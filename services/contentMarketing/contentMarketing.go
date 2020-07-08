package contentmarketing

import (
	"github.com/ciklumtech/models"
	"github.com/ciklumtech/services"
)

const contentMarketingURL = "https://storage.googleapis.com/aller-structure-task/contentmarketing.json"

type ContentMarketing struct{}

func New() *ContentMarketing {
	return &ContentMarketing{}
}

func (c *ContentMarketing) GetAll() (models.Response, error) {
	return services.GetAll(contentMarketingURL)
}
