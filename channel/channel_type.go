package channel

const (
	TypePublic             Type = "C"       // Public channel. Used by channel_* events
	TypePrivate                 = "G"       // Private channel. Used by channel_* events
	TypePublicMessage           = "channel" // Public channel. Used by message_* events
	TypePrivateMessage          = "group"   // Private channel. Used by message_* events
	TypeDirectMessage           = "im"      // Direct message channel. Used by message_* events
	TypeMultiDirectMessage      = "mpim"    // Multiparty direct message channel. Used by message_* events
)

type Type string

func (t Type) IsPublic() bool {
	return t == TypePublic || t == TypePublicMessage
}

func (t Type) IsPrivate() bool {
	return t == TypePrivate || t == TypePrivateMessage
}

func (t Type) IsDirect() bool {
	return t == TypeDirectMessage
}

func (t Type) IsMultiDirect() bool {
	return t == TypeMultiDirectMessage
}
