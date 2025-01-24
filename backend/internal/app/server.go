package app

import (
	"backend/internal/handler"
	"backend/internal/router"
	"log"
	"net/http"
	"os"
)

func NewServer(handlers handler.Handlers) *Server {
	return &Server{
		handler: handlers,
	}
}

func (s *Server) InitRouteAndServe() {
	mux := http.NewServeMux()

	// mux.HandleFunc("/", s.handler.Auth.ReturnHelloWorld)

	mux.Handle("/", router.InitRouter(s.handler))

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s\n", port)
	http.ListenAndServe(":"+port, mux)
}
