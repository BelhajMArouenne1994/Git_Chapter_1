package services

import (
	"context"
	"encoding/xml"

	models "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/soap_to_rest_sfmc/models"
	client "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/soap_to_rest_sfmc/client"
)

func RetrieveDataExtensionFields(ctx context.Context) (*models.RetrieveDEFieldResponseMsg, error) {

	// Construct a RetrieveRequestMsg according to the SFMC API requirements
	retrieveRequest := &models.RetrieveRequestMsg{
		RetrieveRequest: &models.RetrieveRequest{
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
	sfmcClient := client.NewSfmcAuthClient()

	var err error
	var response *models.RetrieveDEFieldResponseMsg
	// Call the Retrieve method and handle the response
	response, err = sfmcClient.RetrieveDataExtensionFields(retrieveRequest)

	if err != nil {
		return response, err
	}

	return response, nil
}

func RetrieveDataExtensionFieldsByDataExtensionCustomerKey(ctx context.Context, dataExtensionCustomerKey models.DataExtensionRequest) (*models.RetrieveDEFieldResponseMsg, error) {
	// Construct a RetrieveRequestMsg according to the SFMC API requirements
	retrieveRequest := &models.RetrieveRequestMsg{
		RetrieveRequest: &models.RetrieveRequest{
			ObjectType: "DataExtensionField", // Assuming you want to retrieve subscriber data
			Properties: []string{
				"ObjectID", "PartnerKey", "CustomerKey", "Name", "DefaultValue",
				"Ordinal", "IsPrimaryKey", "FieldType",
				"CreatedDate", "ModifiedDate",
				"Client.ID",
				"DataExtension.CustomerKey",
			},
			Filter: &models.SimpleFilterPart{
				XMLName:        xml.Name{Local: "Filter"},
				XSIType:        "SimpleFilterPart",
				Property:       "DataExtension.CustomerKey",
				SimpleOperator: "equals",
				Value:          []string{dataExtensionCustomerKey.CustomerKey},
			},
		},
	}

	// Set up your request
	// Call the Retrieve method and handle the response
	sfmcClient := client.NewSfmcAuthClient()

	var err error
	var response *models.RetrieveDEFieldResponseMsg
	// Call the Retrieve method and handle the response
	response, err = sfmcClient.RetrieveDataExtensionFields(retrieveRequest)

	if err != nil {
		return response, err
	}

	return response, nil
}

// TO DO

func RetrieveDataExtensionFieldByDataExtensionCustomerKeyAndFieldCustomerKey(ctx context.Context, dataExtensionRequest models.DataExtensionRequest) (*models.RetrieveDEFieldResponseMsg, error) {

	// Construct a RetrieveRequestMsg according to the SFMC API requirements
	retrieveRequest := &models.RetrieveRequestMsg{
		RetrieveRequest: &models.RetrieveRequest{
			ObjectType: "DataExtensionField", // Assuming you want to retrieve subscriber data
			Properties: []string{
				// essayer de trouver une facon pour mettre aussi automatiquement toutes les propriétés de APIObject struct
				"ObjectID", "PartnerKey", "CustomerKey", "Name", "DefaultValue",
				"Ordinal", "IsPrimaryKey", "FieldType",
				"CreatedDate", "ModifiedDate",
				"Client.ID",
				"DataExtension.CustomerKey",
			},
			Filter: &models.ComplexFilterPart{
				XMLName: xml.Name{Local: "Filter"},
				XSIType: "ComplexFilterPart",
				LeftOperand: &models.SimpleFilterPart{
					XMLName:        xml.Name{Space: "http://exacttarget.com/wsdl/partnerAPI", Local: "LeftOperand"},
					XSIType:        "SimpleFilterPart",
					Property:       "DataExtension.CustomerKey",
					SimpleOperator: "equals",
					Value:          []string{dataExtensionRequest.CustomerKey},
				},
				LogicalOperator: models.LogicalOperatorAND,
				RightOperand: &models.SimpleFilterPart{
					XMLName:        xml.Name{Space: "http://exacttarget.com/wsdl/partnerAPI", Local: "RightOperand"},
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
	sfmcClient := client.NewSfmcAuthClient()

	var err error
	var response *models.RetrieveDEFieldResponseMsg
	// Call the Retrieve method and handle the response
	response, err = sfmcClient.RetrieveDataExtensionFields(retrieveRequest)

	if err != nil {
		return response, err
	}

	return response, nil
}

