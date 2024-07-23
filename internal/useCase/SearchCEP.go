package usecase

import (
	"github.com/nagahshi/pos_go_weather/internal/dto"
)

type SearchCEPUseCase struct{}

func NewSearchCEPUseCase() *SearchCEPUseCase {
	return &SearchCEPUseCase{}
}

func (c *SearchCEPUseCase) Execute() (output dto.CEPOutput, err error) {
	_ = dto.CEPInput{}

	return output, err
}
