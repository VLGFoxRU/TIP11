package main

import (
	"log"
	"net/http"

	"example.com/pz11-notes-api/internal/http"
	"example.com/pz11-notes-api/internal/repo"
)

func main() {
	// Инициализация репозитория
	noteRepo := repo.NewNoteRepoMem()

	// Создание маршрутизатора
	r := router.NewRouter(noteRepo)

	// Запуск сервера
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
