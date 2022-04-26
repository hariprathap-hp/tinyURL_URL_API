package services

import (
	"encoding/json"
	"fmt"
	"test3/hariprathap-hp/DesignTinyURL/tinyURL_URL_API/domain/urldomain"
	"test3/hariprathap-hp/system_design/utils_repo/dateutils"
	"test3/hariprathap-hp/system_design/utils_repo/errors"
	"test3/hariprathap-hp/test/golang-restclient/rest"
	"time"
)

var (
	UrlService     urlServicesInterface = &urlService{}
	kgsRestAPICall                      = rest.RequestBuilder{
		Timeout: 100 * time.Second,
		BaseURL: "http://localhost:8080",
	}
)

type urlService struct {
}

type urlServicesInterface interface {
	CreateURL(urldomain.Url) *errors.RestErr
	DeleteURL()
	RedirectURL()
	ListURLs()
}

func (url *urlService) CreateURL(url_obj urldomain.Url) *errors.RestErr {
	if valErr := url_obj.Validate(); valErr != nil {
		return valErr
	}
	if loadErr := loadUrl(&url_obj); loadErr != nil {
		return loadErr
	}
	createErr := url_obj.Create()
	if createErr != nil {
		return createErr
	}
	return nil
}

func (url *urlService) DeleteURL() {

}

func (url *urlService) RedirectURL() {

}

func (url *urlService) ListURLs() {

}

func loadUrl(url *urldomain.Url) *errors.RestErr {
	url.CreationDate = dateutils.GetNow()
	url.ExpirationDate = dateutils.GetExpiry()
	key, keyErr := getID()
	if keyErr != nil {
		return keyErr
	}
	url.TinyURL = "http://localhost:8081/" + *key

	return nil
}

func getID() (*string, *errors.RestErr) {
	for {
		if key := CacheService.Get(); key != "" {
			return &key, nil
		}

		response := kgsRestAPICall.Get("/getKeys")
		if response == nil || response.Response == nil {
			return nil, errors.NewInternalServerError("invalid rest client response from Key Generation Service")
		}

		var uniq_keys struct {
			Keys []string `json:"keys"`
		}
		json.Unmarshal(response.Bytes(), &uniq_keys)
		fmt.Println("results are - ", uniq_keys.Keys)
		CacheService.Set(uniq_keys.Keys) //load keys into the cache
	}

}
