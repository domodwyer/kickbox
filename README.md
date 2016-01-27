[![Build Status](https://travis-ci.org/domodwyer/kickbox.svg?branch=master)](https://travis-ci.org/domodwyer/kickbox)
# kickbox
An easy-to-use Go library using Kickbox.io for email address validation

##Usage
Install a tagged version with: `go get gopkg.in/domodwyer/kickbox.v1`

```Go
// Create a new client with your API key
kickbox := kickbox.NewClient("KICKBOX_TEST")

// Validate an email address
result, err := kickbox.Verify("dom@itsallbroken.com")
if err != nil {
	panic("Something went wrong :( ")
}

// Interact with the response!
if(result.IsDeliverable()) {
	log.Print("Send some stuff to ", result.Email)
}
```
