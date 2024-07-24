package web

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-chi/chi/v5"

	useCase "github.com/nagahshi/pos_go_weather/internal/useCase"
)

type WeatherHandler struct {
	SearchCEPUseCase  useCase.SearchCEPUseCase
	GetWeatherUseCase useCase.GetWeatherUseCase
}

func NewWeatherHandler(searchCEPUseCase useCase.SearchCEPUseCase, getWeatherUseCase useCase.GetWeatherUseCase) *WeatherHandler {
	return &WeatherHandler{
		SearchCEPUseCase:  searchCEPUseCase,
		GetWeatherUseCase: getWeatherUseCase,
	}
}

// GetWeather - busca de clima pelo CEP
func (wh *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	var CEP = chi.URLParam(r, "cep")
	var re *regexp.Regexp = regexp.MustCompile("[0-9]+")

	CEP = strings.Join(re.FindAllString(CEP, -1), "")
	if len(CEP) != 8 {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	// SearchCEP - busca de endereço pelo CEP
	outputCEP, err := wh.SearchCEPUseCase.Execute(CEP)
	if err != nil {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	log.Println(outputCEP)

	// GetWeather - busca de clima pelo endereço encontrado em SearchCEP
	outputWeather, err := wh.GetWeatherUseCase.Execute(outputCEP.ToWeatherInput())
	if err != nil {
		http.Error(w, "can not find location to weather", http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(outputWeather)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
}
