// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"context"
	"net/http"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
)

// handleCreateUserRequest handles createUser operation.
//
// Create a user in the system.
//
// POST /users
func (s *Server) handleCreateUserRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createUser"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/users"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "CreateUser",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		attrOpt := metric.WithAttributeSet(labeler.AttributeSet())

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributeSet(labeler.AttributeSet()))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "CreateUser",
			ID:   "createUser",
		}
	)
	request, close, err := s.decodeCreateUserRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *User
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "CreateUser",
			OperationSummary: "Create User",
			OperationID:      "createUser",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = OptCreateUserReq
			Params   = struct{}
			Response = *User
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.CreateUser(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.CreateUser(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeCreateUserResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleDeleteUserByIDRequest handles deleteUserByID operation.
//
// Deletes the user from the Sales System.
//
// DELETE /users/{userId}
func (s *Server) handleDeleteUserByIDRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deleteUserByID"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/users/{userId}"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DeleteUserByID",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		attrOpt := metric.WithAttributeSet(labeler.AttributeSet())

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributeSet(labeler.AttributeSet()))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "DeleteUserByID",
			ID:   "deleteUserByID",
		}
	)
	params, err := decodeDeleteUserByIDParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *DeleteUserByIDNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "DeleteUserByID",
			OperationSummary: "Delete user",
			OperationID:      "deleteUserByID",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "userId",
					In:   "path",
				}: params.UserId,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = DeleteUserByIDParams
			Response = *DeleteUserByIDNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackDeleteUserByIDParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.DeleteUserByID(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.DeleteUserByID(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeDeleteUserByIDResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleQueryUserByIDRequest handles queryUserByID operation.
//
// Returns the user details from Sales System.
//
// GET /users/{userId}
func (s *Server) handleQueryUserByIDRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("queryUserByID"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/users/{userId}"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "QueryUserByID",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		attrOpt := metric.WithAttributeSet(labeler.AttributeSet())

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributeSet(labeler.AttributeSet()))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "QueryUserByID",
			ID:   "queryUserByID",
		}
	)
	params, err := decodeQueryUserByIDParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *User
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "QueryUserByID",
			OperationSummary: "Return user details",
			OperationID:      "queryUserByID",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "userId",
					In:   "path",
				}: params.UserId,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = QueryUserByIDParams
			Response = *User
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackQueryUserByIDParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.QueryUserByID(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.QueryUserByID(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeQueryUserByIDResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUpdateUserRequest handles updateUser operation.
//
// Update the user details in the system.
//
// PUT /users
func (s *Server) handleUpdateUserRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateUser"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/users"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "UpdateUser",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		attrOpt := metric.WithAttributeSet(labeler.AttributeSet())

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributeSet(labeler.AttributeSet()))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "UpdateUser",
			ID:   "updateUser",
		}
	)
	request, close, err := s.decodeUpdateUserRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *User
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "UpdateUser",
			OperationSummary: "Update User",
			OperationID:      "updateUser",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = OptUpdateUserReq
			Params   = struct{}
			Response = *User
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.UpdateUser(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.UpdateUser(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUpdateUserResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}