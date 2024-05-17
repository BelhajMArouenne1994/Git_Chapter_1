package handlers

import (
	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"


	"github.com/BelhajMArouenne1994/GIT_CHAPTER_1/soap_to_rest_sfmc/models"
	"github.com/BelhajMArouenne1994/GIT_CHAPTER_1/soap_to_rest_sfmc/services"
)

func GetSubscriberByIDHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataExtensionRequest models.DataExtensionRequest
		if err := c.ShouldBindUri(&dataExtensionRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ctx := c.Request.Context()

		result, err := services.RetrieveDataExtensionByCustomerKey(ctx, dataExtensionRequest)
		if err != nil {
			err := &gin.Error{
				Err:  CreateNewHttpError(err.Error(), err.Error(), 400),
				Type: gin.ErrorTypePublic,
			}
			c.Errors = append(c.Errors, err)
			c.Abort() // Abort the context if the header is not application/json			return
		}

		// You can now call MongoDB service to store the retrieved data
		// mongoService.StoreDataExtensionByCustomerKey(result)

		c.JSON(http.StatusOK, gin.H{"message": result})
	}
}

func GetSubscribersHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataExtensionRequest models.DataExtensionRequest
		if err := c.ShouldBindUri(&dataExtensionRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ctx := c.Request.Context()

		result, err := services.RetrieveDataExtensionByCustomerKey(ctx, dataExtensionRequest)
		if err != nil {
			err := &gin.Error{
				Err:  CreateNewHttpError(err.Error(), err.Error(), 400),
				Type: gin.ErrorTypePublic,
			}
			c.Errors = append(c.Errors, err)
			c.Abort() // Abort the context if the header is not application/json			return
		}

		// You can now call MongoDB service to store the retrieved data
		// mongoService.StoreDataExtensionByCustomerKey(result)

		c.JSON(http.StatusOK, gin.H{"message": result})
	}
}