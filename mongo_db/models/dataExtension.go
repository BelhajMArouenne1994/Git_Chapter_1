package models

import (
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
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



// MongoDB structs
type DataExtensionMongoDB struct {
	ID                         primitive.ObjectID          `bson:"_id,omitempty" json:"_id,omitempty"`
	ObjectID                   string                      `bson:"objectId" json:"objectId" binding:"required"`
	Client                     ClientnMongoDB                      `bson:"client" json:"client"`
	CreatedDate                time.Time                   `bson:"createdDate" json:"createdDate" binding:"required"`
	ModifiedDate               time.Time                   `bson:"modifiedDate" json:"modifiedDate"`
	CustomerKey                string                      `bson:"customerKey" json:"customerKey" binding:"required"`
	Name                       string                      `bson:"name" json:"name"`
	IsSendable                 bool                        `bson:"isSendable" json:"isSendable"`
	SendableDataExtensionField DataExtensionFieldMongoDB   `bson:"sendableDataExtensionField" json:"sendableDataExtensionField"`
	SendableSubscriberField    SubscriberFieldMongoDB      `bson:"sendableSubscriberField" json:"sendableSubscriberField"`
	CategoryID                 int                         `bson:"categoryID" json:"categoryID"`
	Status                     string                      `bson:"status" json:"status"`
	Fields                     []DataExtensionFieldMongoDB `bson:"fields" json:"fields"`
}

type SubscriberFieldMongoDB struct {
	Name        string  `bson:"Name"`
	Value       string  `bson:"Value"`
	Compression *string `bson:"Compression"`
}

type DataExtensionFieldMongoDB struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ObjectID     string             `bson:"objectId" json:"objectId"`
	CustomerKey  string             `bson:"customerKey" json:"customerKey"`
	Name         string             `bson:"name" json:"name"`
	DataType     string             `bson:"dataType" json:"dataType"`
	IsPrimaryKey bool               `bson:"isPrimaryKey" json:"isPrimaryKey"`
	FieldType    string             `bson:"fieldType" json:"fieldType"`
	// ................
	// ................
	// ................
	// ................
	// ................
	// ................
}
