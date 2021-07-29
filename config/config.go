package config

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Grpc     GrpcConfig
	Auth     AuthConfig
}

type ServerConfig struct {
	AppVersion string
	Port       string
	Mode       string
}

type PostgresConfig struct {
	Hostname   string
	Port       string
	Username   string
	Password   string
	DBName     string
	DriverName string
}

type RedisConfig struct {
	Hostname string
	Port     string
	Password string
}

type GrpcConfig struct {
	MaxConnectionIdle time.Duration
	Timeout           time.Duration
	MaxConnectionAge  time.Duration
	Port              string
}

type AuthConfig struct {
	AccessTokenLifeTime  int
	RefreshTokenLifeTime int
}

func ReadConfig(filename string) (*viper.Viper, error) {

	v := viper.New()
	v.SetConfigName(filename)
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			return nil, errors.New("Config file not found")
		}
		// Config file was found but another error was produced
		return nil, err
	}

	return v, nil

}

func GetConfig(pathname string) (*Config, error) {

	configFile, err := ReadConfig(pathname)

	if err != nil {
		return nil, err
	}

	var config Config

	if err = configFile.Unmarshal(&config); err != nil {
		return nil, err
	}

	if os.Getenv("config") == "production" {
		config.Postgres.Username = os.Getenv("POSTGRES_USER")
		config.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")
		config.Postgres.DBName = os.Getenv("POSTGRES_DB")
		config.Redis.Password = os.Getenv("REDIS_PASSWORD")
	}

	return &config, nil

}

func GetConfigByName(name string) string {

	pathname := "config-dev"

	if name == "production" {
		pathname = "config-prod"
	} else if name == "test" {
		pathname = "config-test"
	}

	return fmt.Sprintf("./config/%s", pathname)

}
