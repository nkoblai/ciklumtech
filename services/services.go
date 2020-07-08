package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ciklumtech/models"
)

// Provider is interface for any service
type Provider interface {
	GetAll() (models.Response, error)
}

// GetAll returs response from provided url
func GetAll(url string) (models.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return models.Response{}, err
	}
	defer resp.Body.Close()
	var response models.Response
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.Response{}, err
	}
	if err := json.Unmarshal(b, &response); err != nil {
		return models.Response{}, err
	}
	return response, err
}
