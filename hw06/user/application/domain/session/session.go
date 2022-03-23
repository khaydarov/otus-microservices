package session

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"hw06/user/domain/user"
	"os"
	"time"
)

// Token is a refresh token
type Token struct {
	Value string
}

type Session struct {
	ID 			int
	UserID 		user.ID
	Token 		Token
	UserAgent 	string
	Ip 			string
	Expires		time.Time
	CreatedAt	time.Time
}

// NewToken returns new session token
func NewToken() Token {
	return Token{
		uuid.New().String(),
	}
}

// CreateSession returns new session
func CreateSession(userID user.ID, userAgent string, ip string) Session {
	return Session{
		UserID: userID,
		Token: NewToken(),
		UserAgent: userAgent,
		Ip: ip,
		Expires: time.Now().Add(30 * 24 * time.Hour),
		CreatedAt: time.Now(),
	}
}

// CreateAccessToken returns access token
func CreateAccessToken(user *user.User) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["id"] = user.ID.Value
	atClaims["firstName"] = user.FirstName
	atClaims["lastName"] = user.LastName
	atClaims["expires"] = time.Now().Add(15 * time.Minute).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	return at.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
