package server

// Config is a server configuration structure
// Port is a port that application will listen
type Config struct {
	Port string
}

// Server is an interface for any servers
type Server interface {
	Run(config Config)
}
