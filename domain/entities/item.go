package entities

import "github.com/metrofico/integracion-menu-lib/domain/models"

type Item struct {
	Id           string `json:"id" bson:"id"`
	ExternalData string `json:"external_data" bson:"external_data"`
	Title        string `json:"title" bson:"title"`
	Description  string `json:"description" bson:"description"`
	PriceInfo           /*Herencia por composición*/
	QuantityInfo        /*Herencia por composición*/
}

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
		Quantity Quantity
	}

	Quantity struct {
		MinPermitted           int64 `json:"min_permitted" bson:"min_permitted"`
		MaxPermitted           int64 `json:"max_permitted" bson:"max_permitted"`
		IsMinPermittedOptional bool  `json:"is_min_permitted_optional" bson:"is_min_permitted_optional"`
	}
)

func NewItem() Item {
	return Item{}
}

func (item *Item) Mapper(input models.Item) {
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
			IsMinPermittedOptional: input.QuantityInfo.Quantity.IsMinPermittedOptional == 1,
		},
	}

	item.Id = input.ID
	item.Title = input.Title
	item.Description = input.Description
	item.PriceInfo = price
	item.QuantityInfo = quantity
}
