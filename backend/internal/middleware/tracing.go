package middleware

import (
	
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/AkshatVerma087/go-boilerplate/backend/internal/server"
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type TracingMiddleware struct {
	server *server.Server
	nrApp *newrelic.Application
}

func NewTracingMiddleware(s *server.Server, nrApp *newrelic.Application) *TracingMiddleware {
	return &TracingMiddleware{
		server: s,
		nrApp: nrApp,
	}
}


func (tm *TracingMiddleware) NewRelicMiddleware() echo.MiddlewareFunc {
	if tm.nrApp == nil {

		return func(next echo.HandlerFunc) echo.HandlerFunc {
			return next
		}
	}
		
	return nrecho.Middleware(tm.nrApp)
}


func (tm *TracingMiddleware) EnhanceTracing() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			txn := newrelic.FromContext(c.Request().Context())
			if txn == nil {
				return next(c)
			}

			txn.AddAttribute("service.name", tm.server.Config.Observability.ServiceName)
			txn.AddAttribute("service.enviornment", tm.server.Config.Observability.Environment)
			txn.AddAttribute("http.real_ip", c.RealIP())
			txn.AddAttribute("http.user_agent", c.Request().UserAgent())


			if requestID := GetRequestID(c); requestID !=  "" {
				txn.AddAttribute("request.id", requestID)
			} 

			if userID := c.Get("user_id"); userID != nil {
				if userIDStr, ok := userID.(string); ok {
					txn.AddAttribute("user.id", userIDStr)
				}
			}

			err := next(c)

			if err != nil {
				txn.NoticeError(nrpkgerrors.Wrap(err))
			}

			txn.AddAttribute("http.status_code", c.Response().Status)

			return err

		}
	}
}
