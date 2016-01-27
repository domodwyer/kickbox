package kickbox

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
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

func TestUrl(t *testing.T) {
	client := NewClient("KICKBOX_TEST")

	tests := []struct {
		address string
		url     string
	}{
		{"dom@itsallbroken.com", "test"},
	}

	for _, test := range tests {
		actual := client.url(test.address)
		if test.url != actual {
			t.Error("URL generation failed", test.url, actual)
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
