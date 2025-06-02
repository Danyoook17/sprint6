package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Структура сервера
type Server struct {
	Logger *log.Logger
	Server *http.Server
}

// Функция создания сервера
func NewServer(logger *log.Logger) *Server {
	// Создаем новый роутер
	mux := http.NewServeMux()
	
	// Регистрируем хендлеры
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	// Создаем и настраиваем http.Server
	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ErrorLog:       logger,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    15 * time.Second,
	}

	return &Server{
		Logger: logger,
		Server: server,
	}
}
