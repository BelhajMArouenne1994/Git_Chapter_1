package services

import (
	"context"

	models "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/soap_to_rest_sfmc/models"
	client "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/soap_to_rest_sfmc/client"
)

func RetrieveDataExtensionObject(ctx context.Context) (*models.RetrieveDataExtensionObjectResponseMsg, error) {

	// Construct a RetrieveRequestMsg according to the SFMC API requirements
	retrieveRequest := &models.RetrieveRequestMsg{
		RetrieveRequest: &models.RetrieveRequest{
			ObjectType: "DataExtensionObject[67baea65-a4fc-ee11-b875-f40343c95ac8]", // Assuming you want to retrieve subscriber data
			Properties: []string{
				// essayer de trouver une facon pour mettre aussi automatiquement toutes les propriétés de APIObject struct
				"CustomerKey",
			},
		},
	}

	// Set up your request
	// Call the Retrieve method and handle the response
	sfmcClient := client.NewSfmcAuthClient()

	var err error
	var response *models.RetrieveDataExtensionObjectResponseMsg
	// Call the Retrieve method and handle the response
	response, err = sfmcClient.RetrieveDataExtensionObject(retrieveRequest)

	if err != nil {
		return response, err
	}

	return response, nil
}
