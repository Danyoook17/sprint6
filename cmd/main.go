package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Создаем логгер
	logger := log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Создаем сервер
	srv := server.NewServer(logger)
	
	// Запускаем сервер
	if err := srv.Server.ListenAndServe(); err != nil {
		srv.Logger.Fatal("ListenAndServe:", err)
	}
}
