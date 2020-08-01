package main

import (
	"context"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/joho/godotenv"
	"log"
	"main/Lib"
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

	// Creating mongo client for Logging proposes
	mongoContext, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongo := Lib.NewMongoClient(&Lib.MongoClient{
		Host:    os.Getenv("MONGODB_LOGGER_HOST"),
		Port:    os.Getenv("MONGODB_LOGGER_PORT"),
		Context: mongoContext,
	})

	client, err := mongo.Client()

	if err != nil {
		log.Fatal(err)
	}

	memcachedClient := memcache.New(os.Getenv("MEMCACHED_URL") + ":" + os.Getenv("MEMCACHED_PORT"))
	distanceHandler := handler.NewDistanceHandler(memcachedClient, client)

	sm := http.NewServeMux()
	sm.Handle("/distance", distanceHandler)

	server := http.Server{
		Addr:         os.Getenv("API_URL") + os.Getenv("API_PORT"),
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
