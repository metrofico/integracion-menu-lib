package integracion

import (
	"encoding/json"
	"github.com/metrofico/integracion-menu-lib/pkg/app/application/services"
	services2 "github.com/metrofico/integracion-menu-lib/pkg/app/domain/ports/in/services"
	"log"
)

func main() {
	jsonRe := "{\n  \"menus\": [\n    {\n      \"id\": \"string\",\n      \"title\": \"Titulo del menu\",\n      \"subtitle\": \"Subtitulo del menu\",\n      \"category_ids\": [\n        \"1\"\n      ],\n      \"service_availability\": [\n        {\n          \"day_of_week\": \"monday\",\n          \"time_periods\": [\n            {\n              \"start_time\": \"00:01\",\n              \"end_time\": \"23:59\"\n            }\n          ]\n        }\n      ]\n    }\n  ],\n  \"categories\": [\n    {\n      \"id\": \"1\",\n      \"title\": \"Categoria 1 title\",\n      \"subtitle\": \"Subcategoria 1 title\",\n      \"entities\": [\n        {\n          \"id\": \"1\",\n          \"type\": \"ITEM\"\n        },\n        {\n          \"id\": \"2\",\n          \"type\": \"ITEM\"\n        }\n      ]\n    }\n  ],\n  \"items\": [\n    {\n      \"id\": \"1\",\n      \"external_data\": \"string1\",\n      \"title\": \"string1\",\n      \"description\": \"string1\",\n      \"image_url\": \"string1\",\n      \"price_info\": {\n        \"price\": 2999,\n        \"core_price\": 2999,\n        \"priced_by_unit\": {\n          \"measurement_type\": \"MEASUREMENT_TYPE_COUNT\",\n          \"length_unit\": 1,\n          \"weight_unit\": 1,\n          \"volume_unit\": 1\n        }\n      },\n      \"quantity_info\": {\n        \"quantity\": {\n          \"min_permitted\": 1,\n          \"max_permitted\": 1,\n          \"is_min_permitted_optional\": 1\n        }\n      },\n      \"modifier_group_ids\": {\n        \"ids\": [\n          \"1\",\n          \"2\"\n        ]\n      }\n    },\n    {\n      \"id\": \"2\",\n      \"external_data\": \"string2\",\n      \"title\": \"string1\",\n      \"description\": \"string1\",\n      \"image_url\": \"string1\",\n      \"price_info\": {\n        \"price\": 2999,\n        \"core_price\": 2999,\n        \"priced_by_unit\": {\n          \"measurement_type\": \"MEASUREMENT_TYPE_COUNT\",\n          \"length_unit\": 1,\n          \"weight_unit\": 1,\n          \"volume_unit\": 1\n        }\n      },\n      \"quantity_info\": {\n        \"quantity\": {\n          \"min_permitted\": 1,\n          \"max_permitted\": 1,\n          \"is_min_permitted_optional\": 1\n        }\n      },\n      \"modifier_group_ids\": {\n        \"ids\": [\n          \"3\"\n        ]\n      }\n    }\n  ],\n  \"modifier_groups\": [\n    {\n      \"id\": \"1\",\n      \"external_data\": \"string\",\n      \"title\": \"string\",\n      \"quantity_info\": {\n        \"quantity\": {\n          \"min_permitted\": 0,\n          \"max_permitted\": 0,\n          \"is_min_permitted_optional\": 0\n        }\n      },\n      \"modifier_options\": {\n        \"id\": \"producto1\",\n        \"type\": \"ITEM\"\n      },\n      \"display_type\": \"expanded\"\n    },\n    {\n      \"id\": \"2\",\n      \"external_data\": \"string\",\n      \"title\": \"string\",\n      \"quantity_info\": {\n        \"quantity\": {\n          \"min_permitted\": 0,\n          \"max_permitted\": 0,\n          \"is_min_permitted_optional\": 0\n        }\n      },\n      \"modifier_options\": {\n        \"id\": \"producto1\",\n        \"type\": \"ITEM\"\n      },\n      \"display_type\": \"expanded\"\n    },\n    {\n      \"id\": \"3\",\n      \"external_data\": \"string\",\n      \"title\": \"string\",\n      \"quantity_info\": {\n        \"quantity\": {\n          \"min_permitted\": 0,\n          \"max_permitted\": 0,\n          \"is_min_permitted_optional\": 0\n        }\n      },\n      \"modifier_options\": {\n        \"id\": \"producto1\",\n        \"type\": \"ITEM\"\n      },\n      \"display_type\": \"expanded\"\n    }\n  ]\n}"
	var service services2.MenuService
	service = services.NewMenuServiceImpl()
	contractOut := service.Transform(jsonRe)
	jsonget, err23 := json.Marshal(contractOut)
	if err23 != nil {
		log.Fatalln(err23.Error())
		return
	}
	println(string(jsonget))
}
