package kickbox

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	apiKey string
	http   http.Client
}

func NewClient(apiKey string) Client {
	return Client{
		apiKey: apiKey,
		http:   http.Client{},
	}
}

// Configure the request timeout value (includes connecting, waiting for a response, and reading the response)
func (c Client) SetTimeout(time time.Duration) {
	c.http.Timeout = time
}

func (c Client) Verify(address string) (*Result, error) {
	return c.verify(KickboxResultBuilder{}, c.url(address))
}

func (c Client) url(address string) string {
	return fmt.Sprintf("https://api.kickbox.io/v2/verify?email=%s&apikey=%s", url.QueryEscape(address), url.QueryEscape(c.apiKey))
}

func (c Client) verify(rb ResultBuilder, url string) (*Result, error) {
	// Send our API request
	response, err := c.http.Get(url)
	if err != nil {
		return nil, err
	}

	// Did we get a HTTP 200?
	if response.StatusCode != 200 {
		msg := fmt.Sprintf("Kickbox API returned HTTP %d", response.Status)
		return nil, errors.New(msg)
	}

	// Read the response
	defer response.Body.Close()
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
