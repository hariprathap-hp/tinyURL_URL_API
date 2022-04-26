package urlcontroller

import (
	"fmt"
	"net/http"
	"test3/hariprathap-hp/DesignTinyURL/tinyURL_URL_API/domain/urldomain"
	"test3/hariprathap-hp/DesignTinyURL/tinyURL_URL_API/services"
	"test3/hariprathap-hp/system_design/utils_repo/errors"

	"github.com/gin-gonic/gin"
)

func CreateURL(c *gin.Context) {
	var url urldomain.Url
	bindErr := c.ShouldBindJSON(&url)
	if bindErr != nil {
		c.JSON(http.StatusInternalServerError, errors.NewError("bind err : json binding failed at controller"))
		return
	}

	result, createErr := services.UrlService.CreateURL(url)
	if createErr != nil {
		c.JSON(http.StatusInternalServerError, createErr)
		return
	}
	c.JSON(http.StatusOK, *result)
}

func DeleteURL(c *gin.Context) {
	var url urldomain.Url
	bindErr := c.ShouldBindJSON(&url)
	if bindErr != nil {
		c.JSON(http.StatusInternalServerError, errors.NewError("bind err : json binding failed at controller"))
		return
	}
	delErr := services.UrlService.DeleteURL(url)
	if delErr != nil {
		c.JSON(delErr.Status, delErr)
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("url %s is deleted", url.TinyURL))
}

func ListURLs(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "Implement List")
}

func RedirectURL(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "Implement Redirect")
}
