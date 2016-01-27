# kickbox
An easy-to-use library using Kickbox.io for email address validation

###Example
```Go
// Create a new client with your Kickbox.io API key
kickbox := kickbox.NewClient("KICKBOX_TEST")

// Verify an address
result, err := kickbox.Verify("dom@itsallbroken.com")
if err != nil {
	panic("Something went wrong :( ")
}

if result.IsDeliverable() {
	// Send away!
	log.Print(result.Email)
}
```
