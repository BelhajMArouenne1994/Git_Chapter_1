package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Define soap_to_rest_sfmc Types
type DataExtensionRequest struct {
	//ID          string `uri:"id" binding:"required,uuid"`
	CustomerKey                   string `uri:"customerKey"`                    //binding:"required
	DataExtensionFieldCustomerKey string `uri:"dataExtensionFieldsCustomerKey"` //binding:"required
}

type DataExtensionFieldsRequest struct {
	//ID          string `uri:"id" binding:"required,uuid"`
	CustomerKey                   string `uri:"customerKey"`                    //binding:"required
	DataExtensionFieldCustomerKey string `uri:"dataExtensionFieldsCustomerKey"` //binding:"required
}

// Define mongo_db Types
type DataExtensionMongoDB struct {
	ID                         primitive.ObjectID          `bson:"_id,omitempty"`
	ObjectID                   string                      `bson:"objectId"`
	Client                     Client                      `bson:"client"`
	CreatedDate                time.Time                   `bson:"createdDate"`
	ModifiedDate               time.Time                   `bson:"modifiedDate"`
	CustomerKey                string                      `bson:"customerKey"`
	Name                       string                      `bson:"name"`
	IsSendable                 bool                        `bson:"isSendable"`
	SendableDataExtensionField DataExtensionFieldMongoDB   `bson:"sendableDataExtensionField"`
	SendableSubscriberField    SubscriberFieldMongoDB      `bson:"sendableSubscriberField"`
	CategoryID                 int                         `bson:"categoryID"`
	Status                     string                      `bson:"status"`
	Fields                     []DataExtensionFieldMongoDB `bson:"fields"`
}

type SubscriberFieldMongoDB struct {
	Name        string  `bson:"Name"`
	Value       string  `bson:"Value"`
	Compression *string `bson:"Compression"`
}

type DataExtensionFieldMongoDB struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	ObjectID     string             `bson:"objectId"`
	CustomerKey  string             `bson:"customerKey"`
	Name         string             `bson:"name"`
	DataType     string             `bson:"dataType"`
	IsPrimaryKey bool               `bson:"isPrimaryKey"`
	FieldType    string             `bson:"fieldType"`
	// ................
	// ................
	// ................
	// ................
	// ................
	// ................
}
