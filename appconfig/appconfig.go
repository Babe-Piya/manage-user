package appconfig

import "time"

type AppConfig struct {
	ServerPort    string
	MongoDB       MongoDBConfig
	JwtSecret     string
	TokenTime     time.Duration // minute
	CountUserTime string
}

type MongoDBConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	DatabaseName string
}
