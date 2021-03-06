package urlcontroller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hariprathap-hp/tinyURL_URL_API/domain/urldomain"
	"github.com/hariprathap-hp/tinyURL_URL_API/services"
	"github.com/hariprathap-hp/utils_repo/errors"
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
	var url urldomain.Url
	bindErr := c.ShouldBindJSON(&url)
	if bindErr != nil {
		c.JSON(http.StatusInternalServerError, errors.NewError("bind err : json binding failed at controller"))
		return
	}

	result, err := services.UrlService.ListURLs(url)
	if err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
	c.JSON(http.StatusOK, result)
}

func RedirectURL(c *gin.Context) {
	var url urldomain.Url
	url.TinyURL = c.Request.URL.Query().Get("tiny_url")
	result, err := services.UrlService.RedirectURL(url)
	if err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, *result)
}
