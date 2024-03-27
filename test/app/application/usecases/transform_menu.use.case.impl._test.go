package usecases

import (
	"fmt"
	"github.com/metrofico/integracion-menu-lib/pkg/app/application/usecases"
	"github.com/metrofico/integracion-menu-lib/pkg/app/domain/models"
	"testing"
)

func TestTransformUseCase(t *testing.T) {

	t.Run("No se envía nada en la firma del método", func(t *testing.T) {
		useCase := usecases.NewTransformMenuUseCaseImpl()
		defer func() {
			if r := recover(); r == nil {
				fmt.Println("La función no ha lanzado un panic:", r)

				errorMsg := r.(string)
				t.Errorf(errorMsg)
			}
		}()
		useCase.Execute(nil)
	})

	t.Run("Validar si no se envía ningún menú", func(t *testing.T) {
		useCase := usecases.NewTransformMenuUseCaseImpl()
		input := models.InputMenu{}
		output := useCase.Execute(&input)

		if len(output) > 0 {
			t.Errorf("No debería traer ningún menú")
		}
	})

	t.Run("No se envía categorías", func(t *testing.T) {
		useCase := usecases.NewTransformMenuUseCaseImpl()
		input := models.InputMenu{
			Menus: []models.Menu{
				{
					ID:       "MenuId",
					Title:    "Título",
					Subtitle: "Subtítulo",
				},
			},
		}
		output := useCase.Execute(&input)

		if len(output) == 0 {
			t.Errorf("Debería llegar por lo menos un menú")
		}

		var categories = output[0].Categories
		if len(categories) > 0 {
			t.Errorf("No deberían llegar categorías")
		}
	})

	t.Run("No se envían items", func(t *testing.T) {
		useCase := usecases.NewTransformMenuUseCaseImpl()
		input := models.InputMenu{
			Menus: []models.Menu{
				{
					ID:       "MenuId",
					Title:    "Título",
					Subtitle: "Subtítulo",
				},
			},
			Categories: []models.Category{
				{
					ID:       "categoryId01",
					Title:    "Título",
					Subtitle: "Subtítulo",
				},
				{
					ID:       "categoryId02",
					Title:    "Título",
					Subtitle: "Subtítulo",
				},
			},
		}
		output := useCase.Execute(&input)

		if len(output) == 0 {
			t.Errorf("Debería llegar por lo menos un menú")
		}

		var categories = output[0].Categories
		if len(categories) > 2 {
			t.Errorf("No deberían llegar más de 2 categorías")
		}
	})

	t.Run("Validar categorías por tipo ITEM", func(t *testing.T) {
		useCase := usecases.NewTransformMenuUseCaseImpl()
		input := models.InputMenu{
			Menus: []models.Menu{
				{
					ID:       "MenuId",
					Title:    "Título",
					Subtitle: "Subtítulo",
				},
			},
			Categories: []models.Category{
				{
					ID:       "categoryId01",
					Title:    "Título",
					Subtitle: "Subtítulo",
					Entities: []models.Entity{
						{
							ID:   "item01",
							Type: "ITEM",
						},
						{
							ID:   "item02",
							Type: "NO_ITEM",
						},
					},
				},
				{
					ID:       "categoryId02",
					Title:    "Título",
					Subtitle: "Subtítulo",
				},
			},
			Items: []models.Item{
				{
					ID:    "item02",
					Title: "Título",
				},
			},
		}
		output := useCase.Execute(&input)

		if len(output) == 0 {
			t.Errorf("Debería llegar por lo menos un menú")
		}

		categories := output[0].Categories
		if len(categories) > 2 {
			t.Errorf("No deberían llegar más de 2 categorías")
		}

		items := categories[0].Items
		if len(items) > 0 {
			t.Errorf("No existen items que coincidan con esa categoría")
		}
	})

	t.Run("Conversión del precio de los items", func(t *testing.T) {
		useCase := usecases.NewTransformMenuUseCaseImpl()
		input := models.InputMenu{
			Menus: []models.Menu{
				{
					ID:       "MenuId",
					Title:    "Título",
					Subtitle: "Subtítulo",
				},
			},
			Categories: []models.Category{
				{
					ID:       "categoryId01",
					Title:    "Título",
					Subtitle: "Subtítulo",
					Entities: []models.Entity{
						{
							ID:   "item01",
							Type: "ITEM",
						},
						{
							ID:   "item02",
							Type: "NO_ITEM",
						},
					},
				},
				{
					ID:       "categoryId02",
					Title:    "Título",
					Subtitle: "Subtítulo",
				},
			},
			Items: []models.Item{
				{
					ID:    "item01",
					Title: "Título",
					PriceInfo: models.PriceInfo{
						Price:     29.99,
						CorePrice: 30.99,
					},
				},
			},
		}
		output := useCase.Execute(&input)

		if len(output) == 0 {
			t.Errorf("Debería llegar por lo menos un menú")
		}

		categories := output[0].Categories
		if len(categories) > 2 {
			t.Errorf("No deberían llegar más de 2 categorías")
		}

		items := categories[0].Items
		if len(items) == 0 {
			t.Errorf("Debería llegar el item con id: 'item01'")
		}

		item := items[0]
		if item.Price != 2999 {
			t.Errorf("El valor de conversión de precio es incorrecto, debería ser 2999")
		}

		if item.CorePrice != 3099 {
			t.Errorf("El valor de conversión de core es incorrecto, debería ser 3099")
		}
	})
}
