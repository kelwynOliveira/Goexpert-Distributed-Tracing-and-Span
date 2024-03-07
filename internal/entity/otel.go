package entity

import "go.opentelemetry.io/otel/trace"

type TemplateData struct {
	Title              string
	Name               string
	WeatherApiKey      string
	ExternalCallMethod string
	ExternalCallURL    string
	RequestNameOTEL    string
	OTELTracer         trace.Tracer
}
