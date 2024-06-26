package handlers

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/BelhajMArouenne1994/GIT_CHAPTER_1/soap_to_rest_sfmc/models"
	"github.com/BelhajMArouenne1994/GIT_CHAPTER_1/soap_to_rest_sfmc/services"

	mongo_db_Service "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/mongo_db/services"
	transform_services "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/transform"

	"go.mongodb.org/mongo-driver/mongo"
)

// HttpError structure to represent an HTTP error
type HttpError struct {
	Description string `json:"description,omitempty"`
	Metadata    string `json:"metadata,omitempty"`
	StatusCode  int    `json:"statusCode"`
}

// DataExtensionError structure embedding HttpError
type DataExtensionError struct {
	*HttpError
}

// Implementing the error interface for HttpError
func (e HttpError) Error() string {
	return fmt.Sprintf("description: %s, metadata: %s", e.Description, e.Metadata)
}

// CreateNewHttpError creates a new instance of HttpError
func CreateNewHttpError(description, metadata string, statusCode int) *HttpError {
	return &HttpError{
		Description: description,
		Metadata:    metadata,
		StatusCode:  statusCode,
	}
}

// Assuming mongoService is initialized and passed to the handler
var mongoService *mongo_db_Service.MongoService

func InitMongoService(client *mongo.Client, dbName string) {
	mongoService = mongo_db_Service.NewMongoService(client, dbName)
}

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
				Err:  CreateNewHttpError("Unauthorized", "Missing Authorization Header", http.StatusUnauthorized),
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
				Err:  CreateNewHttpError("Unsupported Media Type", "Invalid Content Type", http.StatusUnsupportedMediaType),
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
			case *HttpError:
				c.AbortWithStatusJSON(e.StatusCode, e)
			default:
				// If the error is not an HttpError, respond with a generic 500 status code
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
			}
		}
	}
}

func RetrieveDataExtensionsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		result, err := services.RetrieveDataExtensions(ctx)
		if err != nil {
			err := &gin.Error{
				Err:  CreateNewHttpError(err.Error(), err.Error(), 400),
				Type: gin.ErrorTypePublic,
			}
			c.Errors = append(c.Errors, err)
			c.Abort() // Abort the context if the header is not application/json			return
		}

		// You can now call MongoDB service to store the retrieved data
		//mongoService.StoreDataExtension(result)

		c.JSON(http.StatusOK, gin.H{"message": result})
	}
}

func RetrieveDataExtensionByCustomerKeyHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataExtensionRequest models.DataExtensionUriRequest
		if err := c.ShouldBindUri(&dataExtensionRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ctx := c.Request.Context()

		sfmcResponse, err := services.RetrieveDataExtensionByCustomerKey(ctx, dataExtensionRequest)
		if err != nil {
			err := &gin.Error{
				Err:  CreateNewHttpError(err.Error(), err.Error(), 400),
				Type: gin.ErrorTypePublic,
			}
			c.Errors = append(c.Errors, err)
			c.Abort() // Abort the context if the header is not application/json			return
		}

		transformedData, err := transform_services.TransformToDataExtensionMongoDB(sfmcResponse)
		if err != nil {
			err := &gin.Error{
				Err:  CreateNewHttpError(err.Error(), err.Error(), 400),
				Type: gin.ErrorTypePublic,
			}
			c.Errors = append(c.Errors, err)
			c.Abort() // Abort the context if the header is not application/json			return
		}

		// Store the transformed data in MongoDB
		for _, data := range transformedData {
			result, err := mongoService.StoreDataExtension(ctx, data)
			if err != nil {
				err := &gin.Error{
					Err:  CreateNewHttpError(err.Error(), err.Error(), 400),
					Type: gin.ErrorTypePublic,
				}
				c.Errors = append(c.Errors, err)
				c.Abort() // Abort the context if the header is not application/json			return
			}
			c.JSON(http.StatusOK, gin.H{"message": result.InsertedID})
		}
	}
}

func CreateDataExtensionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//var request models.CreateDataExtensionRequest
		var requestBody models.DataExtensionCreate
	
		 
		// Lier le corps de la requête JSON à la structure Go
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	
		requestBody.XSIType = "DataExtension"

		request := &models.CreateDataExtensionRequest{
			XMLName: xml.Name{Local: "CreateRequest"},
			Options: &models.CreateOptions{},
			Objects: &requestBody,
		}


		// Appeler le service pour créer la Data Extension dans SFMC
		response, err := services.CreateDataExtension(context.Background(), request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	
		// Retourner la réponse au client
		c.JSON(http.StatusOK, response)
	}
}

func RetrieveDataExtensionFieldByCustomerKeyAndFieldKeyHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataExtensionRequest models.DataExtensionUriRequest
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
