package mid

import (
	"errors"
	"fmt"
	"github.com/mobamoh/service-oas/app/sdk/errs"
	"github.com/mobamoh/service-oas/foundation/logger"
	"github.com/ogen-go/ogen/middleware"
	"time"
)

func Logging(log *logger.Logger) middleware.Middleware {
	return func(
		req middleware.Request,
		next func(req middleware.Request) (middleware.Response, error),
	) (middleware.Response, error) {

		now := time.Now()
		r := req.Raw
		path := r.URL.Path
		if r.URL.RawQuery != "" {
			path = fmt.Sprintf("%s?%s", path, r.URL.RawQuery)
		}
		ctx := req.Context
		log.Info(ctx, "request started", "method", r.Method, "path", path, "remoteaddr", r.RemoteAddr)

		resp, err := next(req)

		var statusCode = errs.OK
		if err != nil {
			statusCode = errs.Internal

			var v *errs.Error
			if errors.As(err, &v) {
				statusCode = v.Code
			}
		}

		log.Info(ctx, "request completed", "method", r.Method, "path", path, "remoteaddr", r.RemoteAddr,
			"statuscode", statusCode, "since", time.Since(now).String())

		return resp, err
	}
}
