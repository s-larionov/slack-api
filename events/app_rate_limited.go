package events

// Indicates your app's event subscriptions are being rate limited
// https://api.slack.com/events-api#rate_limiting
type AppRateLimited struct {
	Type              RequestType `json:"type"`
	Token             string      `json:"token"`
	TeamID            string      `json:"team_id"`
	MinuteRateLimited int64       `json:"minute_rate_limited"`
	APIAppID          string      `json:"api_app_id"`
}
