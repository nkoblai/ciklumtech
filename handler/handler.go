package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/ciklumtech/model"
	"github.com/ciklumtech/service"
	"github.com/ciklumtech/web"
)

const (
	articlesArrLen = 5
)

type Service struct {
	article          service.Provider
	contentMarketing service.Provider
}

func New(article service.Provider, contentMarketing service.Provider) *Service {
	return &Service{article, contentMarketing}
}

func (s *Service) Merge(w http.ResponseWriter, r *http.Request) {
	articles, contentMarketing, httpError := s.getData()
	if httpError != nil {
		web.ResponseWithError(w, httpError)
		return
	}

	b, err := json.Marshal(merge(articles, contentMarketing))
	if err != nil {
		web.ResponseWithError(w, model.NewHTTPError(err.Error()))
		return
	}

	web.ResponseWithStatusAndMessage(w, http.StatusOK, b)
	return
}

func (s *Service) getData() (model.Response, model.Response, error) {
	servicesCount := 2

	var wg sync.WaitGroup
	errsCh := make(chan error, servicesCount)
	var articles model.Response
	var contentMarketing model.Response

	wg.Add(servicesCount)
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

	return articles, contentMarketing, model.NewHTTPError(errMsg)
}

func merge(articles model.Response, contentMarketing model.Response) []model.ContentMarketing {
	articlesLen := len(articles.Response.Items)
	contentMarketingLen := len(contentMarketing.Response.Items)

	result := createSlice(articlesLen, contentMarketingLen)

	for i := 0; ; i++ {
		lastElemIdx := i*articlesArrLen + articlesArrLen
		if lastElemIdx < articlesLen {
			appendArticlesAndContentMarketing(i, &articles, &contentMarketing, &result)
			continue
		}
		if lastElemIdx == articlesLen {
			appendArticlesAndContentMarketing(i, &articles, &contentMarketing, &result)
			return result
		}

		result = append(result, articles.Response.Items[i*articlesArrLen:]...)
		if i < contentMarketingLen {
			result = append(result, contentMarketing.Response.Items[i:]...)
		}
		return result
	}
}

func appendArticlesAndContentMarketing(i int, articles, contentMarketing *model.Response, result *[]model.ContentMarketing) {
	lastElemIdx := i*articlesArrLen + articlesArrLen
	contentMarketingLen := len(contentMarketing.Response.Items)
	*result = append(*result, articles.Response.Items[i*articlesArrLen:lastElemIdx]...)
	if i >= contentMarketingLen {
		*result = append(*result, model.ContentMarketing{Article: model.Article{Type: "Ad"}})
	} else {
		*result = append(*result, contentMarketing.Response.Items[i])
	}
}

func createSlice(articlesLen, contentMarketingLen int) []model.ContentMarketing {
	var adCap int
	if (articlesLen/articlesArrLen)-contentMarketingLen > 0 {
		adCap = (articlesLen / articlesArrLen) - contentMarketingLen
	}
	return make([]model.ContentMarketing, 0, articlesLen+contentMarketingLen+adCap)
}
