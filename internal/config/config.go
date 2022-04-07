package config

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

const (
	defaultHTTPPort               = "8000"
	defaultHTTPRWTimeout          = 10 * time.Second
	defaultHTTPMaxHeaderMegabytes = 1
	defaultAccessTokenTTL         = 15 * time.Minute
	defaultRefreshTokenTTL        = 24 * time.Hour * 30
	defaultPageSize               = 100

	EnvLocal = "local"
	Prod     = "prod"
)

type Config struct {
	Environment string
	HTTP        HTTPConfig
	Postgres    PostgresConfig
	Pagination  PaginationConfig
	Auth        AuthConfig
	CacheTTL    time.Duration `mapstructure:"ttl"`
}

type HTTPConfig struct {
	Host               string        `mapstructure:"host"`
	Port               string        `mapstructure:"port"`
	ReadTimeout        time.Duration `mapstructure:"read_timeout"`
	WriteTimeout       time.Duration `mapstructure:"write_timeout"`
	MaxHeaderMegabytes int           `mapstructure:"max_header_bytes"`
}

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

type PaginationConfig struct {
	PageSize int `mapstructure:"page_size"`
}

type AuthConfig struct {
	JWT          JWTConfig
	PasswordSalt string
}

type JWTConfig struct {
	AccessTokenTTL  time.Duration `mapstructure:"access_token_ttl"`
	RefreshTokenTTL time.Duration `mapstructure:"refresh_token_ttl"`
	SigningKey      string
}

func New(configDir string) (*Config, error) {
	populateDefaults()

	if err := readConfigFile(configDir, os.Getenv("APP_ENV")); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)

	return &cfg, nil
}

func setFromEnv(cfg *Config) {
	cfg.Environment = os.Getenv("APP_ENV")

	cfg.Postgres.Host = os.Getenv("POSTGRES_HOST")
	cfg.Postgres.Port = os.Getenv("POSTGRES_PORT")
	cfg.Postgres.Username = os.Getenv("POSTGRES_USERNAME")
	cfg.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")

	cfg.Auth.PasswordSalt = os.Getenv("PASSWORD_SALT")
	cfg.Auth.JWT.SigningKey = os.Getenv("JWT_SIGNING_KEY")
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("cache.ttl", &cfg.CacheTTL); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("postgres", &cfg.Postgres); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("pagination", &cfg.Pagination); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("jwt", &cfg.Auth.JWT); err != nil {
		return err
	}

	return nil
}

func readConfigFile(configDir string, env string) error {
	viper.AddConfigPath(configDir)
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if env == EnvLocal {
		return nil
	}

	viper.SetConfigName(env)

	return viper.MergeInConfig()
}

func populateDefaults() {
	viper.SetDefault("http.port", defaultHTTPPort)
	viper.SetDefault("http.max_header_megabytes", defaultHTTPMaxHeaderMegabytes)
	viper.SetDefault("http.read_timeout", defaultHTTPRWTimeout)
	viper.SetDefault("http.write_timeout", defaultHTTPRWTimeout)
	viper.SetDefault("auth.access_token_ttl", defaultAccessTokenTTL)
	viper.SetDefault("auth.refresh_token_ttl", defaultRefreshTokenTTL)
	viper.SetDefault("pagination.page_size", defaultPageSize)
}
