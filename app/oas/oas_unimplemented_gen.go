// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CreateUser implements createUser operation.
//
// Create a user in the system.
//
// POST /users
func (UnimplementedHandler) CreateUser(ctx context.Context, req OptUserCommand) (r *User, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteUserByID implements deleteUserByID operation.
//
// Deletes the user from the Sales System.
//
// DELETE /users/{userId}
func (UnimplementedHandler) DeleteUserByID(ctx context.Context, params DeleteUserByIDParams) error {
	return ht.ErrNotImplemented
}

// QueryUserByID implements queryUserByID operation.
//
// Returns the user details from Sales System.
//
// GET /users/{userId}
func (UnimplementedHandler) QueryUserByID(ctx context.Context, params QueryUserByIDParams) (r *User, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateUser implements updateUser operation.
//
// Update the user details in the system.
//
// PUT /users
func (UnimplementedHandler) UpdateUser(ctx context.Context, req OptUpdateUserReq) (r *User, _ error) {
	return r, ht.ErrNotImplemented
}
