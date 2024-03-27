package models

type MenuData struct {
	Menus          []Menu          `json:"menus"`
	Categories     []Category      `json:"categories"`
	Items          []Item          `json:"items"`
	ModifierGroups []ModifierGroup `json:"modifierGroups"`
}

type Menu struct {
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	Subtitle            string                `json:"subtitle"`
	CategoryIDs         []string              `json:"category_ids"`
	ServiceAvailability []ServiceAvailability `json:"service_availability"`
}

type ServiceAvailability struct {
	DayOfWeek   string       `json:"day_of_week"`
	TimePeriods []TimePeriod `json:"time_periods"`
}

type TimePeriod struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type Category struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Subtitle string   `json:"subtitle"`
	Entities []Entity `json:"entities"`
}

type Entity struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type Item struct {
	ID               string           `json:"id"`
	ExternalData     string           `json:"external_data"`
	Title            string           `json:"title"`
	Description      string           `json:"description"`
	ImageURL         string           `json:"image_url"`
	PriceInfo        PriceInfo        `json:"price_info"`
	QuantityInfo     interface{}      `json:"quantity_info"` // Usado interface{} debido a la falta de campos específicos
	ModifierGroupIDs ModifierGroupIDs `json:"modifier_group_ids"`
}

type PriceInfo struct {
	Price            int          `json:"price"`
	CorePrice        int          `json:"core_price"`
	ContainerDeposit int          `json:"container_deposit"`
	Overrides        interface{}  `json:"overrides"` // Usado interface{} debido a la falta de detalles
	PricedByUnit     PricedByUnit `json:"priced_by_unit"`
}

type PricedByUnit struct {
	MeasurementType string      `json:"measurement_type"`
	LengthUnit      interface{} `json:"length_unit"`
	WeightUnit      interface{} `json:"weight_unit"`
	VolumeUnit      interface{} `json:"volume_unit"`
}

type ModifierGroupIDs struct {
	IDs []string `json:"ids"`
}

type ModifierGroup struct {
	QuantityInfo    QuantityInfo     `json:"quantity_info"`
	Title           string           `json:"title"`
	ExternalData    string           `json:"external_data"`
	ID              string           `json:"id"`
	ModifierOptions []ModifierOption `json:"modifier_options"`
}

type QuantityInfo struct {
	Quantity Quantity `json:"quantity"`
}

type Quantity struct {
	MinPermitted int `json:"min_permitted"`
	MaxPermitted int `json:"max_permitted"`
}

type Title struct {
	Translations map[string]string `json:"translations"`
}

type ModifierOption struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

func (menu *Menu) ContainsCategory(categoryId string) bool {
	slice := menu.CategoryIDs
	for _, v := range slice {
		if v == categoryId {
			return true
		}
	}
	return false
}
