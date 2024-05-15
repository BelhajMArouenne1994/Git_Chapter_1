package soap_to_rest_sfmc_api

import (
	"fmt"
	"net/http"

	soap_to_rest_sfmc "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/soap_to_rest_sfmc"
	soap_to_rest_sfmc_handlers "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/soap_to_rest_sfmc_handlers"
	types "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/types"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {

	server := Server{}

	router := gin.New()
	// Utilisation des middlewares
	router.Use(soap_to_rest_sfmc_handlers.ErrorHandler())

	dataExtensionsRoute := router.Group("/DataExtensions")
	{
		// Utilisation de Use pour appliquer les middlewares
		dataExtensionsRoute.Use(
			soap_to_rest_sfmc_handlers.HeaderHandler(),
			soap_to_rest_sfmc_handlers.AuthMiddleware(),
		)
		// Data Extension methods
		dataExtensionsRoute.GET("/", func(c *gin.Context) {
			// Convert *gin.Context to context.Context
			ctx := c.Request.Context()

			// Call the original function
			var result *soap_to_rest_sfmc.RetrieveDEResponseMsg
			var err error
			result, err = soap_to_rest_sfmc.RetrieveDataExtensions(ctx)
			if err != nil {
				// Handle error
				errorResponse(c, err)
				return
			}

			// If needed, use ginCtx for further processing or respond directly
			// For example, assuming a successful operation:
			c.JSON(http.StatusOK, gin.H{"message": result})
		})
		dataExtensionsRoute.GET("/:customerKey", func(c *gin.Context) {
			var dataExtensionRequest = &types.DataExtensionRequest{}
			if err := c.ShouldBindUri(dataExtensionRequest); err != nil {
				c.JSON(400, gin.H{"message": err.Error()})
				return
			}

			// Convert *gin.Context to context.Context
			ctx := c.Request.Context()

			// Call the original function
			var result *soap_to_rest_sfmc.RetrieveDEResponseMsg
			var err error
			result, err = soap_to_rest_sfmc.RetrieveDataExtensionByCustomerKey(ctx, *dataExtensionRequest)
			if err != nil {
				// Handle error
				errorResponse(c, err)
				return
			}

			// If needed, use ginCtx for further processing or respond directly
			// For example, assuming a successful operation:
			c.JSON(http.StatusOK, gin.H{"message": result})
		})
		dataExtensionsRoute.GET("/:customerKey/DataExtensionFields", func(c *gin.Context) {
			var dataExtensionRequest = &types.DataExtensionRequest{}
			if err := c.ShouldBindUri(&dataExtensionRequest); err != nil {
				c.JSON(400, gin.H{"message": err.Error()})
				return
			}

			// Convert *gin.Context to context.Context
			ctx := c.Request.Context()

			// Call the original function
			var result *soap_to_rest_sfmc.RetrieveDEFieldResponseMsg
			var err error
			result, err = soap_to_rest_sfmc.RetrieveDataExtensionFieldsByDataExtensionCustomerKey(ctx, *dataExtensionRequest)
			if err != nil {
				// Handle error
				errorResponse(c, err)
				return
			}

			// If needed, use ginCtx for further processing or respond directly
			// For example, assuming a successful operation:
			c.JSON(http.StatusOK, gin.H{"message": result})
		})
		dataExtensionsRoute.GET("/:customerKey/DataExtensionFields/:dataExtensionFieldsCustomerKey", func(c *gin.Context) {
			var dataExtensionRequest = &types.DataExtensionRequest{}
			if err := c.ShouldBindUri(&dataExtensionRequest); err != nil {
				c.JSON(400, gin.H{"message": err.Error()})
				return
			}

			fmt.Println("CustomerKey: " + dataExtensionRequest.CustomerKey)
			fmt.Println("CustomerKey: " + dataExtensionRequest.DataExtensionFieldCustomerKey)

			// Convert *gin.Context to context.Context
			ctx := c.Request.Context()

			// Call the original function
			var result *soap_to_rest_sfmc.RetrieveDEFieldResponseMsg
			var err error
			result, err = soap_to_rest_sfmc.RetrieveDataExtensionFieldByDataExtensionCustomerKeyAndFieldCustomerKey(ctx, *dataExtensionRequest)
			if err != nil {
				// Handle error
				errorResponse(c, err)
				return
			}

			// If needed, use ginCtx for further processing or respond directly
			// For example, assuming a successful operation:
			c.JSON(http.StatusOK, gin.H{"message": result})
		})
	}

	// Wrap the call to soap_to_rest_sfmc.GetRecipients with a function that adapts it to Gin's expected handler signature
	router.GET("/Describe/:object", func(c *gin.Context) {
		// Convert *gin.Context to context.Context
		ctx := c.Request.Context()

		// Call the original function
		var result *soap_to_rest_sfmc.DefinitionResponseMsg
		var err error
		result, err = soap_to_rest_sfmc.Describe(ctx)
		if err != nil {
			// Handle error
			errorResponse(c, err)
			return
		}

		// If needed, use ginCtx for further processing or respond directly
		// For example, assuming a successful operation:
		c.JSON(http.StatusOK, gin.H{"message": result})
	})

	// Subscriber methods
	router.GET("/Subscribers", func(c *gin.Context) {
		// Convert *gin.Context to context.Context
		ctx := c.Request.Context()

		// Call the original function
		var result *soap_to_rest_sfmc.RetrieveRecipientResponseMsg
		var err error
		result, err = soap_to_rest_sfmc.GetRecipients(ctx)
		if err != nil {
			// Handle error
			errorResponse(c, err)
			return
		}

		// needs middelwares here
		// ...
		// ...
		// ...
		// ... try to creeate a class for each response struct
		// ... chnage date fields from string to time.Time
		// ...
		// ...
		// ...

		// If needed, use ginCtx for further processing or respond directly
		// For example, assuming a successful operation:
		c.JSON(http.StatusOK, gin.H{"message": result})
	})
	router.GET("/Subscriber/:id", func(c *gin.Context) {
		id := c.Param("id")

		fmt.Println(id)

		// Convert *gin.Context to context.Context
		ctx := c.Request.Context()

		// Call the original function
		var result *soap_to_rest_sfmc.RetrieveRecipientResponseMsg
		var err error
		result, err = soap_to_rest_sfmc.GetRecipient(ctx, id)
		if err != nil {
			// Handle error
			errorResponse(c, err)
			return
		}

		// If needed, use ginCtx for further processing or respond directly
		// For example, assuming a successful operation:
		c.JSON(http.StatusOK, gin.H{"message": result})
	})

	// Data Extension Fields methods
	router.GET("/DataExtensionFields", func(c *gin.Context) {
		// Convert *gin.Context to context.Context
		ctx := c.Request.Context()

		// Call the original function
		var result *soap_to_rest_sfmc.RetrieveDEFieldResponseMsg
		var err error
		result, err = soap_to_rest_sfmc.RetrieveDataExtensionFields(ctx)
		if err != nil {
			// Handle error
			errorResponse(c, err)
			return
		}

		// If needed, use ginCtx for further processing or respond directly
		// For example, assuming a successful operation:
		c.JSON(http.StatusOK, gin.H{"message": result})
	})

	// Data Extension Object methods
	// essayer de modifier ca pour avoir un appel avec ody contenant la liste des champs a requeter
	router.GET("/DataExtensionObject", func(c *gin.Context) {
		// Convert *gin.Context to context.Context
		ctx := c.Request.Context()

		// Call the original function
		var result *soap_to_rest_sfmc.RetrieveDataExtensionObjectResponseMsg
		var err error
		result, err = soap_to_rest_sfmc.RetrieveDataExtensionObject(ctx)
		if err != nil {
			// Handle error
			errorResponse(c, err)
			return
		}

		// If needed, use ginCtx for further processing or respond directly
		// For example, assuming a successful operation:
		c.JSON(http.StatusOK, gin.H{"message": result})
	})

	server.router = router
	return &server
}

func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}

func errorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
	c.Abort()
}
