package userapp

import (
	"github.com/mobamoh/service-oas/app/oas"
	"github.com/mobamoh/service-oas/app/sdk/mid"
	"github.com/mobamoh/service-oas/business/domain/userbus"
	"github.com/mobamoh/service-oas/foundation/logger"
	"github.com/mobamoh/service-oas/foundation/web"
	"go.opentelemetry.io/otel/trace"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log     *logger.Logger
	Tracer  trace.Tracer
	UserBus *userbus.Business
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) error {

	apiPathPrefix := "/api/v1"
	mw := []oas.Middleware{
		mid.Otel(cfg.Tracer),
		mid.Logging(cfg.Log),
		mid.Errors(cfg.Log),
		mid.Metrics(),
		mid.Panics(),
	}

	api := newApp(cfg.UserBus)
	oaServer, err := oas.NewServer(api, oas.WithPathPrefix(apiPathPrefix), oas.WithMiddleware(mw...))
	if err != nil {
		return err
	}
	app.Handle("/", oaServer)
	return nil
}
