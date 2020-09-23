package cloudcraft

import (
	"log"
	"net/http"
	"net/url"
)

// Client is the object that handles talking to the cloudcraaft API. This maintains
// state information for a particular application connection.
type Client struct {
	apiKey     string
	baseURL    *url.URL
	HTTPClient *http.Client
}

// NewClient returns a new Cloudcraft.Client which can be used to access the API
// methods. The expected argument is the API key.
func NewClient(apiKey, baseurl string) *Client {
	baseURL, err := url.Parse("https://api.cloudcraft.co")
	if err != nil {
		log.Fatal(err)
	}

	return &Client{
		apiKey:     apiKey,
		baseURL:    baseURL,
		HTTPClient: http.DefaultClient,
	}
}

// SetKeys changes the value of apiKey.
func (c *Client) SetKeys(apiKey string) {
	c.apiKey = apiKey
}

// SetBaseURL changes the value of baseUrl.
func (c *Client) SetBaseURL(baseURL *url.URL) {
	c.baseURL = baseURL
}

// GetBaseURL returns the baseUrl.
func (c *Client) GetBaseURL() *url.URL {
	return c.baseURL
}
