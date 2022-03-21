package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	var static = http.StripPrefix("/", http.FileServer(http.Dir("./static")))
	var wrapper = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		static.ServeHTTP(w, r)
	})

	http.Handle("/", wrapper)

	log.Printf("Listening on :%s", os.Getenv("APP_PORT"))
	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), nil)
	if err != nil {
		log.Fatal(err)
	}
}
