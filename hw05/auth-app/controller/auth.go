package controller

import (
	"auth-app/repository"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return
	}

	session := repository.GetSession(cookie.Value)
	log.Println("auth request", session)
	claims := jwt.MapClaims{}
	claims["session_id"] = session.Id
	claims["user_id"] = session.UserId
	claims["user_email"] = session.UserEmail
	claims["expiration_in"] = session.ExpiresIn.Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return
	}

	w.Header().Add("x-auth-token", token)
}
