package events

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"slack/events/message"
)

const wrapperEvent = `
{
  "token": "token123",
  "team_id": "myteam",
  "api_app_id": "appid",
  "event": {
    "client_msg_id": "msg-id-123",
    "type": "message",
    "text": "test",
    "user": "userid",
    "ts": "1585317662.002300",
    "team": "teamid",
    "blocks": [
      {
        "type": "rich_text",
        "block_id": "0mfAr",
        "elements": [
          {
            "type": "rich_text_section",
            "elements": [
              {
                "type": "text",
                "text": "test"
              }
            ]
          }
        ]
      }
    ],
    "channel": "C010FAL165A",
    "event_ts": "1585317662.002300",
    "channel_type": "channel"
  },
  "type": "event_callback",
  "event_id": "Ev010Z95QUB1",
  "event_time": 1585317662,
  "authed_users": [
    "userid"
  ]
}
`

func TestWrapper_UnmarshalJSON(t *testing.T) {
	Convey("", t, func() {
		var w Wrapper
		err := json.Unmarshal([]byte(wrapperEvent), &w)

		So(err, ShouldBeNil)

		So(w.Type, ShouldEqual, RequestTypeWrapper)
		So(w.AuthedUsers, ShouldHaveLength, 1)
		So(w.AuthedUsers[0], ShouldEqual, "userid")
		So(w.EventTime, ShouldEqual, 1585317662)
		So(w.EventID, ShouldEqual, "Ev010Z95QUB1")
		So(w.APIAppID, ShouldEqual, "appid")
		So(w.TeamID, ShouldEqual, "myteam")
		So(w.Token, ShouldEqual, "token123")

		So(w.Event, ShouldHaveSameTypeAs, &message.Message{})
	})
}
