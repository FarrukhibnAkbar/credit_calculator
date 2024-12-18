package configs

import (
	"errors"
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var (
	conf *Configuration
	once sync.Once
)

func Config() *Configuration {
	once.Do(func() {
		conf = load()
	})
	return conf
}

type Configuration struct {
	HTTPPort    string
	LogLevel    string
	AppName     string
	Environment string

	ServerPort int
	ServerHost string

	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
}

func load() *Configuration {

	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading env file: ", err)
	}

	var config Configuration

	v := viper.New()
	v.AutomaticEnv()

	config.Environment = v.GetString("ENVIRONMENT")
	config.HTTPPort = v.GetString("HTTP_PORT")

	config.LogLevel = v.GetString("LOG_LEVEL")

	config.PostgresDatabase = v.GetString("POSTGRES_DB")
	config.PostgresUser = v.GetString("POSTGRES_USER")
	config.PostgresPassword = v.GetString("POSTGRES_PASSWORD")
	config.PostgresHost = v.GetString("POSTGRES_HOST")
	config.PostgresPort = v.GetInt("POSTGRES_PORT")

	//validate the configuration
	err = config.validate()
	if err != nil {
		log.Fatal("error validating config: ", err)
	}

	return &config
}

func (c *Configuration) validate() error {
	if c.HTTPPort == "" {
		return errors.New("http_port required")
	}
	if c.PostgresDatabase == "" {
		return errors.New("PostgresDatabase required")
	}
	if c.PostgresUser == "" {
		return errors.New("PostgresUser required")
	}
	if c.PostgresPassword == "" {
		return errors.New("PostgresPassword required")
	}
	// ....

	return nil
}
