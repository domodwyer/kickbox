package kickbox

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Client holds the HTTP client for interacting with the API, and the associated
// API key
type Client struct {
	apiKey string
	http   http.Client
}

// NewClient returns an instance of a Kickbox API Client
func NewClient(apiKey string) *Client {
	client := &Client{
		apiKey: apiKey,
		http:   http.Client{},
	}

	// Set the default timeout to 3 seconds
	client.SetTimeout(time.Second * 3)
	return client
}

// SetTimeout configures the request timeout value (includes connecting, waiting for a response, and reading the response)
func (c *Client) SetTimeout(time time.Duration) {
	c.http.Timeout = time
}

// Verify the given email address using Kickbox.io
func (c Client) Verify(address string) (*Result, error) {
	return c.verify(KickboxResultBuilder{}, c.url(address))
}

// Build the API endpoint given the email address and API key
func (c Client) url(address string) string {
	return fmt.Sprintf("https://api.kickbox.io/v2/verify?email=%s&apikey=%s", url.QueryEscape(address), url.QueryEscape(c.apiKey))
}

// Request and read the response of the HTTP request, returning a new Result struct
func (c Client) verify(rb ResultBuilder, url string) (*Result, error) {
	// Send our API request
	response, err := c.http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Did we get a HTTP 200?
	if response.StatusCode != 200 {
		msg := fmt.Sprintf("Kickbox API returned HTTP %d", response.Status)
		return nil, errors.New(msg)
	}

	// Read the response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Build our Result struct
	result, err := rb.NewResult(body)
	if err != nil {
		return nil, err
	}

	return result, nil
}
