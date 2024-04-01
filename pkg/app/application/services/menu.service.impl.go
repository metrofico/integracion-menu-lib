package services

import (
	"encoding/json"
	usecases2 "github.com/metrofico/integracion-menu-lib/pkg/app/application/usecases"
	"github.com/metrofico/integracion-menu-lib/pkg/app/domain/entities"
	"github.com/metrofico/integracion-menu-lib/pkg/app/domain/models"
	"github.com/metrofico/integracion-menu-lib/pkg/app/domain/ports/in/usecases"
)

type MenuServiceImpl struct {
	TransformUseCase usecases.TransformMenuUseCase
}

func NewMenuServiceImpl() *MenuServiceImpl {
	useCase := usecases2.NewTransformMenuUseCaseImpl()
	return &MenuServiceImpl{TransformUseCase: useCase}
}

func (m MenuServiceImpl) Transform(input string) []entities.Menu {
	var contractIn models.InputMenu
	err := json.Unmarshal([]byte(input), &contractIn)
	if err != nil {
		panic(err.Error())
	}

	return m.TransformUseCase.Execute(&contractIn)
}
