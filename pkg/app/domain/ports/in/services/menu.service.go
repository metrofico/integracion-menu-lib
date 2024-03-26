package services

import (
	"shared-library/pkg/app/domain/entities"
)

type MenuService interface {
	Transform(input string) []entities.Menu
}
