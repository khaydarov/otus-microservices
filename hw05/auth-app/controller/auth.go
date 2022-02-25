package controller

import (
	"auth-app/repository"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
)

// Auth is a action that authorizes user with given session identifier
func Auth(sessionRepository repository.SessionRepository) func (w http.ResponseWriter, r *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		sessionCookie, err := r.Cookie("session")
		if err != nil {
			return
		}

		session := sessionRepository.GetSession(sessionCookie.Value)
		if session == nil {
			return
		}

		claims := jwt.MapClaims{}
		claims["user_id"] = session.UserId
		claims["user_name"] = session.UserName
		claims["expiration_in"] = session.ExpiresIn.Unix()
		at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			return
		}

		w.Header().Add("x-auth-token", token)
	}
}
