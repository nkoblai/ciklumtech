package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/ciklumtech/common"
	"github.com/ciklumtech/models"
	"github.com/ciklumtech/services"
)

type Service struct {
	article          services.Provider
	contentMarketing services.Provider
}

func New(article services.Provider, contentMarketing services.Provider) *Service {
	return &Service{article, contentMarketing}
}

func (s *Service) Merge(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	errsCh := make(chan error, 2)
	var articles models.Response
	var contentMarketing models.Response

	wg.Add(2)
	go func() {
		var err error
		articles, err = s.article.GetAll()
		if err != nil {
			errsCh <- fmt.Errorf("getting articles error: %v", err)
		} else {
			errsCh <- nil
		}
		wg.Done()
	}()
	go func() {
		var err error
		contentMarketing, err = s.contentMarketing.GetAll()
		if err != nil {
			errsCh <- fmt.Errorf("getting content marketing error: %v", err)
		} else {
			errsCh <- nil
		}
		wg.Done()
	}()

	wg.Wait()
	close(errsCh)

	var errMsg string
	for err := range errsCh {
		if err != nil {
			errMsg += fmt.Sprintf("%v; ", err)
		}
	}

	httpError := models.NewHTTPError(errMsg)
	if httpError != nil {
		common.ResponseWithError(w, httpError)
		return
	}

	b, err := json.Marshal(merge(articles, contentMarketing))
	if err != nil {
		common.ResponseWithError(w, models.NewHTTPError(err.Error()))
		return
	}

	common.ResponseWithStatusAndMessage(w, http.StatusOK, b)
	return
}

func merge(articles models.Response, contentMarketing models.Response) []models.ContentMarketing {
	articlesLen := len(articles.Response.Items)
	contentMarketingLen := len(contentMarketing.Response.Items)

	articlesArrLen := 5

	var adCap int

	if (articlesLen/articlesArrLen)-contentMarketingLen > 0 {
		adCap = (articlesLen / articlesArrLen) - contentMarketingLen
	}

	result := make([]models.ContentMarketing, 0, articlesLen+contentMarketingLen+adCap)

	for i := 0; ; i++ {
		lastElemIdx := i*articlesArrLen + articlesArrLen
		if lastElemIdx < articlesLen {
			result = append(result, articles.Response.Items[i*articlesArrLen:lastElemIdx]...)
			if i >= contentMarketingLen {
				result = append(result, models.ContentMarketing{Article: models.Article{Type: "Ad"}})
			} else {
				result = append(result, contentMarketing.Response.Items[i])
			}
			continue
		}
		result = append(result, articles.Response.Items[i*articlesArrLen:]...)
		return result
	}
	return result
}
