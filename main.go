package main

import (
	"manage-user/config"
	"manage-user/server"
)

func main() {
	cfg := config.LoadFileConfig("env/config.yaml")
	e, db := server.Start(cfg)

	server.Shutdown(e, db)
}
