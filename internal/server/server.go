package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger     *log.Logger
	HTTPServer *http.Server
}

func NewServer(logger *log.Logger) *Server {
	// Создаем роутер
	mux := http.NewServeMux()

	// Регистрация хендлеров
	mux.HandleFunc("/", handlers.Handler)             // Регистрация основного обработчика
	mux.HandleFunc("/upload", handlers.HandlerUpload) // Регистрация обработчика загрузки

	// Создаем новый HTTP-сервер с заданными параметрами
	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger:     logger,
		HTTPServer: httpServer,
	}
}

