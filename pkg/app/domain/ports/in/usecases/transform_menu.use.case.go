package usecases

import (
	"shared-library/pkg/app/domain/entities"
	"shared-library/pkg/app/domain/models"
)

type TransformMenuUseCase interface {
	Execute(input *models.InputMenu) []entities.Menu
}
