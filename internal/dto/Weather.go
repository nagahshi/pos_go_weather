package dto

type WeatherInput struct {
	Logradouro string
	Bairro     string
	UF         string
	CIDADE     string
	Latitude   string
	Longitude  string
}

type WeatherOutput struct {
	C float64 `json:"celsius"`
	F float64 `json:"fahrenheit"`
	K float64 `json:"kelvin"`
}
