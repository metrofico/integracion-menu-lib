package usecases

import (
	"github.com/metrofico/integracion-menu-lib/domain/entities"
	"github.com/metrofico/integracion-menu-lib/domain/models"
	"github.com/metrofico/integracion-menu-lib/domain/ports.in/usecases"
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

	for _, inputMenu := range input.Menus {
		menu := entities.NewMenu()
		menu.Id = inputMenu.ID
		menu.Title = inputMenu.Title
		menu.Subtitle = inputMenu.Subtitle

		var categories []entities.Category
		for _, inputCategory := range input.Categories {
			hasCategory := inputMenu.ContainsCategory(inputCategory.ID)
			if hasCategory {
				category := entities.NewCategory(
					inputCategory.ID,
					inputCategory.Title,
					inputCategory.Subtitle,
				)

				for _, inputItem := range filterItemsByCategory(input.Items, inputCategory.Entities) {
					item := entities.NewItem()
					item.Mapper(inputItem)
					category.Items = append(category.Items, item)
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
			if menu.ID == entity.ID {
				filtered = append(filtered, menu)
			}
		}
	}
	return filtered
}
