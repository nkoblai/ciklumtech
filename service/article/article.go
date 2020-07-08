package article

import (
	"github.com/ciklumtech/model"
	"github.com/ciklumtech/service"
)

const articlesURL = "https://storage.googleapis.com/aller-structure-task/articles.json"

type Artice struct{}

func New() *Artice {
	return &Artice{}
}

func (a *Artice) GetAll() (model.Response, error) {
	return service.GetAll(articlesURL)
}
