package client

import (
	//"fmt"
	//"log"
	"time"

	gosoap "github.com/hooklift/gowsdl/soap"
	models  "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/soap_to_rest_sfmc/models"
	//gosoap "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/util"
)

const (
	SoapURL = "https://mc166917qx7lk8ldd5cwvlp6ftzq.soap.marketingcloudapis.com/Service.asmx"
)

func NewSfmcAuthClient() models.Soap {

	// Initialize the SOAP client
	soapClient := gosoap.NewClient(SoapURL, gosoap.WithTimeout(1500*time.Millisecond))

	// Add the auth token header
	if err := AddAuthTokenHeader(soapClient); err != nil {
		return nil
	}

	// Create an instance of your SFMC SOAP client
	sfmcClient := models.NewSoap(soapClient)

	return sfmcClient
}
