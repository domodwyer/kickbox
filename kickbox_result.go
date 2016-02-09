package kickbox

import (
	"encoding/json"
)

// KickboxResultBuilder implements our ResultBuilder interface and creates the
// actual Result struct (the response from Kickbox)
type KickboxResultBuilder struct{}

// NewResult creates a new Result object from an JSON API response
func (b KickboxResultBuilder) NewResult(response []byte) (*Result, error) {
	result := Result{}
	if err := json.Unmarshal(response, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// IsDeliverable returns true if the API returns "result: deliverable"
func (r Result) IsDeliverable() bool {
	return (r.Result == "deliverable")
}

// IsUndeliverable returns true if the API returns "result: undeliverable"
func (r Result) IsUndeliverable() bool {
	return (r.Result == "undeliverable")
}

// IsRisky returns true if the API returns "result: risky"
func (r Result) IsRisky() bool {
	return (r.Result == "risky")
}

// IsUnknown returns true if the API returns "result: unknown"
func (r Result) IsUnknown() bool {
	return (r.Result == "unknown")
}
