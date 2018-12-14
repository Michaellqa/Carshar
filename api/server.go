package api

import (
	"fmt"
	"net/http"
)

type Server struct {
	port int
	mux  http.Handler
	done chan struct{}
}

func NewServer(port int, mux http.Handler, done chan struct{}) *Server {
	return &Server{port: port, mux: mux, done: done}
}

func (s *Server) Start() {
	addr := fmt.Sprintf(":%d", s.port)
	http.ListenAndServe(addr, s.mux)
}

func (s *Server) Stop() {
	s.done <- struct{}{}
}
