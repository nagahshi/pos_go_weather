package usecase

import (
	"github.com/nagahshi/pos_go_weather/internal/dto"
	"github.com/nagahshi/pos_go_weather/internal/service"
)

type SearchCEPUseCase struct{}

func NewSearchCEPUseCase() *SearchCEPUseCase {
	return &SearchCEPUseCase{}
}

func (c *SearchCEPUseCase) Execute(CEP string) (output dto.CEPOutput, err error) {
	srvc := service.NewViaCEPService(CEP)

	response, err := srvc.Search()
	if err != nil {
		return output, err
	}

	return response, err
}
