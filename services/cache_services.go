package services

import (
	"github.com/hariprathap-hp/tinyURL_URL_API/domain/cache"
	"github.com/hariprathap-hp/utils_repo/errors"
)

var (
	CacheService cacheServicesInterface = &cacheServices{}
	c            cache.RedisCache
)

type cacheServices struct{}

type cacheServicesInterface interface {
	Get() string
	Set([]string)
	SetKey(string)
	HSet(string, string) *errors.RestErr
	HGet(string) (*string, *errors.RestErr)
}

func (cs *cacheServices) Get() string {
	key := c.Get()
	return key
}

func (cs *cacheServices) Set(keys []string) {
	c.Set(keys)
}

func (cs *cacheServices) SetKey(key string) {
	c.SetKey(key)
}

func (cs *cacheServices) HSet(tiny_url, orig_url string) *errors.RestErr {
	res := c.HSet(tiny_url, orig_url)
	if res != 1 {
		return errors.NewInternalServerError("caching of tiny_url failed")
	}
	return nil
}

func (cs *cacheServices) HGet(tiny_url string) (*string, *errors.RestErr) {
	res := c.HGet(tiny_url)
	if res == "" {
		return nil, errors.NewInternalServerError("unable to find the original url for short url")
	}
	return &res, nil
}
