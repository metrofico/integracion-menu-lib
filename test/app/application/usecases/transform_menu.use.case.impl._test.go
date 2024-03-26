package usecases

import (
	"fmt"
	"shared-library/pkg/app/application/usecases"
	"testing"
)

func TestTransformUseCase(t *testing.T) {
	t.Run("No se envía nada en la firma del método", func(t *testing.T) {
		useCase := usecases.NewTransformMenuUseCaseImpl()
		defer func() {
			if r := recover(); r == nil {
				fmt.Println("La función no ha lanzado un panic:", r)

				errorMsg := r.(string)
				t.Errorf(errorMsg)
			}
		}()
		useCase.Execute(nil)
	})
}
