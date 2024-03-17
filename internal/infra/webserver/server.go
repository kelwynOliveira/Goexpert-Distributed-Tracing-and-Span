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
		// fileServer := http.FileServer(http.Dir("template"))
		// router.Handle("/*", fileServer)
		// router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// 	tpl := template.Must(template.New("index.html").ParseFS(templateContent, "template/index.html"))
		// 	err := tpl.Execute(w, we.TemplateData)
		// 	if err != nil {
		// 		http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
		// 		return
		// 	}
		// })

		router.Post("/", handlers.NewZipcodeHandler(we.TemplateData).ZipcodeHandler)
		// router.Get("/cep", handlers.NewZipcodeHandler(we.TemplateData).GetZipcodeHandler)
	case "serviceB":
		climate := usecases.NewFindByCityNameUseCase(we.TemplateData.WeatherApiKey)
		temperatureHandler := handlers.NewWebClimateHandler(climate, we.TemplateData)
		router.Get("/", temperatureHandler.TemperatureHandler)
	}
	// router.Get("/", handlers.)
	return router
}
