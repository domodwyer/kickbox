[![Build Status](https://travis-ci.org/domodwyer/kickbox.svg?branch=master)](https://travis-ci.org/domodwyer/kickbox) [![GoDoc](https://godoc.org/github.com/domodwyer/kickbox?status.svg)](https://godoc.org/github.com/domodwyer/kickbox)
# kickbox
An easy-to-use Go library using Kickbox.io for email address validation

## Example
Install a tagged version with: `go get github.com/domodwyer/kickbox`

```Go
// Create a new client with your API key
kb := kickbox.NewClient("KICKBOX_TEST")

// Verify an email address
result, err := kb.Verify("dom@itsallbroken.com")
if err != nil {
	panic("something went wrong :( ")
}

// Interact with the response!
if(result.IsDeliverable()) {
	log.Print("Send some stuff to ", result.Email)
}
```

## Usage
The `Verify()` method returns an instance of `Kickbox.Result` and is mapped onto the Kickbox API response:

```Go
type Result struct {
	Success    bool    `json:"success"`
	Result     string  `json:"result"`
	Reason     string  `json:"reason"`
	Disposable bool    `json:"disposable"`
	Role       bool    `json:"role"`
	Free       bool    `json:"free"`
	AcceptAll  bool    `json:"accept_all"`
	Suggested  string  `json:"did_you_mean"`
	Sendex     float32 `json:"sendex"`
	Email      string  `json:"email"`
	User       string  `json:"user"`
	Domain     string  `json:"domain"`
	Message    string  `json:"message"`
}
```

For convenience `Kickbox.Result` also recieves several helper methods:
```Go
func (r Result) IsDeliverable() bool {
	return (r.Result == "deliverable")
}

func (r Result) IsUndeliverable() bool {
	return (r.Result == "undeliverable")
}

func (r Result) IsRisky() bool {
	return (r.Result == "risky")
}

func (r Result) IsUnknown() bool {
	return (r.Result == "unknown")
}
```
