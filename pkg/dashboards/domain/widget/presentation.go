package widget

type Presentation struct {
	Title string `json:"title"`
}

func (p Presentation) getHashKey() string {
	return p.Title
}

func (p Presentation) Equals(other Presentation) bool {
	return p.Title == other.Title
}
