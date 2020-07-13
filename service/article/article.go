package article

import (
	"github.com/ciklumtech/model"
	"github.com/ciklumtech/service"
)

const articlesURL = "https://storage.googleapis.com/aller-structure-task/articles.json"

// Artice represents article service implementation
type Artice struct{}

// New returns article service implementation
func New() *Artice {
	return &Artice{}
}

// GetAll makes API call to article service
func (a *Artice) GetAll() (model.Response, error) {
	return service.GetAll(articlesURL)
}
