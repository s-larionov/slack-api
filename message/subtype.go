package message

const (
	SubTypeGeneric         SubType = ""
	SubTypeBotMessage      SubType = "bot_message"       // A message was posted by an integration
	SubTypeEkmAccessDenied SubType = "ekm_access_denied" // Message content redacted due to Enterprise Key Management (EKM)
	SubTypeMeMessage       SubType = "me_message"        // A /me message was sent
	SubTypeMessageChanged  SubType = "message_changed"   // A message was changed
	SubTypeMessageDeleted  SubType = "message_deleted"   // A message was deleted
	SubTypeMessageReplied  SubType = "message_replied"   // A message thread received a reply
	SubTypeThreadBroadcast SubType = "thread_broadcast"  // A message thread's reply was broadcast to a channel
)

type SubType string
