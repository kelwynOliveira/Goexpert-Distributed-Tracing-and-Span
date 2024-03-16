package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kelwynOliveira/Goexpert-Distributed-Tracing-and-Span/internal/entity"
	"github.com/kelwynOliveira/Goexpert-Distributed-Tracing-and-Span/internal/usecases"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type WebClimateHandlerInterface interface {
	TemperatureHandler(w http.ResponseWriter, r *http.Request)
}

type WebClimateHandler struct {
	FindClimateByCityNameUseCase usecases.FindByCityNameUseCaseInterface
	TemplateData                 *entity.TemplateData
}

// NewServer creates a new server instance
func NewWebClimateHandler(
	findByCityNameUC usecases.FindByCityNameUseCaseInterface,
	TemplateData *entity.TemplateData,
) *WebClimateHandler {
	return &WebClimateHandler{
		FindClimateByCityNameUseCase: findByCityNameUC,
		TemplateData:                 TemplateData,
	}
}

func (h *WebClimateHandler) TemperatureHandler(w http.ResponseWriter, r *http.Request) {
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	_, span := h.TemplateData.OTELTracer.Start(ctx, h.TemplateData.RequestNameOTEL+" GET")
	defer span.End()

	postcode := r.URL.Query().Get("postcode")

	err := usecases.ValidateInput(postcode)
	if err != nil {
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(msg)
		return
	}

	location, err := usecases.GetViaCEP(postcode)
	if err != nil {
		err = errors.New("can not find zipcode")
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(msg)
		return
	}

	climate, err := h.FindClimateByCityNameUseCase.GetWeather(location.City)
	if err != nil {
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(msg)
		return
	}

	fahrenheit, kelvin := convertTemperature(climate.Current.TempC)

	response := entity.Temperature{
		City:       location.City,
		Celcius:    float32(climate.Current.TempC),
		Fahrenheit: float32(fahrenheit),
		Kelvin:     float32(kelvin),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func convertTemperature(celcius float64) (float64, float64) {
	fahrenheit := celcius*1.8 + 32
	kelvin := celcius + 273.15

	return fahrenheit, kelvin
}
