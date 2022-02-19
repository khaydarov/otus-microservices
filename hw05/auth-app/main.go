package main

import (
	"auth-app/controller"
	"auth-app/repository"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = repository.InitStorage()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/auth", controller.Auth)

	log.Printf("Service started at %s", os.Getenv("APP_PORT"))
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), nil)
}