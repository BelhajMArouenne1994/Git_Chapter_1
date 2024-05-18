package models

import (
	"encoding/xml"
)

type DataExtension struct {
	*APIObject

	Name string `json:"name,omitempty" xml:"Name,omitempty"`

	Description string `json:"description,omitempty" xml:"Description,omitempty"`

	IsSendable bool `json:"isSendable,omitempty" xml:"IsSendable,omitempty"`

	IsTestable bool `json:"isTestable,omitempty" xml:"IsTestable,omitempty"`

	SendableDataExtensionField *DataExtensionField `json:"sendableDataExtensionField,omitempty" xml:"SendableDataExtensionField,omitempty"`

	SendableSubscriberField *Attribute `json:"sendableSubscriberField,omitempty" xml:"SendableSubscriberField,omitempty"`

	Template *DataExtensionTemplate `json:"template,omitempty" xml:"Template,omitempty"`

	DataRetentionPeriodLength int32 `json:"dataRetentionPeriodLength,omitempty" xml:"DataRetentionPeriodLength,omitempty"`

	DataRetentionPeriodUnitOfMeasure int32 `json:"dataRetentionPeriodUnitOfMeasure,omitempty" xml:"DataRetentionPeriodUnitOfMeasure,omitempty"`

	RowBasedRetention bool `json:"rowBasedRetention,omitempty" xml:"RowBasedRetention,omitempty"`

	ResetRetentionPeriodOnImport bool `json:"resetRetentionPeriodOnImport,omitempty" xml:"ResetRetentionPeriodOnImport,omitempty"`

	DeleteAtEndOfRetentionPeriod bool `json:"deleteAtEndOfRetentionPeriod,omitempty" xml:"DeleteAtEndOfRetentionPeriod,omitempty"`

	RetainUntil string `json:"retainUntil,omitempty" xml:"RetainUntil,omitempty"`

	DataRetentionPeriod *DateTimeUnitOfMeasure `json:"dataRetentionPeriod,omitempty" xml:"DataRetentionPeriod,omitempty"`

	CategoryID int64 `json:"categoryID,omitempty" xml:"CategoryID,omitempty"`

	Status string `json:"status,omitempty" xml:"Status,omitempty"`
}


type DataExtensionCreate struct {

	XMLName xml.Name `xml:"Objects"`

	XSIType string `xml:"xsi:type,attr,omitempty"`
	
	Client *ClientID `json:"client,omitempty" xml:"Client,omitempty"`

	Name string `json:"name,omitempty" xml:"Name,omitempty"`

	Description string `json:"description,omitempty" xml:"Description,omitempty"`

	IsSendable bool `json:"isSendable,omitempty" xml:"IsSendable,omitempty"`

	IsTestable bool `json:"isTestable,omitempty" xml:"IsTestable,omitempty"`

	SendableDataExtensionField *DataExtensionField `json:"sendableDataExtensionField,omitempty" xml:"SendableDataExtensionField,omitempty"`

	SendableSubscriberField *Attribute `json:"sendableSubscriberField,omitempty" xml:"SendableSubscriberField,omitempty"`

	Template *DataExtensionTemplate `json:"template,omitempty" xml:"Template,omitempty"`

	DataRetentionPeriodLength int32 `json:"dataRetentionPeriodLength,omitempty" xml:"DataRetentionPeriodLength,omitempty"`

	DataRetentionPeriodUnitOfMeasure int32 `json:"dataRetentionPeriodUnitOfMeasure,omitempty" xml:"DataRetentionPeriodUnitOfMeasure,omitempty"`

	RowBasedRetention bool `json:"rowBasedRetention,omitempty" xml:"RowBasedRetention,omitempty"`

	ResetRetentionPeriodOnImport bool `json:"resetRetentionPeriodOnImport,omitempty" xml:"ResetRetentionPeriodOnImport,omitempty"`

	DeleteAtEndOfRetentionPeriod bool `json:"deleteAtEndOfRetentionPeriod,omitempty" xml:"DeleteAtEndOfRetentionPeriod,omitempty"`

	RetainUntil string `json:"retainUntil,omitempty" xml:"RetainUntil,omitempty"`

	DataRetentionPeriod *DateTimeUnitOfMeasure `json:"dataRetentionPeriod,omitempty" xml:"DataRetentionPeriod,omitempty"`

	CategoryID int64 `json:"categoryID,omitempty" xml:"CategoryID,omitempty"`

	Status string `json:"status,omitempty" xml:"Status,omitempty"`
}

type RetrieveDEResponseMsg struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI RetrieveResponseMsg"`

	OverallStatus string `xml:"OverallStatus,omitempty"`

	RequestID string `xml:"RequestID,omitempty"`

	Results []*DataExtension `xml:"Results,omitempty"`
}

type CreateDataExtensionRequest struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI CreateRequest"`

	Options *CreateOptions `json:"options,omitempty" xml:"Options,omitempty"`

	Objects *DataExtensionCreate `json:"objects,omitempty" xml:"Objects,omitempty"`
}

type CreateDataExtensionResponse struct {
	XMLName xml.Name `json:"-" xml:"http://exacttarget.com/wsdl/partnerAPI CreateResponse"`

	Results []*CreateResult `json:"results,omitempty" xml:"Results,omitempty"`

	RequestID string `json:"requestID,omitempty" xml:"RequestID,omitempty"`

	OverallStatus string `json:"overallStatus,omitempty" xml:"OverallStatus,omitempty"`
}

// SFMC REST API Rerquest structs
type DataExtensionUriRequest struct {
	//ID          string `uri:"id" binding:"required,uuid"`
	CustomerKey                   string `uri:"customerKey"`                    //binding:"required
	DataExtensionFieldCustomerKey string `uri:"dataExtensionFieldsCustomerKey"` //binding:"required
}

type DataExtensionFieldsUriRequest struct {
	//ID          string `uri:"id" binding:"required,uuid"`
	CustomerKey                   string `uri:"customerKey"`                    //binding:"required
	DataExtensionFieldCustomerKey string `uri:"dataExtensionFieldsCustomerKey"` //binding:"required
}
