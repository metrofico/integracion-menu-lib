package usecases

import (
	"github.com/metrofico/integracion-menu-lib/pkg/app/domain/entities"
	"github.com/metrofico/integracion-menu-lib/pkg/app/domain/models"
)

type TransformMenuUseCase interface {
	Execute(input *models.InputMenu) []entities.Menu
}
