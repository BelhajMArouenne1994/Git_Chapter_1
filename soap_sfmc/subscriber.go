package soap_sfmc

import (
	"context"
	"encoding/xml"
	"fmt"
	"time"
	//"fmt"
	//"log"
	//"net/http"
	//"github.com/gin-gonic/gin"
)

type RetrieveRecipientResponseMsg struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI RetrieveResponseMsg"`

	OverallStatus string `xml:"OverallStatus,omitempty"`

	RequestID string `xml:"RequestID,omitempty"`

	Results []*RetrieveRecipientResponseMsgAPIOBJECT `xml:"Results,omitempty"`
}

type RetrieveRecipientResponseMsgAPIOBJECT struct {
	*APIObject

	EmailAddress string `xml:"EmailAddress,omitempty" json:"emailAddress,omitempty"`

	Attributes []*Attribute `xml:"Attributes,omitempty" json:"attributes,omitempty"`

	SubscriberKey string `xml:"SubscriberKey,omitempty" json:"subscriberKey,omitempty"`

	UnsubscribedDate time.Time `xml:"UnsubscribedDate,omitempty" json:"unsubscribedDate,omitempty"`

	Status *SubscriberStatus `xml:"Status,omitempty" json:"status,omitempty"`

	PartnerType string `xml:"PartnerType,omitempty" json:"partnerType,omitempty"`

	EmailTypePreference *EmailType `xml:"EmailTypePreference,omitempty" json:"emailTypePreference,omitempty"`

	Lists []*SubscriberList `xml:"Lists,omitempty" json:"lists,omitempty"`

	GlobalUnsubscribeCategory *GlobalUnsubscribeCategory `xml:"GlobalUnsubscribeCategory,omitempty" json:"globalUnsubscribeCategory,omitempty"`

	SubscriberTypeDefinition *SubscriberTypeDefinition `xml:"SubscriberTypeDefinition,omitempty" json:"subscriberTypeDefinition,omitempty"`

	Addresses struct {
		Address []*SubscriberAddress `xml:"Address,omitempty" json:"address,omitempty"`
	} `xml:"Addresses,omitempty" json:"addresses,omitempty"`

	PrimarySMSAddress *SMSAddress `xml:"PrimarySMSAddress,omitempty" json:"primarySMSAddress,omitempty"`

	PrimarySMSPublicationStatus *SubscriberAddressStatus `xml:"PrimarySMSPublicationStatus,omitempty" json:"primarySMSPublicationStatus,omitempty"`

	PrimaryEmailAddress *EmailAddress `xml:"PrimaryEmailAddress,omitempty" json:"primaryEmailAddress,omitempty"`

	Locale *Locale `xml:"Locale,omitempty" json:"locale,omitempty"`
}

func GetRecipients(ctx context.Context) (*RetrieveRecipientResponseMsg, error) {

	// Construct a RetrieveRequestMsg according to the SFMC API requirements
	retrieveRequest := &RetrieveRequestMsg{
		RetrieveRequest: &RetrieveRequest{
			ObjectType: "Subscriber", // Assuming you want to retrieve subscriber data
			Properties: []string{
				"ID", "EmailAddress", "SubscriberKey", "Status", "EmailTypePreference",
			},
		},
	}

	// Set up your request
	// Call the Retrieve method and handle the response
	sfmcClient := NewSfmcAuthClient()

	var err error
	var response *RetrieveRecipientResponseMsg
	// Call the Retrieve method and handle the response
	response, err = sfmcClient.RetrieveRecipients(retrieveRequest)
	fmt.Println(response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func GetRecipient(ctx context.Context, subscriberKey string) (*RetrieveRecipientResponseMsg, error) {
	// Assume `NewSfmcAuthClient` creates a client that is authorized to make requests.
	sfmcClient := NewSfmcAuthClient()

	// Specify the request for retrieving a subscriber
	retrieveRequest := &RetrieveRequestMsg{
		RetrieveRequest: &RetrieveRequest{
			ObjectType: "Subscriber",
			Properties: []string{
				"ID", "EmailAddress", "SubscriberKey", "Status", "EmailTypePreference",
			},
			Filter: &SimpleFilterPart{
				XSIType:        "par:SimpleFilterPart",
				Property:       "SubscriberKey",
				SimpleOperator: "equals",
				Value:          []string{subscriberKey},
			},
		},
	}
	fmt.Println(retrieveRequest.RetrieveRequest)

	var err error
	var response *RetrieveRecipientResponseMsg
	// Call the Retrieve method and handle the response
	response, err = sfmcClient.RetrieveRecipient(retrieveRequest)
	if err != nil {
		return response, err
	}

	return response, nil
}
