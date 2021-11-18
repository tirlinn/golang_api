package lesson

import (
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Server struct {
	EchoServer *echo.Echo
}

func (s *Server) Run(port string) error {
	if err := s.EchoServer.Start(port); err != http.ErrServerClosed {
		log.Fatal(err)
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.EchoServer.Shutdown(ctx)
}

func NewServer() *Server {
	return &Server{
		EchoServer: echo.New(),
	}
}
