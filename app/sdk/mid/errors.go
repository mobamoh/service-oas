package mid

import (
	"errors"
	"github.com/mobamoh/service-oas/app/sdk/errs"
	"github.com/mobamoh/service-oas/foundation/logger"
	"github.com/mobamoh/service-oas/foundation/otel"
	"github.com/ogen-go/ogen/middleware"

	"path"
)

func Errors(log *logger.Logger) middleware.Middleware {
	return func(
		req middleware.Request,
		next func(req middleware.Request) (middleware.Response, error),
	) (middleware.Response, error) {

		resp, err := next(req)
		if err == nil {
			return resp, nil
		}

		_, span := otel.AddSpan(req.Context, "app.sdk.mid.error")
		span.RecordError(err)
		defer span.End()

		var appErr *errs.Error
		if !errors.As(err, &appErr) {
			appErr = errs.Newf(errs.Internal, "Internal Server Error")
		}

		log.Error(req.Context, "handled error during request",
			"err", err,
			"source_err_file", path.Base(appErr.FileName),
			"source_err_func", path.Base(appErr.FuncName))

		if appErr.Code == errs.InternalOnlyLog {
			appErr = errs.Newf(errs.Internal, "Internal Server Error")
		}

		// Send the error to the transport package so the error can be
		// used as the response.

		return resp, appErr
	}
}
