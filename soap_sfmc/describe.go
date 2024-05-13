package soap_sfmc

import (
	"context"
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
func Describe(ctx context.Context) (*DefinitionResponseMsg, error) {

	describeRequest := &DefinitionRequestMsg{
		DescribeRequests: &ArrayOfObjectDefinitionRequest{
			ObjectDefinitionRequest: []*ObjectDefinitionRequest{
				{
					//Client:     &ClientID{ClientID: 12345}, // Example client ID
					ObjectType: "DataExtensionField",
				},
			},
		},
	}

	// Set up your request
	// Call the Retrieve method and handle the response
	sfmcClient := NewSfmcAuthClient()

	var err error
	var response *DefinitionResponseMsg
	// Call the Retrieve method and handle the response
	response, err = sfmcClient.Describe(describeRequest)

	if err != nil {
		return response, err
	}

	return response, nil
}
