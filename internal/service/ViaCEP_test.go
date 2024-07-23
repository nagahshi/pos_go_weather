package service_test

import (
	"testing"

	"github.com/nagahshi/pos_go_weather/internal/service"
)

// TestNewViaCEPService - testa a criação de um novo serviço de busca de endereço pelo CEP
func TestNewViaCEPService(t *testing.T) {
	CEP := "01001000"
	srvc := service.NewViaCEPService(CEP)

	if srvc.CEP != CEP {
		t.Errorf("CEP esperado: %s, recebido: %s", CEP, srvc.CEP)
	}
}

// TestSearchViaCEP - testa a busca de endereço pelo CEP
func TestSearchViaCEP(t *testing.T) {
	// testando um CEP inválido
	srvc := service.NewViaCEPService("00000000")
	data, err := srvc.Search()
	if err == nil {
		t.Errorf("Esperado um erro, recebido: %v", data)
	}

	// testando um CEP válido
	srvc = service.NewViaCEPService("01001000")
	_, err = srvc.Search()
	if err != nil {
		t.Errorf("Erro não esperado para CEP válido, recebido: %v", err)
	}
}
