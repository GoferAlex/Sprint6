package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// структура сервера
type Server struct {
	Addr         string
	Handler      http.Handler
	ErrorLog     *log.Logger
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func CreateServer(flog *log.Logger) *Server {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.DownloadHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	s := Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     flog,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &s
}
