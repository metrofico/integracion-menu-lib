package services

import (
	"shared-library/domain/entities"
)

type MenuService interface {
	Transform(input string) []entities.Menu
}
