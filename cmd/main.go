package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "Server: ", log.LstdFlags|log.Lshortfile)
	s := server.NewServer(logger)

	logger.Println("Запуск на порту :8080")

	err := s.Server.ListenAndServe()
	if err != nil {
		logger.Fatal("Ошибка сервера:", err.Error())
		return
	}

}
