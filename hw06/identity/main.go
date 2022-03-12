package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"hw06/identity/infrastructure/session"
	"hw06/identity/infrastructure/user"
	"hw06/identity/ui/web"
	"log"
	"net/http"
	"os"
)

const currentVersion = "v1"

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	connection, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("DB connection error: %s", err)
	}
	defer connection.Close(context.Background())

	r := gin.Default()
	r.GET("/", func (c *gin.Context) {
		c.JSON(200, "Hello to identity service!")
	})

	r.Use(gin.CustomRecovery(func (c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"error": err,
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"error": "Internal error",
		})
	}))

	group := r.Group(fmt.Sprintf("/%s", currentVersion))
	{
		group.POST("/register", web.Register(user.NewPsqlRepository(connection)))
		group.POST("/login", web.Login(
			user.NewPsqlRepository(connection),
			session.NewSessionRepository(connection),
		))
		group.POST("/auth", web.Auth(user.NewPsqlRepository(connection)))
	}

	r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}

