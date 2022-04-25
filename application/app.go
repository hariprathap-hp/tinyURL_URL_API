package application

import (
	"test3/hariprathap-hp/DesignTinyURL/tinyURL_URL_API/dataResources/postgres"
	"test3/hariprathap-hp/DesignTinyURL/tinyURL_URL_API/dataResources/redis"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	mapURLs()
	postgres.Connect()
	redis.Client.Set("key1", "value1", 0)
	router.Run(":8081")
}
