package soap_to_rest_sfmc_handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
