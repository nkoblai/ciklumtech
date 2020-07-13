package contentmarketing

import (
	"github.com/ciklumtech/model"
	"github.com/ciklumtech/service"
)

const contentMarketingURL = "https://storage.googleapis.com/aller-structure-task/contentmarketing.json"

// ContentMarketing represents content marketing service implementation
type ContentMarketing struct{}

// New represents returns marketing service implementation
func New() *ContentMarketing {
	return &ContentMarketing{}
}

// GetAll makes API call to content marketing service
func (c *ContentMarketing) GetAll() (model.Response, error) {
	return service.GetAll(contentMarketingURL)
}
