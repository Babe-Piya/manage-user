package main

import (
	"os"

	"manage-user/config"
	"manage-user/server"
)

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	cfg := config.LoadFileConfig(configPath)
	e, db := server.Start(cfg)

	server.Shutdown(e, db)
}
