package usecases

import (
	"github.com/metrofico/integracion-menu-lib/domain/entities"
	"github.com/metrofico/integracion-menu-lib/domain/models"
)

type TransformMenuUseCase interface {
	Execute(input *models.InputMenu) []entities.Menu
}
