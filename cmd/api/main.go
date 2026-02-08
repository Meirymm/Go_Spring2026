package main

import (
	"assignment1/internal/handlers"
	"assignment1/internal/middleware"
	"assignment1/internal/models"
	"log"
	"net/http"
)

func main() {
	store := models.NewTaskStore()
	taskHandler := handlers.NewTaskHandler(store)
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", middleware.Chain(
		taskHandler.HandleTasks,
		middleware.AuthMiddleware,
		middleware.LoggingMiddleware,
	))
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}