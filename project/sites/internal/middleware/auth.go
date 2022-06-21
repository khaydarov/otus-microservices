package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"os"
)

const BearerSchema = "Bearer "

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}

		tokenString := authHeader[len(BearerSchema):]
		token, err := jwt.ParseWithClaims(
			tokenString,
			jwt.MapClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			},
		)

		if err != nil {
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			err = errors.New("couldn't parse claims")
			return
		}

		context.Set("UserID", claims["id"])
		context.Set("UserEmail", claims["email"])
		context.Next()
	}
}
