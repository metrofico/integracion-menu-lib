package main

import (
	"encoding/json"
	"log"
	"shared-library/application/services"
	"shared-library/application/usecases"
	services2 "shared-library/domain/ports.in/services"
)

func main() {
	jsonRe := "{ \"menus\": [ { \"id\": \"menu1\", \"title\": \"Menú del Día\", \"subtitle\": \"Saborea nuestras especialidades\", \"category_ids\": [ \"categoryId1\", \"categoryId2\" ], \"service_availability\": [ { \"day_of_week\": \"monday\", \"time_periods\": [ { \"start_time\": \"11:00\", \"end_time\": \"15:00\" }, { \"start_time\": \"18:00\", \"end_time\": \"22:00\" } ] }, { \"day_of_week\": \"tuesday\", \"time_periods\": [ { \"start_time\": \"11:00\", \"end_time\": \"15:00\" }, { \"start_time\": \"18:00\", \"end_time\": \"22:00\" } ] } ] } ], \"categories\": [ { \"id\": \"categoryId1\", \"title\": \"Platos Principales\", \"subtitle\": \"Deliciosos platos principales\", \"entities\": [ { \"id\": \"producto1\", \"type\": \"ITEM\" }, { \"id\": \"producto2\", \"type\": \"ITEM\" } ] }, { \"id\": \"categoryId2\", \"title\": \"Bebidas\", \"subtitle\": \"Refrescos y bebidas\", \"entities\": [ { \"id\": \"producto3\", \"type\": \"ITEM\" }, { \"id\": \"producto4\", \"type\": \"ITEM\" } ] } ], \"items\": [ { \"id\": \"producto1\", \"external_data\": \"producto1-data\", \"title\": \"Hamburguesa Clásica\", \"description\": \"Una hamburguesa jugosa con queso, lechuga y tomate.\", \"image_url\": \"url-de-la-imagen-1\", \"price_info\": { \"price\": 799, \"core_price\": 799, \"priced_by_unit\": { \"measurement_type\": \"MEASUREMENT_TYPE_COUNT\", \"length_unit\": 0, \"weight_unit\": 0, \"volume_unit\": 0 } }, \"quantity_info\": { \"quantity\": { \"min_permitted\": 1, \"max_permitted\": 10, \"is_min_permitted_optional\": 0 } } }, { \"id\": \"producto2\", \"external_data\": \"producto2-data\", \"title\": \"Ensalada César\", \"description\": \"Una fresca ensalada con pollo a la parrilla, aderezo César y crutones.\", \"image_url\": \"url-de-la-imagen-2\", \"price_info\": { \"price\": 899, \"core_price\": 899, \"priced_by_unit\": { \"measurement_type\": \"MEASUREMENT_TYPE_COUNT\", \"length_unit\": 0, \"weight_unit\": 0, \"volume_unit\": 0 } }, \"quantity_info\": { \"quantity\": { \"min_permitted\": 1, \"max_permitted\": 10, \"is_min_permitted_optional\": 0 } } }, { \"id\": \"producto3\", \"external_data\": \"producto3-data\", \"title\": \"Refresco de Cola\", \"description\": \"Una bebida refrescante de cola en lata.\", \"image_url\": \"url-de-la-imagen-3\", \"price_info\": { \"price\": 199, \"core_price\": 199, \"priced_by_unit\": { \"measurement_type\": \"MEASUREMENT_TYPE_COUNT\", \"length_unit\": 0, \"weight_unit\": 0, \"volume_unit\": 0 } }, \"quantity_info\": { \"quantity\": { \"min_permitted\": 1, \"max_permitted\": 10, \"is_min_permitted_optional\": 0 } } }, { \"id\": \"producto4\", \"external_data\": \"producto4-data\", \"title\": \"Agua Mineral\", \"description\": \"Una botella de agua mineral natural.\", \"image_url\": \"url-de-la-imagen-4\", \"price_info\": { \"price\": 149, \"core_price\": 149, \"priced_by_unit\": { \"measurement_type\": \"MEASUREMENT_TYPE_COUNT\", \"length_unit\": 0, \"weight_unit\": 0, \"volume_unit\": 0 } }, \"quantity_info\": { \"quantity\": { \"min_permitted\": 1, \"max_permitted\": 10, \"is_min_permitted_optional\": 0 } } } ] }"
	var service services2.MenuService
	useCase := usecases.NewTransformMenuUseCaseImpl()
	service = services.NewMenuServiceImpl(useCase)
	contractOut := service.Transform(jsonRe)

	jsonget, err23 := json.Marshal(contractOut)
	if err23 != nil {
		log.Fatalln(err23.Error())
		return
	}
	println(string(jsonget))
}
