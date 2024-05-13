package soap_sfmc

import (
	//"fmt"
	//"log"
	"time"

	//util "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/util"
	gosoap "github.com/hooklift/gowsdl/soap"
)

const (
	SoapURL = "https://mc166917qx7lk8ldd5cwvlp6ftzq.soap.marketingcloudapis.com/Service.asmx"
)

func NewSfmcAuthClient() Soap {

	// Initialize the SOAP client
	soapClient := gosoap.NewClient(SoapURL, gosoap.WithTimeout(1500*time.Millisecond))

	// Add the auth token header
	if err := AddAuthTokenHeader(soapClient); err != nil {
		return nil
	}

	// Create an instance of your SFMC SOAP client
	sfmcClient := NewSoap(soapClient)

	return sfmcClient
}