package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

type Claims struct {
	ID 			string `json:"id"`
	FirstName 	string `json:"firstName"`
	LastName	string `json:"lastName"`
	jwt.StandardClaims
}

// AuthMiddleware gets bearers, parses token and passes User if token exists
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")

		if authorization == "" {
			return
		}

		tokenString := strings.Split(authorization, "Bearer ")[1]
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			return
		}

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			c.Set("user", User{
				ID: claims.ID,
				FirstName: claims.FirstName,
				LastName: claims.LastName,
			})
		}
	}
}
