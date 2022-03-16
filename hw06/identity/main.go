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

var postgresConnection *pgx.Conn

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	initDb()
	defer postgresConnection.Close(context.Background())

	r := gin.Default()
	r.GET("/", func (c *gin.Context) {
		c.JSON(200, "Hello to identity service!")
	})

	r.Use(gin.CustomRecovery(func (c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal error",
		})
	}))

	r.POST("/register", web.Register(
		user.NewPsqlRepository(postgresConnection)),
	)
	r.POST("/login", web.Login(
		user.NewPsqlRepository(postgresConnection),
		session.NewSessionRepository(postgresConnection),
	))
	r.POST("/auth", web.Auth(user.NewPsqlRepository(postgresConnection)))
	err := r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("Server is not started: %s", err)
	}
}

func initDb() {
	var err error
	postgresConnection, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("DB connection error: %s", err)
	}
}