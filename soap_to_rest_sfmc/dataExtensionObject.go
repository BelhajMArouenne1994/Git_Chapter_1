package soap_to_rest_sfmc

import (
	"context"
	"encoding/xml"
	//"fmt"
	//"log"
	//"net/http"
	//"github.com/gin-gonic/gin"
)

type RetrieveDataExtensionObjectResponseMsg struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI RetrieveResponseMsg"`

	OverallStatus string `xml:"OverallStatus,omitempty"`

	RequestID string `xml:"RequestID,omitempty"`

	Results []*RetrieveDataExtensionObjectResponseMsgAPIOBJECT `xml:"Results,omitempty"`
}

type RetrieveDataExtensionObjectResponseMsgAPIOBJECT struct {
	*ObjectExtension

	Name string `json:"name,omitempty"`

	Keys struct {
		Key []*APIProperty `nml:"key,omitempty"`
	} `nml:"keys,omitempty"`
}

func RetrieveDataExtensionObject(ctx context.Context) (*RetrieveDataExtensionObjectResponseMsg, error) {

	// Construct a RetrieveRequestMsg according to the SFMC API requirements
	retrieveRequest := &RetrieveRequestMsg{
		RetrieveRequest: &RetrieveRequest{
			ObjectType: "DataExtensionObject[67baea65-a4fc-ee11-b875-f40343c95ac8]", // Assuming you want to retrieve subscriber data
			Properties: []string{
				// essayer de trouver une facon pour mettre aussi automatiquement toutes les propriétés de APIObject struct
				"CustomerKey",
			},
		},
	}

	// Set up your request
	// Call the Retrieve method and handle the response
	sfmcClient := NewSfmcAuthClient()

	var err error
	var response *RetrieveDataExtensionObjectResponseMsg
	// Call the Retrieve method and handle the response
	response, err = sfmcClient.RetrieveDataExtensionObject(retrieveRequest)

	if err != nil {
		return response, err
	}

	return response, nil
}
