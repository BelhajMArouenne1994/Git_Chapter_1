package models

import (
	"encoding/xml"
)

type RetrieveDEFieldResponseMsg struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI RetrieveResponseMsg"`

	OverallStatus string `xml:"OverallStatus,omitempty"`

	RequestID string `xml:"RequestID,omitempty"`

	Results []*RetrieveDEFieldResponseMsgAPIOBJECT `xml:"Results,omitempty"`
}

type RetrieveDEFieldResponseMsgAPIOBJECT struct {
	*PropertyDefinition

	Ordinal int32 `json:"ordinal,omitempty"`

	IsPrimaryKey bool `json:"isPrimaryKey,omitempty"`

	FieldType *DataExtensionFieldType `json:"fieldType,omitempty"`

	DataExtension *DataExtension `json:"dataExtension,omitempty"`

	StorageType *DataExtensionFieldStorageType `json:"storageType,omitempty"`
}