package main

import (
	"log"

	handlers "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/soap_to_rest_sfmc/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	dataExtensionsRoute := router.Group("/DataExtensions")
	{
		dataExtensionsRoute.Use(handlers.AuthMiddleware())
		dataExtensionsRoute.POST("/", handlers.CreateDataExtensionHandler())
		dataExtensionsRoute.GET("/", handlers.RetrieveDataExtensionsHandler())
		dataExtensionsRoute.GET("/:customerKey", handlers.RetrieveDataExtensionByCustomerKeyHandler())
		dataExtensionsRoute.GET("/:customerKey/DataExtensionFields", handlers.RetrieveDataExtensionFieldsByCustomerKeyHandler())
		dataExtensionsRoute.GET("/:customerKey/DataExtensionFields/:dataExtensionFieldsCustomerKey", handlers.RetrieveDataExtensionFieldByCustomerKeyAndFieldKeyHandler())
	}

	router.GET("/Describe/:object", handlers.DescribeHandler())

	router.GET("/Subscribers", handlers.GetSubscribersHandler())
	router.GET("/Subscriber/:id", handlers.GetSubscriberByIDHandler())

	log.Fatal(router.Run(":8080"))
}
