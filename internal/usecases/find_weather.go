package usecases

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/kelwynOliveira/Goexpert-Distributed-Tracing-and-Span/internal/entity"
)

type FindByCityNameUseCaseInterface interface {
	GetWeather(city string) (*entity.Forecast, error)
}

type FindByCityNameUseCase struct {
	APIKey string
}

func NewFindByCityNameUseCase(
	apiKey string,
) *FindByCityNameUseCase {
	return &FindByCityNameUseCase{
		APIKey: apiKey,
	}
}

func (uc *FindByCityNameUseCase) GetWeather(city string) (*entity.Forecast, error) {
	var weather entity.Forecast
	weatherAPIURL := "https://api.weatherapi.com/v1/current.json?key=" + uc.APIKey + "&q=" + url.QueryEscape(city) + "&aqi=no"

	request, err := http.Get(weatherAPIURL)
	if err != nil {
		return nil, err
	}

	result, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(result, &weather)
	if err != nil {
		return nil, err
	}

	return &weather, err
}
