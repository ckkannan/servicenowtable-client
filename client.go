package servicenowtable

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HostURL - Default Hashicups URL
const CSN_URL string = "https://dev161016.service-now.com"

type ServicenowtableProviderInput struct {
	Sn_url    string
	Sn_user   string
	Sn_pass   string
	SSLIgnore bool
	Version   string
}

// Client -
type Client struct {
	sn_url     string
	HTTPClient *http.Client
	Token      string
	Auth       AuthStruct
}

// AuthStruct -
type AuthStruct struct {
	Sn_url  string `json:"sn_url"`
	Sn_user string `json:"sn_user"`
	Sn_pass string `json:"sn_pass"`
}

// AuthResponse -
type AuthResponse struct {
	Sn_user     string `json:"sn_user"`
	Sn_username string `json:"sn_username"`
	Token       string `json:"token"`
}

// NewClient -
func NewClient(servicenow ServicenowtableProviderInput) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Hashicups URL
		sn_url: CSN_URL,
	}

	if servicenow.Sn_url == "" {
		c.sn_url = servicenow.Sn_url
	}

	// If username or password not provided, return empty client
	if servicenow.Sn_user == "" || servicenow.Sn_pass == "" {
		return &c, nil
	}

	c.Auth = AuthStruct{
		Sn_url:  servicenow.Sn_url,
		Sn_user: servicenow.Sn_user,
		Sn_pass: servicenow.Sn_pass,
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	// token := c.Token

	// if authToken != nil {
	// 	token = *authToken
	// }

	// req.Header.Set("Authorization", token)
	req.SetBasicAuth(c.Auth.Sn_user, c.Auth.Sn_pass)
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
