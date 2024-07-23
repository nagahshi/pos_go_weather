package web

import (
	"net/http"

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

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	// output, err := h.CheckHealthyUseCase.Execute()
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// err = json.NewEncoder(w).Encode(output)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}
