package events

const (
	RequestTypeVerification   RequestType = "url_verification"
	RequestTypeWrapper        RequestType = "event_callback"
	RequestTypeAppRateLimited RequestType = "app_rate_limited"
)

type RequestType string
