package events

// Verifies ownership of an Events API Request URL
// https://api.slack.com/events/url_verification
// Expected scopes: none
type Verification struct {
	Type      RequestType `json:"type"` // url_verification
	Token     string      `json:"token"`
	Challenge string      `json:"challenge"`
}
