package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	client "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/mongo_db/client"
	models "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/mongo_db/models"
)
// Handler chains together multiple gin.HandlerFunc
func Handler(handlers ...gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, handler := range handlers {
			handler(c)
			// If one handler calls c.Abort(), stop the chain
			if c.IsAborted() {
				return
			}
		}
	}
}

// Additional middleware to demonstrate chaining
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			err := &gin.Error{
				Err:  models.CreateNewHttpError("Unauthorized", "Missing Authorization Header", http.StatusUnauthorized),
				Type: gin.ErrorTypePublic,
			}
			c.Errors = append(c.Errors, err)
			c.Abort() // Abort the context if the header is not present
		} else {
			c.Next() // Continue to the next middleware/handler if the header is correct
		}
	}
}

// HeaderHandler checks the Content-Type of the request and adds an error if it's not application/json
func HeaderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestHeader := c.Request.Header.Get("Content-Type")
		if requestHeader != "application/json" {
			err := &gin.Error{
				Err:  models.CreateNewHttpError("Unsupported Media Type", "Invalid Content Type", http.StatusUnsupportedMediaType),
				Type: gin.ErrorTypePublic,
			}
			c.Errors = append(c.Errors, err)
			c.Abort() // Abort the context if the header is not application/json
		} else {
			c.Next() // Continue to the next middleware/handler if the header is correct
		}
	}
}

// ErrorHandler handles any errors added to the context and responds with the appropriate message and status code
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Process request, then handle errors

		if len(c.Errors) > 0 {
			// Handle the first error in the list
			err := c.Errors[0]
			switch e := err.Err.(type) {
			case *models.HttpError:
				c.AbortWithStatusJSON(e.StatusCode, e)
			default:
				// If the error is not an HttpError, respond with a generic 500 status code
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
			}
		}
	}
}

func CreateDataExtension() gin.HandlerFunc {
	return func(c *gin.Context) {
		newDataExtension := models.DataExtensionMongoDB{}

		if err := c.ShouldBindBodyWith(&newDataExtension, binding.JSON); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		// Connexion à la base de données
		client, ctx, cancel := client.ConnectDB()
		defer func() {
			cancel()
			client.Disconnect(ctx)
		}()

		collDE := client.Database("sfmcMarouenne").Collection("DataExtension")
		result, err := collDE.InsertOne(ctx, newDataExtension)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			//c.Next()
			return
		}

		collDEFields := client.Database("sfmcMarouenne").Collection("DataExtensionFields")
		newfields := make([]interface{}, len(newDataExtension.Fields))
		if len(newDataExtension.Fields) > 0 {
			for i := range newDataExtension.Fields {
				newfields[i] = newDataExtension.Fields[i]
			}

			results, err := collDEFields.InsertMany(ctx, newfields)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Data extension & related Fields created", "DataExtensionID": result.InsertedID, "DataExtensionFieldsIDs": results.InsertedIDs})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Data extension created", "insertedID": result.InsertedID})
		return
	}
}
