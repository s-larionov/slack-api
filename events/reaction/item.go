package reaction

const (
	ItemTypeMessage     = "message"
	ItemTypeFile        = "file"
	ItemTypeFileComment = "file_comment"
)

type ItemType string

type Item struct {
	Type    ItemType `json:"type"`
	Channel string   `json:"channel"`
	TS      string   `json:"ts"`
}
