package services

import (
	"context"
	
	models "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/soap_to_rest_sfmc/models"
	client "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/soap_to_rest_sfmc/client"

	//"encoding/xml"
	//"fmt"
	//"log"
	//"net/http"
	//"github.com/gin-gonic/gin"
)

/*
	type DefinitionRequestMsg struct {
		XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI DefinitionRequestMsg"`

		DescribeRequests *ArrayOfObjectDefinitionRequest `xml:"DescribeRequests,omitempty"`
	}

	type ArrayOfObjectDefinitionRequest struct {
		ObjectDefinitionRequest []*ObjectDefinitionRequest `xml:"ObjectDefinitionRequest,omitempty"`
	}

	type ObjectDefinitionRequest struct {
		Client *ClientID `xml:"Client,omitempty"`

		ObjectType string `xml:"ObjectType,omitempty"`
	}
*/
func Describe(ctx context.Context) (*models.DefinitionResponseMsg, error) {

	describeRequest := &models.DefinitionRequestMsg{
		DescribeRequests: &models.ArrayOfObjectDefinitionRequest{
			ObjectDefinitionRequest: []*models.ObjectDefinitionRequest{
				{
					//Client:     &ClientID{ClientID: 12345}, // Example client ID
					ObjectType: "DataExtensionField",
				},
			},
		},
	}

	// Set up your request
	// Call the Retrieve method and handle the response
	sfmcClient := client.NewSfmcAuthClient()

	var err error
	var response *models.DefinitionResponseMsg
	// Call the Retrieve method and handle the response
	response, err = sfmcClient.Describe(describeRequest)

	if err != nil {
		return response, err
	}

	return response, nil
}
