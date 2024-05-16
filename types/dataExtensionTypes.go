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
	ID                         primitive.ObjectID          `bson:"_id,omitempty" json:"_id,omitempty"`
	ObjectID                   string                      `bson:"objectId" json:"objectId, required"`
	Client                     Client                      `bson:"client" json:"client"`
	CreatedDate                time.Time                   `bson:"createdDate" json:"createdDate"`
	ModifiedDate               time.Time                   `bson:"modifiedDate" json:"modifiedDate"`
	CustomerKey                string                      `bson:"customerKey" json:"customerKey"`
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
