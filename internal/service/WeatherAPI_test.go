package service_test

import (
	"testing"

	"github.com/nagahshi/pos_go_weather/internal/service"
)

// TestNewWeatherAPI - testa a criação de um novo serviço de busca de clima
func TestNewWeatherAPI(t *testing.T) {
	srvc := service.NewWeatherAPIService("", "local")

	if srvc.Localidade != "local" {
		t.Errorf("Localidade esperado: %s, recebido: %s", "local", srvc.Localidade)
	}
}

// TestSearchWeatherAPI - testa a busca de clima
func TestSearchWeatherAPI(t *testing.T) {
	// testando uma localidade inválida
	srvc := service.NewWeatherAPIService("", "local")
	data, err := srvc.Search()
	if err == nil {
		t.Errorf("Esperado um erro, recebido: %v", data)
	}

	if err.Error() != "chave de acesso não informada" {
		t.Errorf("Erro esperado: chave de acesso não informada, recebido: %v", err)
	}
}
