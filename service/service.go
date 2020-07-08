package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ciklumtech/model"
)

// Provider is interface for any service
type Provider interface {
	GetAll() (model.Response, error)
}

// GetAll returs response from rovided url
func GetAll(url string) (model.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return model.Response{}, err
	}
	defer resp.Body.Close()
	var response model.Response
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Response{}, err
	}
	if err := json.Unmarshal(b, &response); err != nil {
		return model.Response{}, err
	}
	return response, err
}
