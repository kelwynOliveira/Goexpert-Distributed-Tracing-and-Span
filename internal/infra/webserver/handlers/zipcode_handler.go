package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kelwynOliveira/Goexpert-Distributed-Tracing-and-Span/internal/entity"
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

func (h *ZipcodeHandler) ZipcodeHandler(w http.ResponseWriter, r *http.Request) {
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	_, span := h.TemplateData.OTELTracer.Start(ctx, h.TemplateData.RequestNameOTEL)
	defer span.End()

	var zipCode entity.ZipCodeForm
	err := json.NewDecoder(r.Body).Decode(&zipCode)
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

	zipstr := zipCode.Zipcode

	err = usecases.ValidateInput(zipstr)
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

	// Forward request to Service B
	resp, err := http.Post(h.TemplateData.ExternalCallURL+"?postcode="+zipstr, "", nil)
	if err != nil {
		http.Error(w, "failed to forward request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Forward Service B response to client
	for name, values := range resp.Header {
		w.Header()[name] = values
	}
	w.WriteHeader(resp.StatusCode)
	_, err = fmt.Fprint(w, resp.Body)
	if err != nil {
		fmt.Println("Error writing response:", err)
	}

}
