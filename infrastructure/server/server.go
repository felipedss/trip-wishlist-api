package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Port   string
	Server *gin.Engine
}

func NewServer() Server {
	return Server{
		Port:   "8080",
		Server: gin.Default(),
	}
}

func (s *Server) Run(router *gin.Engine) {
	log.Print("Server is running on port " + s.Port)
	log.Fatal(router.Run(":" + s.Port))
}
