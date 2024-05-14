package soap_to_rest_sfmc

import (
	"context"
	"encoding/xml"

	types "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/types"
	//"fmt"
	//"log"
	//"net/http"
	//"github.com/gin-gonic/gin"
)

//needs to set the needed fields

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

func RetrieveDataExtensions(ctx context.Context, optionalArgs ...string) (*RetrieveDEResponseMsg, error) {
	// optional arguments
	// first option : requestID, this value enable to get remaining result from previous request
	var requestID string
	if len(optionalArgs) > 0 {
		requestID = optionalArgs[0] // Use the first optional argument as the filter
	}

	// Construct a RetrieveRequestMsg according to the SFMC API requirements
	retrieveRequest := &RetrieveRequestMsg{
		RetrieveRequest: &RetrieveRequest{
			ObjectType:      "DataExtension",
			ContinueRequest: requestID,
			Properties: []string{

				// essayer de trouver une facon pour mettre aussi automatiquement toutes les propriétés de APIObject struct
				"ObjectID", "CustomerKey", "Name", "Client.ID", "Description",
				"CreatedDate", "ModifiedDate",
				"IsSendable", "IsTestable", "SendableDataExtensionField.Name",
				"SendableSubscriberField.Name", "Template.CustomerKey",
				"IsPlatformObject",
				"DataRetentionPeriodLength", "DataRetentionPeriodUnitOfMeasure",
				"RowBasedRetention", "ResetRetentionPeriodOnImport",
				"DeleteAtEndOfRetentionPeriod", "RetainUntil",
				"DataRetentionPeriod",
				"CategoryID", "Status",
			},
		},
	}

	// Set up your request
	// Call the Retrieve method and handle the response
	sfmcClient := NewSfmcAuthClient()

	var err error
	var response *RetrieveDEResponseMsg
	// Call the Retrieve method and handle the response
	response, err = sfmcClient.RetrieveDataExtensions(retrieveRequest)

	if err != nil {
		return response, err
	}

	return response, nil
}

func RetrieveDataExtensionByCustomerKey(ctx context.Context, dataExtensionCustomerKey types.DataExtensionRequest) (*RetrieveDEResponseMsg, error) {

	// Construct a RetrieveRequestMsg according to the SFMC API requirements
	retrieveRequest := &RetrieveRequestMsg{
		RetrieveRequest: &RetrieveRequest{
			ObjectType: "DataExtension", // Assuming you want to retrieve subscriber data
			Properties: []string{
				// essayer de trouver une facon pour mettre aussi automatiquement toutes les propriétés de APIObject struct
				"ObjectID", "CustomerKey", "Name", "Client.ID", "Description",
				"CreatedDate", "ModifiedDate",
				"IsSendable", "IsTestable", "SendableDataExtensionField.Name",
				"SendableSubscriberField.Name", "Template.CustomerKey",
				"Status", "IsPlatformObject",
				"DataRetentionPeriodLength", "DataRetentionPeriodUnitOfMeasure",
				"RowBasedRetention", "ResetRetentionPeriodOnImport",
				"DeleteAtEndOfRetentionPeriod", "RetainUntil",
				"DataRetentionPeriod",
			},
			Filter: &SimpleFilterPart{
				XSIType:        "SimpleFilterPart",
				Property:       "CustomerKey",
				SimpleOperator: "equals",
				Value:          []string{dataExtensionCustomerKey.CustomerKey},
			},
		},
	}

	// Set up your request
	// Call the Retrieve method and handle the response
	sfmcClient := NewSfmcAuthClient()

	var err error
	var response *RetrieveDEResponseMsg
	// Call the Retrieve method and handle the response
	response, err = sfmcClient.RetrieveDataExtensions(retrieveRequest)

	if err != nil {
		return response, err
	}

	return response, nil
}
