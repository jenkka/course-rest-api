package app

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Server struct {
	Address string
}

func (s *Server) Start() {
	server := &http.Server{
		Addr:         s.Address,
		Handler:      NewRouter(),
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute * 5,
	}

	fmt.Printf("Server started on %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %s", err)
	}
}
