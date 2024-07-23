package main

import (
	"log"
	"os"

	webSrv "github.com/nagahshi/pos_go_weather/internal/infra/web/server"
	usecase "github.com/nagahshi/pos_go_weather/internal/useCase"

	"github.com/nagahshi/pos_go_weather/internal/infra/web"
)

func main() {
	webserver := webSrv.NewWebServer("8080")

	key := os.Getenv("WEATHER_API_KEY")
	if key == "" {
		log.Fatal("chave de consulta [WEATHER_API_KEY] não encontrada")
		return
	}

	weatherHandler := web.NewWeatherHandler(
		*usecase.NewSearchCEPUseCase(),
		*usecase.NewGetWeatherUseCase(key),
	)

	webserver.AddHandler("/cep/{cep}", weatherHandler.GetWeather)

	println("Starting web server on port:", ":8080")
	webserver.Start()
}
