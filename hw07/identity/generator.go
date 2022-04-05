package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	token, err := createAccessToken()
	fmt.Println(token, err)
}

func createAccessToken() (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["id"] = uuid.NewString()
	atClaims["firstName"] = "Sample firstname"
	atClaims["lastName"] = "Sample lastname"
	atClaims["expires"] = time.Now().Add(15 * time.Minute).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	return at.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
