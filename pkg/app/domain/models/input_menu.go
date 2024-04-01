package models

type (
	InputMenu struct {
		Menus          []Menu          `json:"menus"`
		Categories     []Category      `json:"categories"`
		Items          []Item          `json:"items"`
		ModifierGroups []ModifierGroup `json:"modifier_groups"`
	}

	Category struct {
		ID       string   `json:"id"`
		Title    string   `json:"title"`
		Subtitle string   `json:"subtitle"`
		Entities []Entity `json:"entities"`
	}

	Entity struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	}

	Item struct {
		ID                string            `json:"id"`
		ExternalData      string            `json:"external_data"`
		Title             string            `json:"title"`
		Description       string            `json:"description"`
		ImageURL          string            `json:"image_url"`
		PriceInfo         PriceInfo         `json:"price_info"`
		QuantityInfo      QuantityInfo      `json:"quantity_info"`
		ModifierGrouopIds ModifierGrouopIds `json:"modifier_group_ids"`
	}

	PriceInfo struct {
		Price        float32      `json:"price"`
		CorePrice    float32      `json:"core_price"`
		PricedByUnit PricedByUnit `json:"priced_by_unit"`
	}

	PricedByUnit struct {
		MeasurementType string `json:"measurement_type"`
		LengthUnit      int64  `json:"length_unit"`
		WeightUnit      int64  `json:"weight_unit"`
		VolumeUnit      int64  `json:"volume_unit"`
	}

	QuantityInfo struct {
		Quantity Quantity `json:"quantity"`
	}

	Quantity struct {
		MinPermitted           int `json:"min_permitted"`
		MaxPermitted           int `json:"max_permitted"`
		IsMinPermittedOptional int `json:"is_min_permitted_optional"`
	}

	Menu struct {
		ID                  string                `json:"id"`
		Title               string                `json:"title"`
		Subtitle            string                `json:"subtitle"`
		CategoryIds         []string              `json:"category_ids"`
		ServiceAvailability []ServiceAvailability `json:"service_availability"`
	}

	ServiceAvailability struct {
		DayOfWeek   string       `json:"day_of_week"`
		TimePeriods []TimePeriod `json:"time_periods"`
	}

	TimePeriod struct {
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
	}

	ModifierGrouopIds struct {
		Ids []string
	}

	ModifierGroup struct {
		Id              string          `json:"id"`
		ExternalData    string          `json:"external_data"`
		Title           string          `json:"title"`
		Subtitle        string          `json:"subtitle"`
		QuentityInfo    QuantityInfo    `json:"quantity_info"`
		ModifierOptions ModifierOptions `json:"modifier_options"`
		DisplayType     string          `json:"display_type"`
	}

	ModifierOptions struct {
		Id   string `json:"id"`
		Type string `json:"type"`
	}
)

func (menu *Menu) ContainsCategory(categoryId string) bool {
	slice := menu.CategoryIds
	for _, v := range slice {
		if v == categoryId {
			return true
		}
	}
	return false
}

func (inputMenu *InputMenu) MakeMap() (map[string]Category, map[string]Item, map[string]ModifierGroup) {
	categoryMap := make(map[string]Category)
	entityMap := make(map[string]Item)
	modifierMap := make(map[string]ModifierGroup)
	for _, category := range inputMenu.Categories {
		categoryMap[category.ID] = category
	}
	for _, entity := range inputMenu.Items {
		entityMap[entity.ID] = entity
	}
	for _, modifier := range inputMenu.ModifierGroups {
		modifierMap[modifier.Id] = modifier
	}
	return categoryMap, entityMap, modifierMap
}
