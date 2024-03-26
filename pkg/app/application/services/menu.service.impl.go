package services

import (
	"encoding/json"
	"shared-library/pkg/app/domain/entities"
	"shared-library/pkg/app/domain/models"
	"shared-library/pkg/app/domain/ports/in/usecases"
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
