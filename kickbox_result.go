package kickbox

import (
	"encoding/json"
)

// Define our builder type that implements the ResultBuilder interface
type KickboxResultBuilder struct{}

// Creates a new Result object from an JSON API response
func (b KickboxResultBuilder) NewResult(response []byte) (*Result, error) {
	result := Result{}
	if err := json.Unmarshal(response, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

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
