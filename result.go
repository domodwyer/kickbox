package kickbox

// Define our interface
type ResultBuilder interface {
	NewResult(response []byte) (*Result, error)
}

// This is our actual Result struct we store the API response in
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
