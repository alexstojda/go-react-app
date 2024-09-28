package api

import (
	"github.com/gin-gonic/gin"
	"go-react-app/internal/app/api/health"
	"go-react-app/internal/app/api/hello"
)

type Server struct {
	Hello  *hello.Controller
	Health *health.Controller
}

func NewServer() *Server {
	return &Server{
		Hello:  hello.NewHello(),
		Health: health.NewHealth(),
	}
}

func (s *Server) GetHealth(c *gin.Context) {
	s.Health.Get(c)
}

func (s *Server) GetHello(c *gin.Context) {
	s.Hello.Get(c)
}
