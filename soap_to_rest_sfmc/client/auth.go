package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	util "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/soap_to_rest_sfmc/util"
	//gosoap "github.com/hooklift/gowsdl/soap"
	gosoap "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/soap_to_rest_sfmc/util"
)

// Assuming the previously shared structs and types...

// Configurations for SFMC API
const (
	InstanceAuthURL      = "https://mc166917qx7lk8ldd5cwvlp6ftzq.auth.marketingcloudapis.com/v2/token"
	InstanceClientID     = "bum2l8ir8ijnjqke1q1h0fno"
	InstanceClientSecret = "xFiOLKz1M2cDkCac1zFNwFll"
	InstanceAccountID    = "510009405"
)

type AuthRequest struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	AccountID    string `json:"account_id"`
}

// AuthResponse holds the structure for authentication response
type AuthResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	TokenType        string `json:"token_type"`
	Scope            string `json:"scope"`
	SoapInstanceUrl  string `json:"soap_instance_url"`
	RestInstanceUrl  string `json:"rest_instance_url"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorURI         string `json:"error_uri"`
}

// tokenCache to store the token and its expiration time
var tokenCache *AuthResponse

// GetToken retrieves a new token or uses the cached one
func GetToken() (*AuthResponse, error) {
	if tokenCache != nil && tokenCache.ExpiresIn > int(time.Now().Unix()) {
		return tokenCache, nil
	}

	authReq := AuthRequest{
		GrantType:    "client_credentials",
		ClientID:     InstanceClientID,
		ClientSecret: InstanceClientSecret,
		AccountID:    InstanceAccountID,
	}

	reqBodyBytes, err := json.Marshal(authReq)
	if err != nil {
		log.Fatal(err) // Handle the error according to your error handling strategy
	}

	req, err := http.NewRequest("POST", InstanceAuthURL, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		log.Fatal(err) // Handle the error appropriately
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	log.Println(resp)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Failed to read response body", err)
		}
		log.Printf("Server error response: %s\n", string(bodyBytes))
		return nil, err // or handle the error appropriately
	}

	var authResponse AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		return nil, err
	}

	authResponse.ExpiresIn += int(time.Now().Unix()) // update token expiration time based on current time
	tokenCache = &authResponse
	return tokenCache, nil
}

// AddAuthTokenHeader adds the authentication token to the SOAP header
func AddAuthTokenHeader(client *gosoap.Client) error {
	token, err := GetToken()

	if err != nil {
		return err
	}

	authHeader := &util.FuelOAuth{AccessToken: token.AccessToken}
	client.AddHeader(authHeader)
	return nil
}
