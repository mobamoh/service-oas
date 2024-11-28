package mux

import (
	"github.com/mobamoh/service-oas/foundation/logger"
	"github.com/mobamoh/service-oas/foundation/web"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Build  string
	Log    *logger.Logger
	Tracer trace.Tracer
}

// RouteAdder defines behavior that sets the routes to bind for an instance
// of the service.
type RouteAdder interface {
	Add(app *web.App, cfg Config) error
}

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(cfg Config, routeAdder RouteAdder) (http.Handler, error) {

	app := web.NewApp()

	if err := routeAdder.Add(app, cfg); err != nil {
		return nil, err
	}
	return app, nil
}
