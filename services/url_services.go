package services

import (
	"encoding/json"
	"fmt"
	"strings"

	"time"

	"github.com/golang-restclient/rest"
	"github.com/hariprathap-hp/tinyURL_URL_API/dataResources/lru"
	"github.com/hariprathap-hp/tinyURL_URL_API/domain/urldomain"

	"github.com/hariprathap-hp/utils_repo/dateutils"
	"github.com/hariprathap-hp/utils_repo/errors"
)

var (
	UrlService       urlServicesInterface = &urlService{}
	kgsRestClientURL                      = rest.RequestBuilder{
		BaseURL: "http://localhost:8080",
		Timeout: 100 * time.Second,
	}
)

type urlService struct {
}

type urlServicesInterface interface {
	CreateURL(urldomain.Url) (*urldomain.ListURLs, *errors.RestErr)
	DeleteURL(urldomain.Url) *errors.RestErr
	RedirectURL(urldomain.Url) (*string, *errors.RestErr)
	ListURLs(urldomain.Url) (urldomain.UrlsList, *errors.RestErr)
}

func (url *urlService) CreateURL(url_obj urldomain.Url) (*urldomain.ListURLs, *errors.RestErr) {
	if valErr := url_obj.Validate(); valErr != nil {
		return nil, valErr
	}
	key, loadErr := loadUrl(&url_obj)
	if loadErr != nil {
		return nil, loadErr
	}

	createErr := url_obj.Create()
	if createErr != nil {
		if strings.Contains(createErr.Message, "already present in the database") {
			CacheService.SetKey(*key)
		}
		return nil, createErr
	}
	result := urldomain.ListURLs{
		OriginalURL: url_obj.OriginalURL,
		TinyURL:     url_obj.TinyURL,
	}
	lru.Cache.Add(url_obj.TinyURL, url_obj.OriginalURL)
	return &result, nil
}

func (url *urlService) DeleteURL(url_obj urldomain.Url) *errors.RestErr {
	delErr := url_obj.Delete()
	if delErr != nil {
		return delErr
	}
	lru.Cache.Remove(url_obj.TinyURL)
	return nil
}

func (url *urlService) RedirectURL(url_obj urldomain.Url) (*string, *errors.RestErr) {
	val, ok := lru.Cache.Get(url_obj.TinyURL)
	if !ok {
		result, redirErr := url_obj.Redirect()
		if redirErr != nil {
			return nil, redirErr
		}
		lru.Cache.Add(url_obj.TinyURL, *result)
		return result, nil
	}
	fmt.Println("Found in Cache")
	res := fmt.Sprintf("%v", val)
	return &res, nil

}

func (url *urlService) ListURLs(url_obj urldomain.Url) (urldomain.UrlsList, *errors.RestErr) {
	result, listErr := url_obj.List()
	if listErr != nil {
		return nil, listErr
	}
	return result, nil
}

func loadUrl(url *urldomain.Url) (*string, *errors.RestErr) {
	url.CreationDate = dateutils.GetNow()
	url.ExpirationDate = dateutils.GetExpiry()
	key, keyErr := getID()
	if keyErr != nil {
		return nil, keyErr
	}
	url.TinyURL = "http://localhost:8081/" + *key
	return key, nil
}

func getID() (*string, *errors.RestErr) {
	for {
		if key := CacheService.Get(); key != "" {
			return &key, nil
		}
		//we need to make an internal API call to KGS
		response := kgsRestClientURL.Get("/getKeys")
		if response == nil || response.Response == nil {
			return nil, errors.NewInternalServerError("invalid rest client response while trying to contact KGS")
		}
		var keys urldomain.UniqKeys
		if err := json.Unmarshal(response.Bytes(), &keys); err != nil {
			return nil, errors.NewInternalServerError("error while decoding the json value received from KGS")
		}
		CacheService.Set(keys.Keys)
	}
}
