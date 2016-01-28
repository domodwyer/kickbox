package kickbox

import "testing"

func TestNewResult(t *testing.T) {
	rb := KickboxResultBuilder{}

	in := []byte("{\"result\":\"deliverable\",\"reason\":\"accepted_email\",\"role\":true,\"free\":true,\"disposable\":true,\"accept_all\":true,\"did_you_mean\":\"test\",\"sendex\":0.6,\"email\":\"dom@itsallbroken.com\",\"user\":\"dom\",\"domain\":\"itsallbroken.com\",\"success\":true,\"message\":\"message\"}")
	result, err := rb.NewResult(in)
	if err != nil {
		t.Error(err)
	}

	if !result.IsDeliverable() {
		t.Error("IsDeliverable() fail")
	}

	if result.IsUndeliverable() {
		t.Error("IsUndeliverable() fail")
	}

	if result.Reason != "accepted_email" {
		t.Error("Reason fail")
	}

	if !result.Success {
		t.Error("Success fail")
	}

	if !result.Role {
		t.Error("Role fail")
	}

	if !result.Free {
		t.Error("Free fail")
	}

	if !result.Disposable {
		t.Error("Disposable fail")
	}

	if !result.AcceptAll {
		t.Error("AcceptAll fail")
	}

	if result.Suggested != "test" {
		t.Error("Suggested fail")
	}

	if result.Sendex != 0.6 {
		t.Error("Sendex fail")
	}

	if result.Email != "dom@itsallbroken.com" {
		t.Error("Email fail")
	}

	if result.User != "dom" {
		t.Error("User fail")
	}

	if result.Domain != "itsallbroken.com" {
		t.Error("Domain fail")
	}

	if result.Message != "message" {
		t.Error("Message fail")
	}

}

func TestNewResult_inverted(t *testing.T) {
	rb := KickboxResultBuilder{}

	in := []byte("{\"result\":\"undeliverable\",\"reason\":\"nope\",\"role\":false,\"free\":false,\"disposable\":false,\"accept_all\":false,\"did_you_mean\":\"test\",\"sendex\":0.6,\"email\":\"dom@itsallbroken.com\",\"user\":\"dom\",\"domain\":\"itsallbroken.com\",\"success\":false,\"message\":\"message\"}")
	result, err := rb.NewResult(in)
	if err != nil {
		t.Error(err)
	}

	if result.IsDeliverable() {
		t.Error("IsDeliverable() fail")
	}

	if !result.IsUndeliverable() {
		t.Error("IsUndeliverable() fail")
	}

	if result.Reason != "nope" {
		t.Error("Reason fail")
	}

	if result.Success {
		t.Error("Success fail")
	}

	if result.Role {
		t.Error("Role fail")
	}

	if result.Free {
		t.Error("Free fail")
	}

	if result.Disposable {
		t.Error("Disposable fail")
	}

	if result.AcceptAll {
		t.Error("AcceptAll fail")
	}

	if result.Suggested != "test" {
		t.Error("Suggested fail")
	}

	if result.Sendex != 0.6 {
		t.Error("Sendex fail")
	}

	if result.Email != "dom@itsallbroken.com" {
		t.Error("Email fail")
	}

	if result.User != "dom" {
		t.Error("User fail")
	}

	if result.Domain != "itsallbroken.com" {
		t.Error("Domain fail")
	}

	if result.Message != "message" {
		t.Error("Message fail")
	}

}

func TestNewResult_risky(t *testing.T) {
	rb := KickboxResultBuilder{}

	in := []byte("{\"result\":\"risky\",\"reason\":\"nope\",\"role\":false,\"free\":false,\"disposable\":false,\"accept_all\":false,\"did_you_mean\":\"test\",\"sendex\":0.6,\"email\":\"dom@itsallbroken.com\",\"user\":\"dom\",\"domain\":\"itsallbroken.com\",\"success\":false,\"message\":\"message\"}")
	result, err := rb.NewResult(in)
	if err != nil {
		t.Error(err)
	}

	if !result.IsRisky() {
		t.Error("IsRisky() fail")
	}
}

func TestNewResult_unknown(t *testing.T) {
	rb := KickboxResultBuilder{}

	in := []byte("{\"result\":\"unknown\",\"reason\":\"nope\",\"role\":false,\"free\":false,\"disposable\":false,\"accept_all\":false,\"did_you_mean\":\"test\",\"sendex\":0.6,\"email\":\"dom@itsallbroken.com\",\"user\":\"dom\",\"domain\":\"itsallbroken.com\",\"success\":false,\"message\":\"message\"}")
	result, err := rb.NewResult(in)
	if err != nil {
		t.Error(err)
	}

	if !result.IsUnknown() {
		t.Error("IsRisky() fail")
	}
}
