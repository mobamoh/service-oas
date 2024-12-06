package checkapp

import (
	"github.com/mobamoh/service-oas/foundation/logger"
	"github.com/mobamoh/service-oas/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Build string
	Log   *logger.Logger
	//DB    *sqlx.DB
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "/v1"

	api := newApp(cfg.Build, cfg.Log)

	app.HandleFunc(version+"/readiness", api.readiness)
	app.HandleFunc(version+"/liveness", api.liveness)
}
