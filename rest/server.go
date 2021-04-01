package rest

import "net/http"

// Server - API Server
type Server struct {
	//storage
	//publisher
	//log
	//config
	//gin engine
}

// NewAPIServer - wrapper on gin engine, sets up routes i.e endpoints
func NewAPIServer() (s *Server) {
	s = &Server{}
	s.endpoints()
	return
}

// ServerHTTP ...
func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//gin serves
}
