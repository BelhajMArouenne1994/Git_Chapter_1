package soap_sfmc

import (
	"context"
	"encoding/xml"

	types "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/types"
	//"fmt"
	//"log"
	//"net/http"
	//"github.com/gin-gonic/gin"
)

type RetrieveDEFieldResponseMsg struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI RetrieveResponseMsg"`

	OverallStatus string `xml:"OverallStatus,omitempty"`

	RequestID string `xml:"RequestID,omitempty"`

	Results []*RetrieveDEFieldResponseMsgAPIOBJECT `xml:"Results,omitempty"`
}

type RetrieveDEFieldResponseMsgAPIOBJECT struct {
	*PropertyDefinition

	Ordinal int32 `json:"ordinal,omitempty"`

	IsPrimaryKey bool `json:"isPrimaryKey,omitempty"`

	FieldType *DataExtensionFieldType `json:"fieldType,omitempty"`

	DataExtension *DataExtension `json:"dataExtension,omitempty"`

	StorageType *DataExtensionFieldStorageType `json:"storageType,omitempty"`
}

func RetrieveDataExtensionFields(ctx context.Context) (*RetrieveDEFieldResponseMsg, error) {

	// Construct a RetrieveRequestMsg according to the SFMC API requirements
	retrieveRequest := &RetrieveRequestMsg{
		RetrieveRequest: &RetrieveRequest{
			ObjectType: "DataExtensionField", // Assuming you want to retrieve subscriber data
			Properties: []string{
				"ObjectID", "PartnerKey", "CustomerKey", "Name", "DefaultValue",
				"Ordinal", "IsPrimaryKey", "FieldType",
				"CreatedDate", "ModifiedDate",
				"Client.ID",
				"DataExtension.CustomerKey",

				// ZID CHOUFHOM
				//"Name": "StorageType", 			"DataType": "DataExtensionFieldStorageType",
				//"Name": "DataExtension", 		"DataType": "DataExtension",

				// Attributes related to PropertyDefinition
				//"MaxLength", "IsRequired", "Scale", "DataType",
				//"IsCreatable", "IsUpdatable", "IsRetrievable", "IsQueryable", "IsFilterable",
			},
		},
	}

	// Set up your request
	// Call the Retrieve method and handle the response
	sfmcClient := NewSfmcAuthClient()

	var err error
	var response *RetrieveDEFieldResponseMsg
	// Call the Retrieve method and handle the response
	response, err = sfmcClient.RetrieveDataExtensionFields(retrieveRequest)

	if err != nil {
		return response, err
	}

	return response, nil
}

func RetrieveDataExtensionFieldsByDataExtensionCustomerKey(ctx context.Context, dataExtensionCustomerKey types.DataExtensionRequest) (*RetrieveDEFieldResponseMsg, error) {
	// Construct a RetrieveRequestMsg according to the SFMC API requirements
	retrieveRequest := &RetrieveRequestMsg{
		RetrieveRequest: &RetrieveRequest{
			ObjectType: "DataExtensionField", // Assuming you want to retrieve subscriber data
			Properties: []string{
				"ObjectID", "PartnerKey", "CustomerKey", "Name", "DefaultValue",
				"Ordinal", "IsPrimaryKey", "FieldType",
				"CreatedDate", "ModifiedDate",
				"Client.ID",
				"DataExtension.CustomerKey",
			},
			Filter: &SimpleFilterPart{
				XMLName: xml.Name{Local: "Filter"},
				XSIType:        "SimpleFilterPart",
				Property:       "DataExtension.CustomerKey",
				SimpleOperator: "equals",
				Value:          []string{dataExtensionCustomerKey.CustomerKey},
			},
		},
	}

	// Set up your request
	// Call the Retrieve method and handle the response
	sfmcClient := NewSfmcAuthClient()

	var err error
	var response *RetrieveDEFieldResponseMsg
	// Call the Retrieve method and handle the response
	response, err = sfmcClient.RetrieveDataExtensionFields(retrieveRequest)

	if err != nil {
		return response, err
	}

	return response, nil
}

// TO DO

func RetrieveDataExtensionFieldByDataExtensionCustomerKeyAndFieldCustomerKey(ctx context.Context, dataExtensionRequest types.DataExtensionRequest) (*RetrieveDEFieldResponseMsg, error) {

	// Construct a RetrieveRequestMsg according to the SFMC API requirements
	retrieveRequest := &RetrieveRequestMsg{
		RetrieveRequest: &RetrieveRequest{
			ObjectType: "DataExtensionField", // Assuming you want to retrieve subscriber data
			Properties: []string{
				// essayer de trouver une facon pour mettre aussi automatiquement toutes les propriétés de APIObject struct
				"ObjectID", "PartnerKey", "CustomerKey", "Name", "DefaultValue",
				"Ordinal", "IsPrimaryKey", "FieldType",
				"CreatedDate", "ModifiedDate",
				"Client.ID",
				"DataExtension.CustomerKey",
			},
			Filter: &ComplexFilterPart{
				XMLName: xml.Name{Local: "Filter"},
				XSIType: "ComplexFilterPart",
				LeftOperand:  &SimpleFilterPart{
					XMLName: xml.Name{Space: "http://exacttarget.com/wsdl/partnerAPI", Local: "LeftOperand"},
					XSIType:        "SimpleFilterPart",
					Property:       "DataExtension.CustomerKey",
					SimpleOperator: "equals",
					Value:          []string{dataExtensionRequest.CustomerKey},
				},
				LogicalOperator: LogicalOperatorAND,
				RightOperand: &SimpleFilterPart{
					XMLName: xml.Name{Space: "http://exacttarget.com/wsdl/partnerAPI", Local: "RightOperand"},
					XSIType:        "SimpleFilterPart",
					Property:       "CustomerKey",
					SimpleOperator: "equals",
					Value:          []string{dataExtensionRequest.DataExtensionFieldCustomerKey},
				},
			},
		},
	}

	// Set up your request
	// Call the Retrieve method and handle the response
	sfmcClient := NewSfmcAuthClient()

	var err error
	var response *RetrieveDEFieldResponseMsg
	// Call the Retrieve method and handle the response
	response, err = sfmcClient.RetrieveDataExtensionFields(retrieveRequest)

	if err != nil {
		return response, err
	}

	return response, nil
}
