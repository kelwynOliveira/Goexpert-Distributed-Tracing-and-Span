package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kelwynOliveira/Goexpert-Distributed-Tracing-and-Span/internal/entity"
	"github.com/kelwynOliveira/Goexpert-Distributed-Tracing-and-Span/internal/infra/database"
	"github.com/kelwynOliveira/Goexpert-Distributed-Tracing-and-Span/internal/usecases"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type ZipcodeHandler struct {
	TemplateData *entity.TemplateData
}

// NewServer creates a new server instance
func NewZipcodeHandler(TemplateData *entity.TemplateData) *ZipcodeHandler {
	return &ZipcodeHandler{
		TemplateData: TemplateData,
	}
}

func (h *ZipcodeHandler) SaveZipcodeHandler(w http.ResponseWriter, r *http.Request) {
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	_, span := h.TemplateData.OTELTracer.Start(ctx, h.TemplateData.RequestNameOTEL+" POST")
	defer span.End()

	// if err := r.ParseForm(); err != nil {
	// 	fmt.Fprintf(w, "ParseForm() err: %v", err)
	// }
	// zipstr := r.FormValue("cep")

	var zipCode entity.ZipCodeForm
	err := json.NewDecoder(r.Body).Decode(&zipCode)
	if err != nil {
		fmt.Println(err)
	}

	zipstr := zipCode.Zipcode

	zipcode, err := usecases.CreateZipCode(zipstr)
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

	// File
	err = database.CreateFile(zipstr)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(zipcode)

}

func (h *ZipcodeHandler) GetZipcodeHandler(w http.ResponseWriter, r *http.Request) {
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	_, span := h.TemplateData.OTELTracer.Start(ctx, h.TemplateData.RequestNameOTEL+" GET")
	defer span.End()

	zipstr := database.ReadFile()

	zipcode, err := usecases.CreateZipCode(zipstr)
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(zipcode)
}
