package config

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	type env struct {
		appEnv           string
		postgresHost     string
		postgresPort     string
		postgresUsername string
		postgresPassword string
		passwordSalt     string
		jwtSigningKey    string
	}

	type args struct {
		path string
		env  env
	}

	setEnv := func(env env) {
		os.Setenv("APP_ENV", env.appEnv)
		os.Setenv("POSTGRES_HOST", env.postgresHost)
		os.Setenv("POSTGRES_PORT", env.postgresPort)
		os.Setenv("POSTGRES_USERNAME", env.postgresUsername)
		os.Setenv("POSTGRES_PASSWORD", env.postgresPassword)
		os.Setenv("PASSWORD_SALT", env.passwordSalt)
		os.Setenv("JWT_SIGNING_KEY", env.jwtSigningKey)
	}

	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "Ok!",
			args: args{
				path: "fixtures",
				env: env{
					appEnv:           "local",
					postgresHost:     "localhost",
					postgresPort:     "5432",
					postgresUsername: "postgres",
					postgresPassword: "postgres",
					passwordSalt:     "pass",
					jwtSigningKey:    "jwt",
				},
			},
			want: &Config{
				Environment: "local",
				CacheTTL:    3600 * time.Second,
				HTTP: HTTPConfig{
					Host:               "localhost",
					Port:               "80",
					ReadTimeout:        time.Second * 10,
					MaxHeaderMegabytes: 1,
					WriteTimeout:       time.Second * 10,
				},
				Pagination: PaginationConfig{
					PageSize: 10,
				},
				Postgres: PostgresConfig{
					Host:     "localhost",
					Port:     "5432",
					Username: "postgres",
					Password: "postgres",
					DBName:   "test_db",
					SSLMode:  "disable",
				},
				Auth: AuthConfig{
					PasswordSalt: "pass",
					JWT: JWTConfig{
						AccessTokenTTL:  time.Minute * 15,
						RefreshTokenTTL: time.Hour * 24,
						SigningKey:      "jwt",
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setEnv(tt.args.env)

			got, err := New(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr = %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}
