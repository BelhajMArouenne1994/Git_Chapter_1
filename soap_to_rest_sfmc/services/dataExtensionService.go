package services

import (
	"context"

	models "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/soap_to_rest_sfmc/models"
	client "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/soap_to_rest_sfmc/client"
)

func RetrieveDataExtensions(ctx context.Context, optionalArgs ...string) (*models.RetrieveDEResponseMsg, error) {
	// optional arguments
	// first option : requestID, this value enable to get remaining result from previous request
	var requestID string
	if len(optionalArgs) > 0 {
		requestID = optionalArgs[0] // Use the first optional argument as the filter
	}

	// Construct a RetrieveRequestMsg according to the SFMC API requirements
	retrieveRequest := &models.RetrieveRequestMsg{
		RetrieveRequest: &models.RetrieveRequest{
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
	sfmcClient := client.NewSfmcAuthClient()

	var err error
	var response *models.RetrieveDEResponseMsg
	// Call the Retrieve method and handle the response
	response, err = sfmcClient.RetrieveDataExtensions(retrieveRequest)

	if err != nil {
		return response, err
	}

	return response, nil
}

func RetrieveDataExtensionByCustomerKey(ctx context.Context, dataExtensionCustomerKey models.DataExtensionRequest) (*models.RetrieveDEResponseMsg, error) {

	// Construct a RetrieveRequestMsg according to the SFMC API requirements
	retrieveRequest := &models.RetrieveRequestMsg{
		RetrieveRequest: &models.RetrieveRequest{
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
			Filter: &models.SimpleFilterPart{
				XSIType:        "SimpleFilterPart",
				Property:       "CustomerKey",
				SimpleOperator: "equals",
				Value:          []string{dataExtensionCustomerKey.CustomerKey},
			},
		},
	}

	// Set up your request
	// Call the Retrieve method and handle the response
	sfmcClient := client.NewSfmcAuthClient()

	var err error
	var response *models.RetrieveDEResponseMsg
	// Call the Retrieve method and handle the response
	response, err = sfmcClient.RetrieveDataExtensions(retrieveRequest)

	if err != nil {
		return response, err
	}

	return response, nil
}


