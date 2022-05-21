package cache

import (
	"context"

	"github.com/hariprathap-hp/tinyURL_URL_API/dataResources/redis"
)

func (rc *RedisCache) Set(keys []string) {
	//we are implement LPUSH. Create a new list "url_keys" and push the keys received from KGS
	redis.Client.LPush(context.Background(), "urlkeys", keys)
}

func (rc *RedisCache) SetKey(key string) {
	redis.Client.LPush(context.Background(), "urlkeys", key)
}

func (rc *RedisCache) Get() string {
	//at any point of time, we are going to require only one key. So, we can use LPOP command to get the top most key
	key := redis.Client.LPop(context.Background(), "urlkeys")
	return key.Val()
}
