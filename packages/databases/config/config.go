package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	POSTGRESQL    PostgreSQLConfig
	MYSQL         MySQLConfig
	SERVER        ServerConfig
	SMTP   		  SMTPConfig
}

type (
	PostgreSQLConfig struct {
		POSTGRESQL_USER string
		POSTGRESQL_PASS string
		POSTGRESQL_HOST string
		POSTGRESQL_PORT string
		POSTGRESQL_NAME string
	}
	MySQLConfig struct {
		MYSQL_USER string
		MYSQL_PASS string
		MYSQL_HOST string
		MYSQL_PORT string
		MYSQL_NAME string
	}
	ServerConfig struct {
		SERVER_HOST string
		SERVER_PORT string
	}
	SMTPConfig struct {
		SMTP_SERVER string
		SMTP_PORT   string
		SMTP_USERNAME string
		SMTP_PASSWORD string
	}
)

func LoadConfig() (*Configuration, error) {

	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			return nil, fmt.Errorf("error loading .env file: %v", err)
		}
	} else {
		fmt.Println(".env file not found, using default values")
	}

	return &Configuration{
		POSTGRESQL: PostgreSQLConfig{
			POSTGRESQL_USER: os.Getenv("POSTGRESQL_USER"),
			POSTGRESQL_PASS: os.Getenv("POSTGRESQL_PASS"),
			POSTGRESQL_HOST: os.Getenv("POSTGRESQL_HOST"),
			POSTGRESQL_PORT: os.Getenv("POSTGRESQL_PORT"),
			POSTGRESQL_NAME: os.Getenv("POSTGRESQL_NAME"),
		},
		MYSQL: MySQLConfig{
			MYSQL_USER: os.Getenv("MYSQL_USER"),
			MYSQL_PASS: os.Getenv("MYSQL_PASS"),
			MYSQL_HOST: os.Getenv("MYSQL_HOST"),
			MYSQL_PORT: os.Getenv("MYSQL_PORT"),
			MYSQL_NAME: os.Getenv("MYSQL_NAME"),
		},
		SERVER: ServerConfig{
			SERVER_HOST: os.Getenv("SERVER_HOST"),
			SERVER_PORT: os.Getenv("SERVER_PORT"),
		},
		SMTP: SMTPConfig{
			SMTP_SERVER: os.Getenv("SMTP_HOST"),
			SMTP_PORT:   os.Getenv("SMTP_PORT"),
			SMTP_USERNAME: os.Getenv("SMTP_USER"),
			SMTP_PASSWORD: os.Getenv("SMTP_PASSWORD"),
		},
	}, nil
}