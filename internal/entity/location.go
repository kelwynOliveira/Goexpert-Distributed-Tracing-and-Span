package entity

import (
	"errors"
	"regexp"
)

type ZipCodeForm struct {
	Zipcode string `json:"cep"`
}

type ViaCEP struct {
	Zipcode      string `json:"cep"`
	AddressLine1 string `json:"logradouro"`
	AddressLine2 string `json:"complemento"`
	Neighborhood string `json:"bairro"`
	City         string `json:"localidade"`
	State        string `json:"uf"`
	IBGE         string `json:"ibge"`
	GIA          string `json:"gia"`
	Area         string `json:"ddd"`
	SIAFI        string `json:"siafi"`
}

func NewZipcode(zipstr string) (*ZipCodeForm, error) {
	zipcode := &ZipCodeForm{Zipcode: zipstr}
	err := zipcode.validateInput(zipstr)
	if err != nil {
		return nil, err
	}
	return zipcode, nil
}

func (z *ZipCodeForm) validateInput(zipCode string) error {
	matched, err := regexp.MatchString(`\b\d{8}\b`, zipCode)
	if !matched || err != nil {
		return errors.New("invalid zipcode")
	}

	return nil
}
