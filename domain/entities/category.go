package entities

type Category struct {
	Id       string `json:"id" bson:"id"`
	Title    string `json:"title" bson:"title"`
	Subtitle string `json:"subtitle" bson:"subtitle"`
	Items    []Item `json:"items" bson:"items"`
}

func NewCategory(
	Id string,
	Title string,
	Subtitle string,
) Category {
	return Category{
		Id:       Id,
		Title:    Title,
		Subtitle: Subtitle,
		Items:    []Item{},
	}
}
