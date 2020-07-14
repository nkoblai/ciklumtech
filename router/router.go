package router

import (
	"net/http"

	"github.com/ciklumtech/handler"
	"github.com/ciklumtech/service"
)

func registerHandler(s *http.ServeMux, h handler.Service) {
	s.HandleFunc("/merge", h.Merge)
}

// New creates new ServeMux
func New(a service.Provider, cm service.Provider) *http.ServeMux {
	server := http.NewServeMux()
	handler := handler.New(a, cm)
	registerHandler(server, *handler)
	return server
}
