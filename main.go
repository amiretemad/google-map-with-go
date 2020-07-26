package main

import (
	"github.com/joho/godotenv"
	"log"
	"main/handler"
	"net/http"
	"os"
	"time"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	distanceHandler := handler.NewDistanceHandler()

	sm := http.NewServeMux()
	sm.Handle("/distance", distanceHandler)

	server := http.Server{
		Addr:         os.Getenv("API_PORT"),
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	_ = server.ListenAndServe()
}
