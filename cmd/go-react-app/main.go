package main

import (
	"os"

	"go-react-app/web"
)

func main() {
	spaPath := os.Getenv("SPA_PATH")
	server := web.NewServer(spaPath)

	server.StartServer()
}
