package main

import (
	"manage-user/server"
)

func main() {
	e, db := server.Start()

	server.Shutdown(e, db)
}
