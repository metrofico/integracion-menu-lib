package usecases

import (
	"github.com/metrofico/integracion-menu-lib/pkg/app/domain/entities"
	"github.com/metrofico/integracion-menu-lib/pkg/app/domain/models"
	"github.com/metrofico/integracion-menu-lib/pkg/app/domain/ports/in/usecases"
)

type TransformMenuUseCaseImpl struct {
}

func NewTransformMenuUseCaseImpl() usecases.TransformMenuUseCase {
	return &TransformMenuUseCaseImpl{}
}

func (t TransformMenuUseCaseImpl) Execute(input *models.InputMenu) []entities.Menu {
	if input == nil {
		panic("El menú es obligatorio")
	}

	var menus []entities.Menu
	mapCategory, mapItem, mapModifiers := input.MakeMap()
	for _, inputMenu := range input.Menus {
		menu := entities.NewMenu()
		menu.Id = inputMenu.ID
		menu.Title = inputMenu.Title
		menu.Subtitle = inputMenu.Subtitle

		var categories []entities.Category
		for _, inputCategory := range input.Categories {
			_, ok := mapCategory[inputCategory.ID]
			if ok {
				category := entities.NewCategory(
					inputCategory.ID,
					inputCategory.Title,
					inputCategory.Subtitle,
				)
				for _, inputItem := range inputCategory.Entities {
					if inputItem.Type != "ITEM" {
						continue
					}
					if entityItem, ok2 := mapItem[inputItem.ID]; ok2 {
						item := entities.NewItem()
						item.Mapper(entityItem, mapModifiers)

						category.Items = append(category.Items, item)
					}
				}
				categories = append(categories, category)
			}
		}

		menu.Categories = categories
		menus = append(menus, menu)
	}

	return menus
}

func filterItemsByCategory(items []models.Item, categoryEntities []models.Entity) []models.Item {
	var filtered []models.Item
	for _, entity := range categoryEntities {
		for _, menu := range items {
			if menu.ID == entity.ID && entity.Type == "ITEM" {
				filtered = append(filtered, menu)
			}
		}
	}
	return filtered
}
