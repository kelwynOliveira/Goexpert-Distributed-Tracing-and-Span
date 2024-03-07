package usecases

import (
	"errors"
	"regexp"

	"github.com/kelwynOliveira/Goexpert-Distributed-Tracing-and-Span/internal/entity"
)

func CreateZipCode(zipcode string) (*entity.ZipCodeForm, error) {
	err := validateInput(zipcode)
	if err != nil {
		return nil, err
	}

	return &entity.ZipCodeForm{
		Zipcode: zipcode,
	}, nil
}

func validateInput(zipCode string) error {
	matched, err := regexp.MatchString(`\b\d{8}\b`, zipCode)
	if !matched || err != nil {
		return errors.New("invalid zipcode")
	}

	return nil
}
