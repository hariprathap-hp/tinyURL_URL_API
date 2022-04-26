package services

import "test3/hariprathap-hp/DesignTinyURL/tinyURL_URL_API/domain/cache"

var (
	CacheService cacheServicesInterface = &cacheServices{}
)

type cacheServices struct{}

type cacheServicesInterface interface {
	Get() string
	Set([]string)
	SetKey(key string)
}

func (cs *cacheServices) Get() string {
	var c cache.RedisCache
	key := c.Get()
	return key
}

func (cs *cacheServices) Set(keys []string) {
	var c cache.RedisCache
	c.Set(keys)
}

func (cs *cacheServices) SetKey(key string) {
	var c cache.RedisCache
	c.SetKey(key)
}
