package util

import (
	"os"
	"strconv"
)

type APPConfig struct {
	AppName                 string
	Port                    string
	Debug                   string
	ThirdPartyServiceConfig ThirdPartyServiceConfig
	Database                DB
	Redis                   Redis
}

type ThirdPartyServiceConfig struct {
	Timeout   int
	Debug     bool
	UseProxy  bool
	HttpProxy string
}

type DB struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DBName       string
	DBSchemaName string
}

type Redis struct {
	Host     string
	Username string
	Password string
}

func InitConfig() APPConfig {
	useProxy, _ := strconv.ParseBool(os.Getenv("USE_PROXY"))
	timeout, _ := strconv.Atoi(os.Getenv("TIMEOUT"))
	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	appConfig := APPConfig{
		AppName: os.Getenv("APP_NAME"),
		Port:    os.Getenv("PORT"),
		Debug:   os.Getenv("DEBUG"),
		ThirdPartyServiceConfig: ThirdPartyServiceConfig{
			Timeout:   timeout,
			Debug:     debug,
			UseProxy:  useProxy,
			HttpProxy: os.Getenv("HTTP_PROXY"),
		},
		Database: DB{
			Host:         os.Getenv("DB_HOST"),
			Port:         os.Getenv("DB_PORT"),
			Username:     os.Getenv("DB_USERNAME"),
			Password:     os.Getenv("DB_PASSWORD"),
			DBName:       os.Getenv("DB_NAME"),
			DBSchemaName: os.Getenv("DB_SCHEMA_NAME"),
		},
		Redis: Redis{
			Host:     os.Getenv("REDIS_HOST"),
			Username: os.Getenv("REDIS_USERNAME"),
			Password: os.Getenv("REDIS_PASSWORD"),
		},
	}

	return appConfig
}
