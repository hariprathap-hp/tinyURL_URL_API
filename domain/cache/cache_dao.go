package cache

import "test3/hariprathap-hp/DesignTinyURL/tinyURL_URL_API/dataResources/redis"

func (rc *RedisCache) Set(keys []string) {
	//we are implement LPUSH. Create a new list "url_keys" and push the keys received from KGS
	redis.Client.LPush("urlkeys", keys)
}

func (rc *RedisCache) SetKey(key string) {
	redis.Client.LPush("urlkeys", key)
}

func (rc *RedisCache) Get() string {
	//at any point of time, we are going to require only one key. So, we can use LPOP command to get the top most key
	key := redis.Client.LPop("urlkeys")
	return key.Val()
}
