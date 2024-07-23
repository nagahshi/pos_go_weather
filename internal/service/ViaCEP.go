package service

import (
	"errors"
	"io"
	"strings"

	"github.com/nagahshi/pos_go_weather/internal/dto"
	"github.com/valyala/fastjson"
)

type ViaCEP struct {
	CEP string
}

func NewViaCEPService(CEP string) *ViaCEP {
	return &ViaCEP{
		CEP: CEP,
	}
}

func getUFName(uf string) string {
	switch uf {
	case "RO":
		return "rondônia"
	case "AC":
		return "acre"
	case "AM":
		return "amazonas"
	case "RR":
		return "roraima"
	case "PA":
		return "pará"
	case "AP":
		return "amapá"
	case "TO":
		return "tocantins"
	case "MA":
		return "maranhão"
	case "PI":
		return "piauí"
	case "CE":
		return "ceará"
	case "RN":
		return "rio grande do norte"
	case "PB":
		return "paraíba"
	case "PE":
		return "pernambuco"
	case "AL":
		return "alagoas"
	case "SE":
		return "sergipe"
	case "BA":
		return "bahia"
	case "MG":
		return "minas gerais"
	case "ES":
		return "espírito santo"
	case "RJ":
		return "rio de janeiro"
	case "SP":
		return "são paulo"
	case "PR":
		return "paraná"
	case "SC":
		return "santa catarina"
	case "RS":
		return "rio grande do sul"
	case "MS":
		return "mato grosso do sul"
	case "MT":
		return "mato grosso"
	case "GO":
		return "goiás"
	case "DF":
		return "distrito federal"
	}

	return "brazil"
}

func (c *ViaCEP) Search() (CEPOutput dto.CEPOutput, err error) {
	var client = getNewClient()

	// realizo pesquisas cada um em sua rotina
	resp, err := client.Get("https://viacep.com.br/ws/" + c.CEP + "/json/")
	if err != nil {
		return CEPOutput, errors.New("ocorreu um erro, ao buscar informações: " + err.Error())
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return CEPOutput, errors.New("ocorreu um erro, ao ler informações")
	}

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		var p fastjson.Parser
		v, err := p.Parse(string(respBody))
		if err != nil {
			return CEPOutput, errors.New("ocorreu um erro, ao tratar informações")
		}

		CEPOutput.Logradouro = string(v.GetStringBytes("logradouro"))
		CEPOutput.Bairro = string(v.GetStringBytes("bairro"))
		CEPOutput.UF = getUFName(strings.ToUpper(string(v.GetStringBytes("uf"))))
		CEPOutput.CIDADE = string(v.GetStringBytes("localidade"))

		if CEPOutput.CIDADE == "" && CEPOutput.UF == "brazil" {
			return CEPOutput, errors.New("CEP não encontrado")
		}

		return CEPOutput, nil
	}

	return CEPOutput, errors.New("ocorreu um erro, ao buscar informações: bad request via CEP")
}
