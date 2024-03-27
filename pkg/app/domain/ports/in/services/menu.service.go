package services

import (
	"github.com/metrofico/integracion-menu-lib/pkg/app/domain/entities"
)

type MenuService interface {
	Transform(input string) []entities.Menu
}
