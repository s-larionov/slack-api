package subteam

type Subteam struct {
	ID          string
	TeamID      string
	IsUserGroup bool `json:"is_usergroup"`
	Name        string
	Description string
	Handle      string
	IsExternal  bool
	DateCreate  int64
	DateUpdate  int64
	DateDelete  int64
	AutoType    interface{} // FIXME
	CreatedBy   string
	UpdatedBy   string
	DeletedBy   string
	Preferences Preferences `json:"prefs"`
	Users       []string
	UserCount   int64
}
