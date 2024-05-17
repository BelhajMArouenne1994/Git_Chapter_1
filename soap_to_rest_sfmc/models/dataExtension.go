package models

import (
	"encoding/xml"
)

// SOAP Request Body Definition structs
type RetrieveDEResponseMsg struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI RetrieveResponseMsg"`

	OverallStatus string `xml:"OverallStatus,omitempty"`

	RequestID string `xml:"RequestID,omitempty"`

	Results []*RetrieveDEResponseMsgAPIOBJECT `xml:"Results,omitempty"`
}

type RetrieveDEResponseMsgAPIOBJECT struct {
	*APIObject

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	IsSendable bool `json:"isSendable,omitempty"`

	IsTestable bool `json:"isTestable,omitempty"`

	SendableDataExtensionField *DataExtensionField `json:"sendableDataExtensionField,omitempty"`

	SendableSubscriberField *Attribute `json:"sendableSubscriberField,omitempty"`

	Template *DataExtensionTemplate `json:"template,omitempty"`

	DataRetentionPeriodLength int32 `json:"dataRetentionPeriodLength,omitempty"`

	DataRetentionPeriodUnitOfMeasure int32 `json:"dataRetentionPeriodUnitOfMeasure,omitempty"`

	RowBasedRetention bool `json:"rowBasedRetention,omitempty"`

	ResetRetentionPeriodOnImport bool `json:"resetRetentionPeriodOnImport,omitempty"`

	DeleteAtEndOfRetentionPeriod bool `json:"deleteAtEndOfRetentionPeriod,omitempty"`

	RetainUntil string `json:"retainUntil,omitempty"`

	//Fields struct { Field []*DataExtensionField `xml:"Field,omitempty"`} `xml:"Fields,omitempty"`

	DataRetentionPeriod *DateTimeUnitOfMeasure `json:"dataRetentionPeriod,omitempty"`

	CategoryID int64 `json:"categoryID,omitempty"`

	Status string `json:"status,omitempty"`
}


// SFMC REST API Rerquest structs
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