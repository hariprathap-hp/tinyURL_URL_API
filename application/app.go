package application

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hariprathap-hp/tinyURL_URL_API/dataResources/postgres"
	"github.com/hariprathap-hp/tinyURL_URL_API/dataResources/redis"
)

var (
	router = gin.Default()
)

func StartApp() {
	mapURLs()
	postgres.Connect()
	redis.Client.Set(context.Background(), "key1", "value1", 0)
	router.Run(":8081")
}
