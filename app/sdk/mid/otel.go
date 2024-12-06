package mid

import (
	"github.com/mobamoh/service-oas/foundation/otel"
	"github.com/ogen-go/ogen/middleware"
	"go.opentelemetry.io/otel/trace"
)

func Otel(tracer trace.Tracer) middleware.Middleware {
	return func(
		req middleware.Request,
		next func(req middleware.Request) (middleware.Response, error),
	) (middleware.Response, error) {
		ctx := otel.InjectTracing(req.Context, tracer)
		req.SetContext(ctx)
		return next(req)
	}
}
