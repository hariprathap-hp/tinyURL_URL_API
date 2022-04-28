package services

import (
	"encoding/json"
	"strings"
	"test3/hariprathap-hp/DesignTinyURL/tinyURL_URL_API/domain/urldomain"
	"test3/hariprathap-hp/system_design/utils_repo/dateutils"
	"test3/hariprathap-hp/system_design/utils_repo/errors"
	"test3/hariprathap-hp/test/golang-restclient/rest"
	"time"
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
	return &result, nil
}

func (url *urlService) DeleteURL(url_obj urldomain.Url) *errors.RestErr {
	delErr := url_obj.Delete()
	if delErr != nil {
		return delErr
	}
	return nil
}

func (url *urlService) RedirectURL(url_obj urldomain.Url) (*string, *errors.RestErr) {
	result, redirErr := url_obj.Redirect()
	if redirErr != nil {
		return nil, redirErr
	}
	return result, nil
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
