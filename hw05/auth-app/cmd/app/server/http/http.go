package http

import (
	"auth-app/cmd/app/server"
	"auth-app/controller"
	"auth-app/repository"
	"fmt"
	"log"
	"net/http"
)

// Server is basic HTTP server that uses net/http
type Server struct {
	SessionRepository repository.SessionRepository
}

// Run initializes routes and listens requests
func (s *Server) Run(config server.Config) {
	http.HandleFunc("/login", controller.Login(s.SessionRepository))
	http.HandleFunc("/auth", controller.Auth(s.SessionRepository))

	log.Printf("Auth application started at %s", config.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", config.Port), nil)
}
