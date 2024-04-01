package entities

import "github.com/metrofico/integracion-menu-lib/pkg/app/domain/models"

type Item struct {
	Id            string                   `json:"id" bson:"id"`
	ExternalData  string                   `json:"external_data" bson:"external_data"`
	Title         string                   `json:"title" bson:"title"`
	Description   string                   `json:"description" bson:"description"`
	Image         string                   `json:"image_url"`
	PriceInfo                              /*Herencia por composición*/
	QuantityInfo                           /*Herencia por composición*/
	ModifierGroup map[string]ModifierGroup `json:"modifier_group"`
}

type (
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

type (
	PriceInfo struct {
		Price       int64        `json:"price" bson:"price"`
		CorePrice   int64        `json:"core_price" bson:"core_price"`
		PriceByUnit PricedByUnit `json:"priced_by_unit" bson:"priced_by_unit"`
	}

	PricedByUnit struct {
		MeasurementType string `json:"measurement_type" bson:"measurement_type"`
		LengthUnit      int64  `json:"length_unit" bson:"length_unit"`
		WeightUnit      int64  `json:"weight_unit" bson:"weight_unit"`
		VolumeUnit      int64  `json:"volume_unit" bson:"volume_unit"`
	}
)

type (
	QuantityInfo struct {
		Quantity Quantity `json:"quantity" bson:"volume_unit"`
	}

	Quantity struct {
		MinPermitted           int `json:"min_permitted" bson:"min_permitted"`
		MaxPermitted           int `json:"max_permitted" bson:"max_permitted"`
		IsMinPermittedOptional int `json:"is_min_permitted_optional" bson:"is_min_permitted_optional"`
	}
)

func NewItem() Item {
	return Item{}
}

func (item *Item) Mapper(input models.Item, mapModifiers map[string]models.ModifierGroup) {
	modifierFinal := make(map[string]ModifierGroup)
	for _, ids := range input.ModifierGrouopIds.Ids {
		if modifier, ok3 := mapModifiers[ids]; ok3 {
			modifierFinal[modifier.Id] = ModifierGroup{
				Id:           modifier.Id,
				ExternalData: modifier.ExternalData,
				Title:        modifier.Title,
				Subtitle:     modifier.Subtitle,
				QuentityInfo: QuantityInfo{Quantity: Quantity{
					MinPermitted: modifier.QuentityInfo.Quantity.MinPermitted,
					MaxPermitted: modifier.QuentityInfo.Quantity.MaxPermitted,
				}},
				ModifierOptions: ModifierOptions{
					Id:   modifier.ModifierOptions.Id,
					Type: modifier.ModifierOptions.Type,
				},
				DisplayType: modifier.DisplayType,
			}
		}
	}
	price := PriceInfo{
		Price:     int64(input.PriceInfo.Price * 100),
		CorePrice: int64(input.PriceInfo.CorePrice * 100),
		PriceByUnit: PricedByUnit{
			MeasurementType: input.PriceInfo.PricedByUnit.MeasurementType,
			LengthUnit:      input.PriceInfo.PricedByUnit.LengthUnit,
			WeightUnit:      input.PriceInfo.PricedByUnit.WeightUnit,
			VolumeUnit:      input.PriceInfo.PricedByUnit.VolumeUnit,
		},
	}

	quantity := QuantityInfo{
		Quantity: Quantity{
			MinPermitted:           input.QuantityInfo.Quantity.MinPermitted,
			MaxPermitted:           input.QuantityInfo.Quantity.MaxPermitted,
			IsMinPermittedOptional: input.QuantityInfo.Quantity.IsMinPermittedOptional,
		},
	}

	item.Id = input.ID
	item.Title = input.Title
	item.Description = input.Description
	item.PriceInfo = price
	item.QuantityInfo = quantity
	item.ModifierGroup = modifierFinal
}
