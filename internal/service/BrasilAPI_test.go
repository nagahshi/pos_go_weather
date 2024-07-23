package service_test

import (
	"testing"

	"github.com/nagahshi/pos_go_weather/internal/service"
)

// TestNewBrasilAPIService - testa a criação de um novo serviço de busca de endereço pelo CEP
func TestNewBrasilAPIService(t *testing.T) {
	CEP := "01001000"
	srvc := service.NewBrasilAPIService(CEP)

	if srvc.CEP != CEP {
		t.Errorf("CEP esperado: %s, recebido: %s", CEP, srvc.CEP)
	}
}

// TestSearchBrasilAPI - testa a busca de endereço pelo CEP
func TestSearchBrasilAPI(t *testing.T) {
	// testando um CEP inválido
	srvc := service.NewBrasilAPIService("00000000")
	data, err := srvc.Search()
	if err == nil {
		t.Errorf("Esperado um erro, recebido: %v", data)
	}

	// testando um CEP válido
	srvc = service.NewBrasilAPIService("01001000")
	_, err = srvc.Search()
	if err != nil {
		t.Errorf("Erro não esperado para CEP válido, recebido: %v", err)
	}
}
