package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Создаем новый логгер
	logger := log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Создаем новый сервер
	server := server.NewServer(logger)

	// Логируем сообщение о старте сервера
	server.Logger.Println("Starting server on", server.HTTPServer.Addr)

	// Запускаем сервер
	if err := server.HTTPServer.ListenAndServe(); err != nil {
		server.Logger.Fatal("ListenAndServe:", err)
	}
}
