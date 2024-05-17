package models

import (
	"encoding/xml"
	"time"
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