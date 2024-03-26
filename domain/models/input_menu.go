package models

type (
	InputMenu struct {
		Menus      []Menu     `json:"menus"`
		Categories []Category `json:"categories"`
		Items      []Item     `json:"items"`
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
		ID           string       `json:"id"`
		ExternalData string       `json:"external_data"`
		Title        string       `json:"title"`
		Description  string       `json:"description"`
		ImageURL     string       `json:"image_url"`
		PriceInfo    PriceInfo    `json:"price_info"`
		QuantityInfo QuantityInfo `json:"quantity_info"`
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
		MinPermitted           int64 `json:"min_permitted"`
		MaxPermitted           int64 `json:"max_permitted"`
		IsMinPermittedOptional int64 `json:"is_min_permitted_optional"`
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
