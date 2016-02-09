package kickbox

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type TestResultBuilder struct{}

// Creates a new Result object from an JSON API response
func (b TestResultBuilder) NewResult(response []byte) (*Result, error) {
	return &Result{Message: string(response)}, nil
}

// Ensure the API key is set
func TestNewClient(t *testing.T) {
	expect := "KICKBOX_TEST"
	client := NewClient(expect)

	if client.apiKey != expect {
		t.Error("Client API key not as expected")
	}
}

// Ensure the default timeout is set
func TestNewClient_setDefaultTimeout(t *testing.T) {
	client := NewClient("KICKBOX_TEST")

	if client.http.Timeout != (time.Second * 3) {
		t.Error("Default timeout not set")
	}
}

// Ensure the timeout is settable
func TestNewClient_setTimeout(t *testing.T) {
	client := NewClient("KICKBOX_TEST")

	client.SetTimeout(time.Second * 8)
	if client.http.Timeout != (time.Second * 8) {
		t.Error("Timeout not set")
	}
}

// Test the correct API URL is generated - email
func TestUrl_email(t *testing.T) {
	client := NewClient("KICKBOX_TEST")

	tests := []struct {
		address string
		url     string
	}{
		{"dom@itsallbroken.com", "https://api.kickbox.io/v2/verify?email=dom%40itsallbroken.com&apikey=KICKBOX_TEST"},
		{"somEjuNK@*($@.2coen19e,1.2e12e.1", "https://api.kickbox.io/v2/verify?email=somEjuNK%40%2A%28%24%40.2coen19e%2C1.2e12e.1&apikey=KICKBOX_TEST"},
		{"a", "https://api.kickbox.io/v2/verify?email=a&apikey=KICKBOX_TEST"},
		{"123", "https://api.kickbox.io/v2/verify?email=123&apikey=KICKBOX_TEST"},
		{"", "https://api.kickbox.io/v2/verify?email=&apikey=KICKBOX_TEST"},
	}

	for _, test := range tests {
		actual := client.url(test.address)
		if test.url != actual {
			t.Error("URL generation failed expected: ", test.url, " got: ", actual)
		}
	}
}

// Ensure a Result object is returned with a successful request
func TestVerify_200(t *testing.T) {
	tb := TestResultBuilder{}
	client := NewClient("KICKBOX_TEST")

	// Start our test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "PASSED")
	}))
	defer ts.Close()

	// Make an API request to our test server
	result, err := client.verify(tb, ts.URL)
	if err != nil {
		t.Error("Failed to interact with the API server")
	}

	// Ensure the Result object was returned
	if result.Message != "PASSED" {
		t.Error("Unexpected result")
	}
}
