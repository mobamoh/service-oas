// Package all binds all the routes into the specified app.
package all

import (
	"github.com/mobamoh/service-oas/app/domain/checkapp"
	"github.com/mobamoh/service-oas/app/domain/userapp"
	"github.com/mobamoh/service-oas/app/sdk/mux"
	"github.com/mobamoh/service-oas/business/domain/userbus"
	"github.com/mobamoh/service-oas/business/domain/userbus/stores/userdb"
	"github.com/mobamoh/service-oas/foundation/web"
)

// Routes constructs the add value which provides the implementation
// of RouteAdder for specifying what routes to bind to this instance.
func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouterAdder interface.
func (add) Add(app *web.App, cfg mux.Config) error {

	// Construct the business domain packages we need here so we are using the
	// sames instances for the different set of domain apis.
	noteBus := userbus.NewBusiness(cfg.Log, userdb.NewStore(cfg.Log, cfg.Pool))

	checkapp.Routes(app, checkapp.Config{
		Build: cfg.Build,
		Log:   cfg.Log,
	})

	if err := userapp.Routes(app, userapp.Config{
		Log:     cfg.Log,
		UserBus: noteBus,
		Tracer:  cfg.Tracer,
	}); err != nil {
		return err
	}

	return nil
}
