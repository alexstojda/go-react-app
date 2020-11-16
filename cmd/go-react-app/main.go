package main

import (
	"go-react-app/web"
	"os"
)

func main() {

	spaPath := os.Getenv("SPA_PATH")
	if spaPath == "" {
		spaPath = "./html"
	}

	server := web.NewServer(spaPath)

	server.StartServer()
}
