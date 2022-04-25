package services

import (
	"test3/hariprathap-hp/system_design/utils_repo/errors"
)

var (
	UrlService urlServicesInterface = &urlService{}
)

type urlService struct {
}

type urlServicesInterface interface {
	CreateURL() *errors.RestErr
	DeleteURL()
	RedirectURL()
	ListURLs()
}

func (url *urlService) CreateURL() *errors.RestErr {

	return nil
}

func (url *urlService) DeleteURL() {

}

func (url *urlService) RedirectURL() {

}

func (url *urlService) ListURLs() {

}
