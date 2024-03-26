package entities

type Menu struct {
	Id           string       `json:"id" bson:"id"`
	Title        string       `json:"title" bson:"title"`
	Subtitle     string       `json:"subtitle" bson:"subtitle"`
	Availability Availability `json:"availability" bson:"availability"`
	Categories   []Category   `json:"categories" bson:"categories"`
}

func NewMenu() Menu {
	return Menu{}
}
