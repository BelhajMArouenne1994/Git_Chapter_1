package services

import (
	"context"
	"fmt"

	models "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/soap_to_rest_sfmc/models"
	client "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/soap_to_rest_sfmc/client")

func GetRecipients(ctx context.Context) (*models.RetrieveRecipientResponseMsg, error) {

	// Construct a RetrieveRequestMsg according to the SFMC API requirements
	retrieveRequest := &models.RetrieveRequestMsg{
		RetrieveRequest: &models.RetrieveRequest{
			ObjectType: "Subscriber", // Assuming you want to retrieve subscriber data
			Properties: []string{
				"ID", "EmailAddress", "SubscriberKey", "Status", "EmailTypePreference",
			},
		},
	}

	// Set up your request
	// Call the Retrieve method and handle the response
	sfmcClient := client.NewSfmcAuthClient()

	var err error
	var response *models.RetrieveRecipientResponseMsg
	// Call the Retrieve method and handle the response
	response, err = sfmcClient.RetrieveRecipients(retrieveRequest)
	fmt.Println(response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func GetRecipient(ctx context.Context, subscriberKey string) (*models.RetrieveRecipientResponseMsg, error) {
	// Assume `NewSfmcAuthClient` creates a client that is authorized to make requests.
	sfmcClient := client.NewSfmcAuthClient()

	// Specify the request for retrieving a subscriber
	retrieveRequest := &models.RetrieveRequestMsg{
		RetrieveRequest: &models.RetrieveRequest{
			ObjectType: "Subscriber",
			Properties: []string{
				"ID", "EmailAddress", "SubscriberKey", "Status", "EmailTypePreference",
			},
			Filter: &models.SimpleFilterPart{
				XSIType:        "par:SimpleFilterPart",
				Property:       "SubscriberKey",
				SimpleOperator: "equals",
				Value:          []string{subscriberKey},
			},
		},
	}
	fmt.Println(retrieveRequest.RetrieveRequest)

	var err error
	var response *models.RetrieveRecipientResponseMsg
	// Call the Retrieve method and handle the response
	response, err = sfmcClient.RetrieveRecipient(retrieveRequest)
	if err != nil {
		return response, err
	}

	return response, nil
}
