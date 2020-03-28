package message

const (
	SubTypeGeneric         SubType = ""
	SubTypeBotMessage              = "bot_message"       // A message was posted by an integration
	SubTypeEkmAccessDenied         = "ekm_access_denied" // Message content redacted due to Enterprise Key Management (EKM)
	SubTypeMeMessage               = "me_message"        // A /me message was sent
	SubTypeMessageChanged          = "message_changed"   // A message was changed
	SubTypeMessageDeleted          = "message_deleted"   // A message was deleted
	SubTypeMessageReplied          = "message_replied"   // A message thread received a reply
	SubTypeThreadBroadcast         = "thread_broadcast"  // A message thread's reply was broadcast to a channel
)

type SubType string
