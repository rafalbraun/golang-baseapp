package config

import (
	"context"
	"fmt"
	"log"

	"github.com/gorilla/securecookie"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

// Config defines all the configuration variables
type Config struct {
	BaseURL           string `env:"BASE_URL,default=http://localhost/"`
	Mode              string `env:"MODE,default=prod"`
	PortHttp          string `env:"HTTP_PORT,default=8080"`
	PortHttps         string `env:"HTTPS_PORT,default=4040"`
	Database          string `env:"DATABASE,default=sqlite"`
	DatabaseHost      string `env:"MYSQL_HOST,default=127.0.0.1"`
	DatabasePort      string `env:"MYSQL_PORT,default=3306"`
	DatabaseName      string `env:"MYSQL_DATABASE,default=baseapp"`
	DatabaseUsername  string `env:"MYSQL_USER,default=baseapp"`
	DatabasePassword  string `env:"MYSQL_PASSWORD,default=baseapp"`
	SMTPUsername      string `env:"SMTP_USERNAME"`
	SMTPPassword      string `env:"SMTP_PASSWORD"`
	SMTPHost          string `env:"SMTP_HOST"`
	SMTPPort          string `env:"SMTP_PORT,default=2525"`
	SMTPSender        string `env:"SMTP_SENDER"`
	RequestsPerMinute int    `env:"REQUESTS_PER_MINUTE,default=5"`
	DefaultPageSize   int    `env:"REQUESTS_PER_MINUTE,default=10"`
	CacheParameter    string `env:"CACHE_PARAMETER"`
	CacheMaxAge       int    `env:"CACHE_MAX_AGE"`
	CookieSecret      string `env:"COOKIE_SECRET"`
	CrtFile           string `env:"CRT_FILE,default="`
	KeyFile           string `env:"KEY_FILE,default="`
}

// https://stackoverflow.com/questions/39623052/os-getenv-dosent-work-after-set-env-variable
func LoadEnv() *Config {
	var config Config

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ctx := context.Background()
	if err := envconfig.Process(ctx, &config); err != nil {
		log.Fatal(err)
	}
	if config.CookieSecret == "" {
		config.CookieSecret = string(securecookie.GenerateRandomKey(32))
	}

	return &config
}

func (config *Config) HttpPortString() string {
	return fmt.Sprintf(":%s", config.PortHttp)
}

func (config *Config) HttpsPortString() string {
	return fmt.Sprintf(":%s", config.PortHttps)
}
