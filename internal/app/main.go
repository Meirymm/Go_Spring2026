package app

import (
	"assignment2/internal/handlers"
	"assignment2/internal/repository"
	"assignment2/internal/repository/_postgres"
	"assignment2/internal/usecase"
	"assignment2/pkg/modules"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(".env file loaded successfully!")
	dbConfig := initPostgreConfig()
	db := _postgres.NewPGXDialect(dbConfig)
	log.Println("Successfully connected to database!")
	repos := repository.NewRepositories(db)
	userUsecase := usecase.NewUserUsecase(repos.UserRepository)
	userHandler := handlers.NewUserHandler(userUsecase)
	mux := http.NewServeMux()
	mux.HandleFunc("/health", userHandler.HealthCheck)
	mux.HandleFunc("/users", handlers.Chain(
		userHandler.HandleUsers,
		handlers.AuthMiddleware,
		handlers.LoggingMiddleware,
	))
	mux.HandleFunc("/users/id", handlers.Chain(
		userHandler.HandleUserByID,
		handlers.AuthMiddleware,
		handlers.LoggingMiddleware,
	))
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
func initPostgreConfig() *modules.PostgreConfig {
	return &modules.PostgreConfig{
		Host:        os.Getenv("DB_HOST"),
		Port:        os.Getenv("DB_PORT"),
		Username:    os.Getenv("DB_USERNAME"),
		Password:    os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
		SSLMode:     os.Getenv("DB_SSLMODE"),
		ExecTimeout: 5 * time.Second,
	}
}