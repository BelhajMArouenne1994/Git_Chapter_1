package soap_to_rest_sfmc_handlers

import (
	"fmt"
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
