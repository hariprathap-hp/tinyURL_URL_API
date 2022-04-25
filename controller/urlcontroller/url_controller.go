package urlcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateURL(c *gin.Context) {

}

func DeleteURL(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "Implement Delete")
}

func ListURLs(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "Implement List")
}

func RedirectURL(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "Implement Redirect")
}
