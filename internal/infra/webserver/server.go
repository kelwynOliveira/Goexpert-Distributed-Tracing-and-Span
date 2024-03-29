package webserver

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kelwynOliveira/Goexpert-Distributed-Tracing-and-Span/internal/entity"
	"github.com/kelwynOliveira/Goexpert-Distributed-Tracing-and-Span/internal/infra/webserver/handlers"
	"github.com/kelwynOliveira/Goexpert-Distributed-Tracing-and-Span/internal/usecases"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// go:embed template/*
// var templateContent embed.FS

type Webserver struct {
	TemplateData *entity.TemplateData
}

// NewServer creates a new server instance
func NewServer(templateData *entity.TemplateData) *Webserver {
	return &Webserver{
		TemplateData: templateData,
	}
}

// createServer creates a new server instance with go chi router
func (we *Webserver) CreateServer() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(60 * time.Second))
	// promhttp
	router.Handle("/metrics", promhttp.Handler())

	switch we.TemplateData.Name {
	case "serviceA":
		router.Post("/", handlers.NewZipcodeHandler(we.TemplateData).ZipcodeHandler)
	case "serviceB":
		climate := usecases.NewFindByCityNameUseCase(we.TemplateData.WeatherApiKey)
		temperatureHandler := handlers.NewWebClimateHandler(climate, we.TemplateData)
		router.HandleFunc("/", temperatureHandler.TemperatureHandler)
	}
	return router
}
