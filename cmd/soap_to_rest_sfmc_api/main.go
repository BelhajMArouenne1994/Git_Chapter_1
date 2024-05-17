package soap_to_rest_sfmc_api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	models "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/soap_to_rest_sfmc/models"
	services "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/soap_to_rest_sfmc/services"
	handlers "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/soap_to_rest_sfmc/handlers"

	//mongo_db_models "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/mongo_db/models"
	//mongo_db_services "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/mongo_db/services"
	mongo_db_handlers "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/mongo_db/handlers"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {

	server := Server{}

	router := gin.New()
	// Utilisation des middlewares
	router.Use(handlers.ErrorHandler())

	dataExtensionsRoute := router.Group("/DataExtensions")
	{
		// Utilisation de Use pour appliquer les middlewares
		dataExtensionsRoute.Use(
			//handlers.HeaderHandler(),
			handlers.AuthMiddleware(),
		)
		// Data Extension methods
		dataExtensionsRoute.POST("/", mongo_db_handlers.CreateDataExtension())


		dataExtensionsRoute.GET("/", func(c *gin.Context) {
			// Convert *gin.Context to context.Context
			ctx := c.Request.Context()

			// Call the original function
			var result *models.RetrieveDEResponseMsg
			var err error
			result, err = services.RetrieveDataExtensions(ctx)
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
			var dataExtensionRequest = &models.DataExtensionRequest{}
			if err := c.ShouldBindUri(dataExtensionRequest); err != nil {
				c.JSON(400, gin.H{"message": err.Error()})
				return
			}

			// Convert *gin.Context to context.Context
			ctx := c.Request.Context()

			// Call the original function
			var result *models.RetrieveDEResponseMsg
			var err error
			result, err = services.RetrieveDataExtensionByCustomerKey(ctx, *dataExtensionRequest)
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
			var dataExtensionRequest = &models.DataExtensionRequest{}
			if err := c.ShouldBindUri(&dataExtensionRequest); err != nil {
				c.JSON(400, gin.H{"message": err.Error()})
				return
			}

			// Convert *gin.Context to context.Context
			ctx := c.Request.Context()

			// Call the original function
			var result *models.RetrieveDEFieldResponseMsg
			var err error
			result, err = services.RetrieveDataExtensionFieldsByDataExtensionCustomerKey(ctx, *dataExtensionRequest)
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
			var dataExtensionRequest = &models.DataExtensionRequest{}
			if err := c.ShouldBindUri(&dataExtensionRequest); err != nil {
				c.JSON(400, gin.H{"message": err.Error()})
				return
			}

			fmt.Println("CustomerKey: " + dataExtensionRequest.CustomerKey)
			fmt.Println("CustomerKey: " + dataExtensionRequest.DataExtensionFieldCustomerKey)

			// Convert *gin.Context to context.Context
			ctx := c.Request.Context()

			// Call the original function
			var result *models.RetrieveDEFieldResponseMsg
			var err error
			result, err = services.RetrieveDataExtensionFieldByDataExtensionCustomerKeyAndFieldCustomerKey(ctx, *dataExtensionRequest)
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

	router.GET("/Describe/:object", func(c *gin.Context) {
		// Convert *gin.Context to context.Context
		ctx := c.Request.Context()

		// Call the original function
		var result *models.DefinitionResponseMsg
		var err error
		result, err = services.Describe(ctx)
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
		var result *models.RetrieveRecipientResponseMsg
		var err error
		result, err = services.GetRecipients(ctx)
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
		var result *models.RetrieveRecipientResponseMsg
		var err error
		result, err = services.GetRecipient(ctx, id)
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
		var result *models.RetrieveDEFieldResponseMsg
		var err error
		result, err = services.RetrieveDataExtensionFields(ctx)
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
		var result *models.RetrieveDataExtensionObjectResponseMsg
		var err error
		result, err = services.RetrieveDataExtensionObject(ctx)
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
