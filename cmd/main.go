package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "Server: ", log.LstdFlags|log.Lshortfile)
	s := server.NewServer(logger)

	logger.Println("Server started :8080")

	err := s.Server.ListenAndServe()
	if err != nil {
		logger.Fatal("Server error:", err.Error())
		return
	}

}
