package usecases

import (
	"shared-library/domain/entities"
	"shared-library/domain/models"
)

type TransformMenuUseCase interface {
	Execute(input *models.InputMenu) []entities.Menu
}
