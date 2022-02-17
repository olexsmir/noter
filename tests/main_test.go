package tests

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/flof-ik/noter/internal/repository"
	"github.com/flof-ik/noter/internal/service"
	v1 "github.com/flof-ik/noter/internal/transport/rest/v1"
	"github.com/flof-ik/noter/pkg/cache"
	"github.com/flof-ik/noter/pkg/database"
	"github.com/flof-ik/noter/pkg/hash"
	"github.com/flof-ik/noter/pkg/token"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

type APITestSuite struct {
	suite.Suite

	db     *sqlx.DB
	stopDB func()

	handler  *v1.Handler
	services *service.Services
	repos    *repository.Repositorys

	hasher       hash.PasswordHasher
	tokenManager token.TokenManager
}

func TestAPISuite(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	suite.Run(t, new(APITestSuite))
}

func (s *APITestSuite) SetupSuite() {
	if db, stop, err := startDbAndConnect(context.Background()); err != nil {
		s.FailNow("Failed to connect to database")
	} else {
		s.db = db
		s.stopDB = stop
	}

	s.initDeps()

	if err := s.populateDB(); err != nil {
		s.FailNow("Failed to populate database")
	}
}

func (s *APITestSuite) TearDownSuite() {
}

func (s *APITestSuite) initDeps() {
	repos := repository.NewRepositorys(s.db)
	memCache := cache.NewMemoryCache()
	hasher := hash.NewSHA1Hasher("salt")
	tokenManager, err := token.NewManager("key")
	if err != nil {
		s.FailNow("Failed to initialize token manager", err)
	}

	services := service.NewServices(service.Deps{
		Repos:           repos,
		Hasher:          hasher,
		TokenManager:    tokenManager,
		Cache:           memCache,
		AccessTokenTTL:  15 * time.Minute,
		RefreshTokenTTL: 30 * time.Minute,
		CacheTTL:        int64(time.Minute.Seconds()),
	})

	s.repos = repos
	s.services = services
	s.handler = v1.NewHandler(services, tokenManager)
	s.hasher = hasher
	s.tokenManager = tokenManager
}

func startDbAndConnect(ctx context.Context) (*sqlx.DB, func(), error) {
	req := testcontainers.ContainerRequest{
		Image:        "postgres:14",
		ExposedPorts: []string{"5432/tcp"},
		Env:          map[string]string{"POSTGRES_PASSWORD": "qwerty"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
	}

	conn, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, func() {}, err
	}

	stop := func() { _ = conn.Terminate(ctx) }
	defer func() {
		if err != nil {
			stop()
		}
	}()

	port, err := conn.MappedPort(ctx, "5432")
	if err != nil {
		return nil, func() {}, err
	}

	host, err := conn.Host(ctx)
	if err != nil {
		return nil, func() {}, err
	}

	db, err := database.NewConnection(database.ConnInfo{
		Host:     host,
		Port:     port.Port(),
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	})

	return db, stop, nil
}

func (s *APITestSuite) populateDB() error {
	_, err := s.db.Exec(`
        CREATE TABLE users (id serial not null unique, name varchar(255) not null, email varchar(255) not null unique, password varchar(255) not null, registred_at time not null, last_visit_at time not null);
        CREATE TABLE sessions (id serial not null unique, user_id  integer REFERENCES users (id) unique, refresh_token varchar(255) not null, expires_at time not null);
        CREATE TABLE notes (id serial not null unique, author_id integer REFERENCES users (id), title varchar(255) not null, content text not null, created_at time not null, updated_at time not null);
        CREATE TABLE notebooks (id serial not null unique, author_id integer REFERENCES users (id), name varchar(255) not null unique, description varchar(255), created_at time not null, updated_at time not null);
        ALTER TABLE notes ADD notebook_id integer REFERENCES notebooks (id) not null;
        ALTER TABLE notes ADD pinted boolean DEFAULT false;
    `)
	if err != nil {
		return err
	}

	_, err = s.db.NamedExec("INSERT INTO users (name, email, password, registred_at, last_visit_at) VALUES (:name, :email, :password, :registred_at, :last_visit_at)", &user)
	if err != nil {
		return err
	}

	_, err = s.db.NamedExec("INSERT INTO notebooks (author_id, name, description, created_at, updated_at) VALUES (:author_id, :name, :description, :created_at, :updated_at)", &notebook)
	if err != nil {
		return err
	}

	_, err = s.db.NamedExec("INSERT INTO notes (author_id, notebook_id, pinted, title, content, created_at, updated_at) VALUES (:author_id, :notebook_id, :pinted, :title, :content, :created_at, :updated_at)", &note)
	if err != nil {
		return err
	}

	return err
}
