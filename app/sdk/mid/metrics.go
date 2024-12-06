package mid

import (
	"github.com/mobamoh/service-oas/app/sdk/metrics"
	"github.com/ogen-go/ogen/middleware"
)

func Metrics() middleware.Middleware {
	return func(
		req middleware.Request,
		next func(req middleware.Request) (middleware.Response, error),
	) (middleware.Response, error) {

		ctx := metrics.Set(req.Context)

		resp, err := next(req)

		n := metrics.AddRequests(ctx)

		if n%1000 == 0 {
			metrics.AddGoroutines(ctx)
		}

		if err != nil {
			metrics.AddErrors(ctx)
		}

		return resp, err
	}
}
