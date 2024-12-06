package userapp

import (
	"context"
	"errors"
	"github.com/mobamoh/service-oas/app/oas"
	"github.com/mobamoh/service-oas/app/sdk/errs"
	"github.com/mobamoh/service-oas/business/domain/userbus"
)

type app struct {
	userBus *userbus.Business
}

func newApp(userBus *userbus.Business) *app {
	return &app{
		userBus: userBus,
	}
}

func (a *app) CreateUser(ctx context.Context, req oas.OptUserCommand) (*oas.User, error) {

	usrReq, ok := req.Get()
	if !ok {
		return nil, errs.New(errs.InvalidArgument, errors.New("invalid request"))
	}
	nu, err := toBusNewUser(usrReq)
	if err != nil {
		return nil, errs.New(errs.InvalidArgument, err)
	}

	usr, err := a.userBus.Create(ctx, nu)
	if err != nil {
		return nil, errs.Newf(errs.Internal, "create: usr[%+v]: %s", usr, err)
	}

	return toAppUser(usr), nil
}

func (a *app) UpdateUser(ctx context.Context, req oas.OptUpdateUserReq) (*oas.User, error) {
	//TODO implement me
	panic("implement me")
}

func (a *app) DeleteUserByID(ctx context.Context, params oas.DeleteUserByIDParams) error {
	userID, err := toBusUserID(params.UserId)
	if err != nil {
		return err
	}

	if err := a.userBus.Delete(ctx, userID); err != nil {
		return errs.Newf(errs.Internal, "delete: userID[%s]: %s", userID, err)
	}
	return nil
}

func (a *app) QueryUserByID(ctx context.Context, params oas.QueryUserByIDParams) (*oas.User, error) {
	userID, err := toBusUserID(params.UserId)
	if err != nil {
		return nil, err
	}

	usr, err := a.userBus.QueryByID(ctx, userID)
	if err != nil {
		return nil, errs.Newf(errs.Internal, "query: %s", err)
	}

	return toAppUser(usr), nil
}
