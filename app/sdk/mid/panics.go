package mid

import (
	"github.com/mobamoh/service-oas/app/sdk/errs"
	"github.com/mobamoh/service-oas/app/sdk/metrics"
	"github.com/ogen-go/ogen/middleware"

	"runtime/debug"
)

func Panics() middleware.Middleware {
	return func(
		req middleware.Request,
		next func(req middleware.Request) (middleware.Response, error),
	) (resp middleware.Response, err error) {
		defer func() {
			if rec := recover(); rec != nil {
				trace := debug.Stack()
				err = errs.Newf(errs.InternalOnlyLog, "PANIC [%v] TRACE[%s]", rec, string(trace))

				metrics.AddPanics(req.Context)
			}
		}()

		return next(req)
	}
}
