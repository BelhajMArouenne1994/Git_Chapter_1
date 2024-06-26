package util

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type FuelOAuth struct {
    XMLName     xml.Name `xml:"http://exacttarget.com fueloauth"`
    AccessToken string   `xml:",chardata"`
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  *SOAPHeader
	XmlnsXsi string   `xml:"xmlns:xsi,attr"`
	Body    SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Items []interface{} `xml:",omitempty"`
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

// UnmarshalXML unmarshals SOAPBody xml
func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, _ xml.StartElement) error {
    if b.Content == nil {
        return xml.UnmarshalError("Content must be a pointer to a struct")
    }

    var (
        token    xml.Token
        err      error
        consumed bool
    )

Loop:
    for {
        if token, err = d.Token(); err != nil {
            return err
        }

        if token == nil {
            break
        }

        switch se := token.(type) {
        case xml.StartElement:
            fmt.Printf("Encountered Element: %+v\n", se)  // Log every element encountered
            if consumed {
                return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
            } else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
                b.Fault = &SOAPFault{}
                b.Content = nil

                err = d.DecodeElement(b.Fault, &se)
                if err != nil {
                    return err
                }

                consumed = true
            } else {
                if err = d.DecodeElement(b.Content, &se); err != nil {
                    return err
                }

                consumed = true
            }
        case xml.EndElement:
            break Loop
        }
    }

    return nil
}

/*
func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, _ xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}
*/

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

func (f *SOAPFault) Error() string {
	return f.String
}

const (
	// Predefined WSS namespaces to be used in
	WssNsWSSE string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	WssNsWSU  string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
	WssNsType string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText"
)

type WSSSecurityHeader struct {
	XMLName   xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ wsse:Security"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	MustUnderstand string `xml:"mustUnderstand,attr,omitempty"`

	Token *WSSUsernameToken `xml:",omitempty"`
}

type WSSUsernameToken struct {
	XMLName   xml.Name `xml:"wsse:UsernameToken"`
	XmlNSWsu  string   `xml:"xmlns:wsu,attr"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	Id string `xml:"wsu:Id,attr,omitempty"`

	Username *WSSUsername `xml:",omitempty"`
	Password *WSSPassword `xml:",omitempty"`
}

type WSSUsername struct {
	XMLName   xml.Name `xml:"wsse:Username"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	Data string `xml:",chardata"`
}

type WSSPassword struct {
	XMLName   xml.Name `xml:"wsse:Password"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`
	XmlNSType string   `xml:"Type,attr"`

	Data string `xml:",chardata"`
}

// NewWSSSecurityHeader creates WSSSecurityHeader instance
func NewWSSSecurityHeader(user, pass, tokenID, mustUnderstand string) *WSSSecurityHeader {
	hdr := &WSSSecurityHeader{XmlNSWsse: WssNsWSSE, MustUnderstand: mustUnderstand}
	hdr.Token = &WSSUsernameToken{XmlNSWsu: WssNsWSU, XmlNSWsse: WssNsWSSE, Id: tokenID}
	hdr.Token.Username = &WSSUsername{XmlNSWsse: WssNsWSSE, Data: user}
	hdr.Token.Password = &WSSPassword{XmlNSWsse: WssNsWSSE, XmlNSType: WssNsType, Data: pass}
	return hdr
}

type basicAuth struct {
	Login    string
	Password string
}

type options struct {
	tlsCfg      *tls.Config
	auth        *basicAuth
	timeout     time.Duration
	httpHeaders map[string]string
}

var defaultOptions = options{
	timeout: time.Duration(30 * time.Second),
}

// A Option sets options such as credentials, tls, etc.
type Option func(*options)

// WithBasicAuth is an Option to set BasicAuth
func WithBasicAuth(login, password string) Option {
	return func(o *options) {
		o.auth = &basicAuth{Login: login, Password: password}
	}
}

// WithTLS is an Option to set tls config
func WithTLS(tls *tls.Config) Option {
	return func(o *options) {
		o.tlsCfg = tls
	}
}

// WithTimeout is an Option to set default HTTP dial timeout
func WithTimeout(t time.Duration) Option {
	return func(o *options) {
		o.timeout = t
	}
}

// WithHTTPHeaders is an Option to set global HTTP headers for all requests
func WithHTTPHeaders(headers map[string]string) Option {
	return func(o *options) {
		o.httpHeaders = headers
	}
}

// Client is soap client
type Client struct {
	url     string
	opts    *options
	headers []interface{}
}

// NewClient creates new SOAP client instance
func NewClient(url string, opt ...Option) *Client {
	opts := defaultOptions
	for _, o := range opt {
		o(&opts)
	}
	return &Client{
		url:  url,
		opts: &opts,
	}
}

// AddHeader adds envelope header
func (s *Client) AddHeader(header interface{}) {
	s.headers = append(s.headers, header)
}

// Call performs HTTP POST request
/*func (s *Client) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{
		Body: SOAPBody{
			Content: request,
		},
	}

	// Handling custom headers if any
	if s.headers != nil && len(s.headers) > 0 {
		envelope.Header = &SOAPHeader{Items: s.headers}
	}

	// Marshal the envelope using xml.MarshalIndent
	output, err := xml.MarshalIndent(envelope, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling XML:", err)
		return err
	}

	buffer := bytes.NewBuffer(output)
	fmt.Println(string(output)) // Print XML for debugging

	// Create HTTP request
	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}

	if s.opts.auth != nil {
		req.SetBasicAuth(s.opts.auth.Login, s.opts.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	req.Header.Add("SOAPAction", soapAction)
	req.Header.Set("User-Agent", "gowsdl/0.1")
	if s.opts.httpHeaders != nil {
		for k, v := range s.opts.httpHeaders {
			req.Header.Set(k, v)
		}
	}
	req.Close = true

	// Configure HTTP client and send request
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: s.opts.tlsCfg,
			Dial: func(network, addr string) (net.Conn, error) {
				return net.DialTimeout(network, addr, s.opts.timeout)
			},
		},
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Read and process response
	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		return nil
	}

	respEnvelope := new(SOAPEnvelope)
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	if respEnvelope.Body.Fault != nil {
		return fmt.Errorf("SOAP Fault: %s", respEnvelope.Body.Fault)
	}

	return nil
}
*/

/*
func (s *Client) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{}

	if s.headers != nil && len(s.headers) > 0 {
		soapHeader := &SOAPHeader{Items: make([]interface{}, len(s.headers))}
		copy(soapHeader.Items, s.headers)
		envelope.Header = soapHeader
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)

	if err := encoder.Encode(envelope); err != nil {
		fmt.Println(envelope)

		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}


	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.opts.auth != nil {
		req.SetBasicAuth(s.opts.auth.Login, s.opts.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	req.Header.Add("SOAPAction", soapAction)
	req.Header.Set("User-Agent", "gowsdl/0.1")
	if s.opts.httpHeaders != nil {
		for k, v := range s.opts.httpHeaders {
			req.Header.Set(k, v)
		}
	}
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: s.opts.tlsCfg,
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, s.opts.timeout)
		},
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		return nil
	}

	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	
	fmt.Println(respEnvelope.Body)
	
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}
*/
func (s *Client) Call(soapAction string, request, response interface{}) error {
    envelope := SOAPEnvelope{}

    // Add SOAP headers if any
    if s.headers != nil && len(s.headers) > 0 {
        soapHeader := &SOAPHeader{Items: make([]interface{}, len(s.headers))}
        copy(soapHeader.Items, s.headers)
        envelope.Header = soapHeader
    }

    envelope.Body.Content = request
    buffer := new(bytes.Buffer)

    // Create a new XML encoder
    encoder := xml.NewEncoder(buffer)
    encoder.Indent("", "  ") // Optional: Add indentation to XML for better readability

    // Encode the envelope struct to XML
    if err := encoder.Encode(envelope); err != nil {
        fmt.Println("Error encoding XML:", err)
        return err
    }

    if err := encoder.Flush(); err != nil {
        fmt.Println("Error flushing encoder:", err)
        return err
    }


	// Print the request body for debugging purposes
	fmt.Println(buffer.String())

    // Convert buffer to string and print it for debugging
    xmlString := buffer.String()
    fmt.Println("SOAP Request:", xmlString)

    // Create an HTTP request
    req, err := http.NewRequest("POST", s.url, buffer)
    if err != nil {
        fmt.Println("Error creating HTTP request:", err)
        return err
    }

    // Set headers
    if s.opts.auth != nil {
        req.SetBasicAuth(s.opts.auth.Login, s.opts.auth.Password)
    }
    req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
    req.Header.Add("SOAPAction", soapAction)
    req.Header.Set("User-Agent", "gowsdl/0.1")
    if s.opts.httpHeaders != nil {
        for k, v := range s.opts.httpHeaders {
            req.Header.Set(k, v)
        }
    }
    req.Close = true

    // Set up the transport with the provided TLS configuration and timeout
    tr := &http.Transport{
        TLSClientConfig: s.opts.tlsCfg,
        Dial: func(network, addr string) (net.Conn, error) {
            return net.DialTimeout(network, addr, s.opts.timeout)
        },
    }

    // Send the request
    client := &http.Client{Transport: tr}
    res, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return err
    }
    defer res.Body.Close()

    // Read and handle the response
    rawBody, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return err
    }

    // Check if the response body is empty
    if len(rawBody) == 0 {
        fmt.Println("Received empty response body")
        return nil
    }

    // Unmarshal the response into the provided response struct
    respEnvelope := new(SOAPEnvelope)
    respEnvelope.Body = SOAPBody{Content: response}
    if err := xml.Unmarshal(rawBody, respEnvelope); err != nil {
        fmt.Println("Error unmarshalling response:", err)
        return err
    }

    // Handle SOAP fault if any
    if fault := respEnvelope.Body.Fault; fault != nil {
        fmt.Println("SOAP Fault:", fault)
        return fault
    }

    return nil
}
