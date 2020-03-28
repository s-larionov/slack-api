package events

const (
	RequestTypeVerification   RequestType = "url_verification"
	RequestTypeWrapper                    = "event_callback"
	RequestTypeAppRateLimited             = "app_rate_limited"
)

type RequestType string
