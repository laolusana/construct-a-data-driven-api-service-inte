package main

import (
	"encoding/json"
	"errors"
)

// Service represents a single API service
type Service struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// ServiceIntegrator represents the data-driven API service integrator
type ServiceIntegrator struct {
	Services []Service `json:"services"`
	Config   Config    `json:"config"`
}

// Config represents the configuration for the integrator
type Config struct {
	BaseURL string `json:"base_url"`
	Timeout int    `json:"timeout"`
}

// ServiceIntegration represents the integration details for a single service
type ServiceIntegration struct {
	ServiceID string                 `json:"service_id"`
	Endpoint string                 `json:"endpoint"`
	Method   string                 `json:"method"`
	Query    map[string]string      `json:"query"`
	Headers  map[string]string      `json:"headers"`
	Body     json.RawMessage        `json:"body"`
	Auth     ServiceIntegrationAuth `json:"auth"`
}

// ServiceIntegrationAuth represents the authentication details for a service integration
type ServiceIntegrationAuth struct {
	Type   string `json:"type"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

// Error represents a custom error type
type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e *Error) Error() string {
	return e.Message
}

// NewError returns a new error instance
func NewError(message string, code int) error {
	return &Error{Message: message, Code: code}
}

// validate returns an error if the service integrator is invalid
func (si *ServiceIntegrator) validate() error {
	if len(si.Services) == 0 {
		return NewError("at least one service is required", 400)
	}
	for _, service := range si.Services {
		if service.ID == "" {
			return NewError("service id is required", 400)
		}
		if service.URL == "" {
			return NewError("service url is required", 400)
		}
	}
	return nil
}