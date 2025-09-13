package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

func NewServer(logger *log.Logger) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.MainHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	return &Server{
		Logger: logger,
		Server: &http.Server{
			Addr:         ":8080",
			Handler:      mux,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
			IdleTimeout:  15 * time.Second,
			ErrorLog:     logger,
		},
	}
}
