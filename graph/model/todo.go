package model

type Todo struct {
	ID     ID     `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID ID     `json:"-"`
}

func (t *Todo) User() {
	panic("not implemented")
}
