package router

import (
	"net/http"

	"github.com/ciklumtech/handlers"
	"github.com/ciklumtech/services/article"
	"github.com/ciklumtech/services/contentmarketing"
)

func registerHandlers(s *http.ServeMux, h handlers.Service) {
	s.HandleFunc("/merge", h.Merge)
}

func New() *http.ServeMux {

	articleService := article.New()
	contentMarketingService := contentmarketing.New()

	server := http.NewServeMux()
	handlers := handlers.New(articleService, contentMarketingService)

	registerHandlers(server, *handlers)
	return server
}
