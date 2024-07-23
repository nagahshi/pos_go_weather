package main

import (
	webSrv "github.com/nagahshi/pos_go_weather/internal/infra/web/server"
	usecase "github.com/nagahshi/pos_go_weather/internal/useCase"

	"github.com/nagahshi/pos_go_weather/internal/infra/web"
)

func main() {
	webserver := webSrv.NewWebServer("8080")

	healthyHandler := web.NewWeatherHandler(
		*usecase.NewSearchCEPUseCase(),
		*usecase.NewGetWeatherUseCase(),
	)

	webserver.AddHandler("/cep/{cep}", healthyHandler.GetWeather)

	println("Starting web server on port:", ":8080")
	webserver.Start()
}
