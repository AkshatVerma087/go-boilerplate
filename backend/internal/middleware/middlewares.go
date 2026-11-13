package middleware

import (
	"github.com/AkshatVerma087/go-boilerplate/backend/internal/server"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Middlewares struct {
	Global          *GlobalMiddlewares
	Auth            *AuthMiddleware
	ContextEnhancer *ContextEnhancer
	Tracing         *TracingMiddleware
	RateLimit       *RateLimitMiddleware
}

func NewMiddlewares(s *server.Server) *Middlewares {
	var nrApp *newrelic.Application 
	if s.LoggerService != nil {
		nrApp = s.LoggerService.GetApplication()
	}

	return &Middlewares{
		Global: NewGlobalMiddlewares(s),
		Auth: NewAuthMiddleware(s),
		ContextEnhancer: NewContextEnhancer(s),
		Tracing: NewTracingMiddleware(s, nrApp),
		RateLimit: NewRateLimitMiddleware(s),
	}
}