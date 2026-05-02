package main

import (
	"log"
	"os"
	server "sprint-6/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	serv := server.CreateServer(logger)

	logger.Fatal(serv.HTTPServer().ListenAndServe())
}
