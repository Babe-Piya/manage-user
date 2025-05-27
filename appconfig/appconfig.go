package appconfig

type AppConfig struct {
	ServerPort string
	MongoDB    MongoDBConfig
	JwtSecret  string
}

type MongoDBConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	DatabaseName string
}
