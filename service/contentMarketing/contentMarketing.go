package contentmarketing

import (
	"github.com/ciklumtech/model"
	"github.com/ciklumtech/service"
)

const contentMarketingURL = "https://storage.googleapis.com/aller-structure-task/contentmarketing.json"

type ContentMarketing struct{}

func New() *ContentMarketing {
	return &ContentMarketing{}
}

func (c *ContentMarketing) GetAll() (model.Response, error) {
	return service.GetAll(contentMarketingURL)
}
