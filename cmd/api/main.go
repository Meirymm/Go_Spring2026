package main

<<<<<<< HEAD
import "assignment2/internal/app"

func main() {
	app.Run()
=======
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
>>>>>>> a0bd3467941b72497a83026b703c10a46cd15d12
}