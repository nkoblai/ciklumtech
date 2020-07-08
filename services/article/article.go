package article

import (
	"github.com/ciklumtech/models"
	"github.com/ciklumtech/services"
)

const articlesURL = "https://storage.googleapis.com/aller-structure-task/articles.json"

type Artice struct{}

func New() *Artice {
	return &Artice{}
}

func (a *Artice) GetAll() (models.Response, error) {
	return services.GetAll(articlesURL)
}
