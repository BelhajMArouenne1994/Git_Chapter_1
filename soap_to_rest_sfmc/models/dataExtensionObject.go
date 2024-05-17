package models

import (
	"encoding/xml"
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
