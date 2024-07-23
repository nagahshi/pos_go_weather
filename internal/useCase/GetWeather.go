package usecase

import (
	"github.com/nagahshi/pos_go_weather/internal/dto"
)

type GetWeatherUseCase struct{}

func NewGetWeatherUseCase() *GetWeatherUseCase {
	return &GetWeatherUseCase{}
}

func (c *GetWeatherUseCase) Execute() (output dto.CEPOutput, err error) {
	_ = dto.CEPInput{}

	return output, err
}
