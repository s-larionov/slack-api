package star

const (
	ItemTypeMessage     = "message"
	ItemTypeFile        = "file"
	ItemTypeFileComment = "file_comment"
	ItemTypeChannel     = "channel"
)

type ItemType string

type Item struct {
	Type ItemType `json:"type"`
}

type ItemMessage struct {
	Item
	Channel string      `json:"channel"`
	Message interface{} `json:"message"` // TODO
}

type ItemChannel struct {
	Item
	Channel string `json:"channel"`
}

type ItemFile struct {
	Item
	Channel string      `json:"channel"`
	File    interface{} `json:"file"` // TODO
}

type ItemFileComment struct {
	Item
	File    interface{} `json:"file"`    // TODO
	Comment interface{} `json:"comment"` // TODO
}
