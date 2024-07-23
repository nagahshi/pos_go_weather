package usecase_test

import (
	"testing"

	"github.com/nagahshi/pos_go_weather/internal/dto"

	usecase "github.com/nagahshi/pos_go_weather/internal/useCase"
)

// TestNewGetWeatherUseCase - testa a criação de um novo caso de uso para busca de clima
func TestNewGetWeatherUseCase(t *testing.T) {
	uc := usecase.NewGetWeatherUseCase("")

	d := dto.CEPOutput{}

	_, err := uc.Execute(d.ToWeatherInput())

	if err.Error() != "chave de acesso não informada" {
		t.Errorf("Erro esperado: chave de acesso não informada, recebido: %v", err)
	}
}
