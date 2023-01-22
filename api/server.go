package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Server struct {
	e *echo.Echo
}

func NewServer() *Server {
	e := NewAPIServer()
	return &Server{e: e}
}

// Start server will block until exit with gracefull shutdown
func (s *Server) Start(port int) {
	log.Println("starting server...")
	go func() {
		err := s.e.Start(":" + strconv.Itoa(port))
		if err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Printf("Listening on %d\n", port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Println("Shutting down server... Reason:", sig)
	// teardown logic...0
	if err := s.e.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	log.Println("Server gracefully stopped")
}
