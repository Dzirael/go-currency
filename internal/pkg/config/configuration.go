package config

import (
	"log"

	"github.com/jessevdk/go-flags"
)

var Config *Configuration

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
	Email    EmailConfiguration
}

type DatabaseConfiguration struct {
	Driver       string `long:"driver" env:"DB_DRIVER" default:"sqlite" description:"Database driver (e.g., sqlite, postgres, mysql)"`
	Dbname       string `long:"dbname" env:"DB_NAME" default:"data/database" description:"Database name"`
	Username     string `long:"username" env:"DB_USERNAME" default:"user" description:"Database username"`
	Password     string `long:"password" env:"DB_PASSWORD" default:"password" description:"Database password"`
	Host         string `long:"host" env:"DB_HOST" default:"localhost" description:"Database host"`
	Port         string `long:"port" env:"DB_PORT" default:"5432" description:"Database port"`
	MaxLifetime  int    `long:"max-lifetime" env:"DB_MAX_LIFETIME" default:"7200" description:"Maximum lifetime for database connections"`
	MaxOpenConns int    `long:"max-open-conns" env:"DB_MAX_OPEN_CONNS" default:"150" description:"Maximum number of open database connections"`
	MaxIdleConns int    `long:"max-idle-conns" env:"DB_MAX_IDLE_CONNS" default:"50" description:"Maximum number of idle database connections"`
}

type ServerConfiguration struct {
	Port string `long:"server-port" env:"SERVER_PORT" default:"8080" description:"Listen to http traffic on this tcp address"`
	Mode string `long:"server-mode" env:"SERVER_MODE" default:"release" description:"Server mode (e.g., debug, release)"`
}

type EmailConfiguration struct {
	From     string `long:"email-from" env:"EMAIL_FROM" description:"Email address to send from"`
	Password string `long:"email-password" env:"EMAIL_PASSWORD" description:"Password for the email account"`
	Smtp     string `long:"email-smtp" env:"EMAIL_SMTP" description:"SMTP server address without port"`
}

// SetupDB initialize configuration
func Setup() {
	var configuration Configuration

	parser := flags.NewParser(&configuration, flags.Default)
	if _, err := parser.Parse(); err != nil {
		log.Fatalf("Error reading eviroment variables, %s", err)
	}

	Config = &configuration
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	return Config
}
