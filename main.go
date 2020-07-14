package main

import (
	"net/http"

	"github.com/ciklumtech/router"
	"github.com/ciklumtech/service"
	"github.com/ciklumtech/service/article"
	contentmarketing "github.com/ciklumtech/service/contentMarketing"
)

func main() {
	articleSvc, contentmarketingSvc := setupServices()
	if err := http.ListenAndServe(":8080", router.New(articleSvc, contentmarketingSvc)); err != nil {
		panic(err)
	}
}

func setupServices() (service.Provider, service.Provider) {
	return article.New(), contentmarketing.New()
}
