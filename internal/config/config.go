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
)

type Config struct {
	HTTP     HTTPConfig
	Postgres PostgresConfig
	Auth     AuthConfig
}

type HTTPConfig struct {
	Port               string        `mapstructure:"port"`
	ReadTimeout        time.Duration `mapstructure:"readTimeout"`
	WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
	MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
}

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

type AuthConfig struct {
	JWT          JWTConfig
	PasswordSalt string
}

type JWTConfig struct {
	AccessTokenTTL  time.Duration `mapstructure:"accessTokenTTL"`
	RefreshTokenTTL time.Duration `mapstructure:"refreshTokenTTL"`
	SigningKey      string
}

func New(configDir string) (*Config, error) {
	populateDefaults()

	if err := readConfigFile(configDir); err != nil {
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
	cfg.Postgres.Host = os.Getenv("POSTGRES_HOST")
	cfg.Postgres.Port = os.Getenv("POSTGRES_PORT")
	cfg.Postgres.Username = os.Getenv("POSTGRES_USERNAME")
	cfg.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")

	cfg.Auth.PasswordSalt = os.Getenv("PASSWORD_SALT")
	cfg.Auth.JWT.SigningKey = os.Getenv("JWT_SIGNING_KEY")
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("postgres", &cfg.Postgres); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("jwt", &cfg.Auth.JWT); err != nil {
		return err
	}

	return nil
}

func readConfigFile(configDir string) error {
	viper.AddConfigPath(configDir)
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.MergeInConfig()
}

func populateDefaults() {
	viper.SetDefault("http.port", defaultHTTPPort)
	viper.SetDefault("http.max_header_megabytes", defaultHTTPMaxHeaderMegabytes)
	viper.SetDefault("http.timeouts.read", defaultHTTPRWTimeout)
	viper.SetDefault("http.timeouts.write", defaultHTTPRWTimeout)
	viper.SetDefault("auth.accessTokenTTL", defaultAccessTokenTTL)
	viper.SetDefault("auth.refreshTokenTTL", defaultRefreshTokenTTL)
}
