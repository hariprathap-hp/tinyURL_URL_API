package application

import "test3/hariprathap-hp/DesignTinyURL/tinyURL_URL_API/controller/urlcontroller"

func mapURLs() {
	router.POST("/create", urlcontroller.CreateURL)
	router.POST("/delete", urlcontroller.DeleteURL)
	router.GET("/redirect", urlcontroller.RedirectURL)
	router.GET("/list", urlcontroller.ListURLs)
}
