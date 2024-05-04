package main

import (
	"encoding/xml"
	"fmt"
	//"go/token"
	"log"

	sfmc "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/sfmc" // Update the import path to your package
	gosoap "github.com/hooklift/gowsdl/soap"                 // Assuming this is the SOAP client library
)


type FuelOAuth struct {
    XMLName     xml.Name `xml:"http://exacttarget.com fueloauth"`
    AccessToken string   `xml:",chardata"`
}

func main() {
    fuelOAuthHeader := FuelOAuth{
        AccessToken: "eyJhbGciOiJIUzI1NiIsImtpZCI6IjQiLCJ2ZXIiOiIxIiwidHlwIjoiSldUIn0.eyJhY2Nlc3NfdG9rZW4iOiJYNHowSHJzUDJEMXVmTk52WUFtMFYyWnYiLCJjbGllbnRfaWQiOiJidW0ybDhpcjhpam5qcWtlMXExaDBmbm8iLCJlaWQiOjUxMDAwOTI2Miwic3RhY2tfa2V5IjoiUzUwIiwicGxhdGZvcm1fdmVyc2lvbiI6MiwiY2xpZW50X3R5cGUiOiJTZXJ2ZXJUb1NlcnZlciIsInBpZCI6OTN9.4svWd6eFTI83bQOGj0XnSkaaMX6wKOJcpvt-dkwBo1w.fFdsTNKmltkOWnmx0ymJawJxyqq7oIsER2IwaI9JkbATnsI6SrZU2XgeVbTOW2EPRDGQZDjV89vLaAVvdi1JxMnInTVmTiFMHfAynUvXNesAICNeuS14Pqvnb-Gw5GNK1JDrBb0rqPUPQ0dHolWQ9DIZb1UK6wucUF-Kf98",
    }

    // Assuming you have a way to create and configure a SOAP client
    // Setup details like endpoint, credentials, etc., should be configured here
    url := "https://mc166917qx7lk8ldd5cwvlp6ftzq.soap.marketingcloudapis.com/Service.asmx" // Replace with the actual endpoint
    //token := "eyJhbGciOiJIUzI1NiIsImtpZCI6IjQiLCJ2ZXIiOiIxIiwidHlwIjoiSldUIn0.eyJhY2Nlc3NfdG9rZW4iOiJYaUNrcjY2VnkwdnN5YVhxNHB0UkZOSXciLCJjbGllbnRfaWQiOiJidW0ybDhpcjhpam5qcWtlMXExaDBmbm8iLCJlaWQiOjUxMDAwOTI2Miwic3RhY2tfa2V5IjoiUzUwIiwicGxhdGZvcm1fdmVyc2lvbiI6MiwiY2xpZW50X3R5cGUiOiJTZXJ2ZXJUb1NlcnZlciIsInBpZCI6OTN9.b1AMFcePnoXzCtpgtE3fmBqA6ZQaOgQb0bNyJNu5VvI.o1fYgMOzdhpvm4HxOaHqydyandqQ1ProEDr9qVEjdxOwObYVP2kYzIBR1yTzfqELI6-7oGT9fc9d72-ItszISyNtQf_r1YpygdFZP_zV9ZKZ8YaiaNeusrzmuFD6jxWziYsBhTyhIE0kSB6Toz3Sh_nd78VGJeNMv45CBB0"
    client := gosoap.NewClient(url) // This method would be your configured SOAP client constructor


    client.AddHeader(fuelOAuthHeader)

    // Create an instance of your SFMC SOAP client
    soapClient := sfmc.NewSoap(client)

    // Construct a RetrieveRequestMsg according to the SFMC API requirements
    retrieveRequest := &sfmc.RetrieveRequestMsg{
        RetrieveRequest: &sfmc.RetrieveRequest{
            ObjectType: "Subscriber", // Assuming you want to retrieve subscriber data
            Properties: []string{"ID", "EmailAddress"}, // Specify the properties you need
        },
    } // Set up your request
    // Call the Retrieve method and handle the response
    response, err := soapClient.Retrieve(retrieveRequest)
    if err != nil {
        log.Fatalf("Error retrieving data: %v", err)

    }

    // Assuming the response includes a slice of results that you want to print
    fmt.Println("Retrieve Successful, Response Data:")
    for _, result := range response.Results {
        fmt.Printf("ID: %d \n", result.ID)
    }
}
