package server

import (
	"log"
	"net/http"
	handlers "sprint-6/internal/handlers"
	"time"
)

type Server struct {
	logger *log.Logger
	server *http.Server
}

func CreateServer(logger *log.Logger) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.MainHandle)
	mux.HandleFunc("/upload", handlers.UploadHandle)

	serv := &Server{
		logger: logger,
	}

	serv.server = &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return serv
}

func (serv *Server) HTTPServer() *http.Server {
	return serv.server
}
