package usecase_test

import (
	"testing"

	usecase "github.com/nagahshi/pos_go_weather/internal/useCase"
)

// TestNewSearchCEPUseCase - testa a criação de um novo caso de uso para busca de endereço pelo CEP
func TestNewSearchCEPUseCase(t *testing.T) {
	uc := usecase.NewSearchCEPUseCase()

	_, err := uc.Execute("01001000")
	if err != nil {
		t.Errorf("Erro não esperado para CEP válido, recebido: %v", err)
	}

	_, err = uc.Execute("00000000")
	if err == nil {
		t.Errorf("Esperado um erro, recebido: %v", err)
	}
}
