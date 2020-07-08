package router

import (
	"net/http"

	"github.com/ciklumtech/handler"
	"github.com/ciklumtech/service/article"
	"github.com/ciklumtech/service/contentmarketing"
)

func registerHandler(s *http.ServeMux, h handler.Service) {
	s.HandleFunc("/merge", h.Merge)
}

func New() *http.ServeMux {

	articleService := article.New()
	contentMarketingService := contentmarketing.New()

	server := http.NewServeMux()
	handler := handler.New(articleService, contentMarketingService)

	registerHandler(server, *handler)
	return server
}
