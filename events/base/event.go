package base

type EventInterface interface {
	GetType() Type
}

type Event struct {
	Type Type `json:"type"`
}

func (e Event) GetType() Type {
	return e.Type
}
