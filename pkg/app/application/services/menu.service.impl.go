package services

import (
	"encoding/json"
	"github.com/metrofico/integracion-menu-lib/pkg/app/domain/entities"
	"github.com/metrofico/integracion-menu-lib/pkg/app/domain/models"
	"github.com/metrofico/integracion-menu-lib/pkg/app/domain/ports/in/usecases"
)

type MenuServiceImpl struct {
	TransformUseCase usecases.TransformMenuUseCase
}

func NewMenuServiceImpl(transform usecases.TransformMenuUseCase) *MenuServiceImpl {
	return &MenuServiceImpl{TransformUseCase: transform}
}

func (m MenuServiceImpl) Transform(input string) []entities.Menu {
	var contractIn models.InputMenu
	err := json.Unmarshal([]byte(input), &contractIn)
	if err != nil {
		panic(err.Error())
	}

	return m.TransformUseCase.Execute(&contractIn)
}
