package userbus

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/mobamoh/service-oas/foundation/logger"
	"github.com/mobamoh/service-oas/foundation/otel"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
	"time"
)

// Set of error variables for CRUD operations.
var (
	ErrNotFound              = errors.New("user not found")
	ErrUniqueEmail           = errors.New("email is not unique")
	ErrAuthenticationFailure = errors.New("authentication failed")
)

// Storer interface declares the behavior this package needs to persist and
// retrieve data.
type Storer interface {
	Create(ctx context.Context, usr User) error
	Update(ctx context.Context, usr User) error
	Delete(ctx context.Context, userID uuid.UUID) error
	QueryByID(ctx context.Context, userID uuid.UUID) (User, error)
	QueryByEmail(ctx context.Context, email mail.Address) (User, error)
}

// Business manages the set of APIs for user access.
type Business struct {
	log    *logger.Logger
	storer Storer
}

// NewBusiness constructs a user business API for use.
func NewBusiness(log *logger.Logger, storer Storer) *Business {
	return &Business{
		log:    log,
		storer: storer,
	}
}

// Create adds a new user to the system.
func (b *Business) Create(ctx context.Context, nu NewUser) (User, error) {
	ctx, span := otel.AddSpan(ctx, "business.userbus.create")
	defer span.End()

	hash, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, fmt.Errorf("generatefrompassword: %w", err)
	}

	now := time.Now()

	usr := User{
		ID:           uuid.New(),
		Name:         nu.Name,
		Email:        nu.Email,
		PasswordHash: hash,
		Roles:        nu.Roles,
		Department:   nu.Department,
		Enabled:      true,
		DateCreated:  now,
		DateUpdated:  now,
	}

	if err := b.storer.Create(ctx, usr); err != nil {
		return User{}, fmt.Errorf("create: %w", err)
	}

	return usr, nil
}

// Update modifies information about a user.
func (b *Business) Update(ctx context.Context, usr User, uu UpdateUser) (User, error) {
	//TODO implement me
	panic("implement me")
}

// Delete removes the specified user.
func (b *Business) Delete(ctx context.Context, userID uuid.UUID) error {
	ctx, span := otel.AddSpan(ctx, "business.userbus.delete")
	defer span.End()

	if err := b.storer.Delete(ctx, userID); err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}

// QueryByID finds the user by the specified ID.
func (b *Business) QueryByID(ctx context.Context, userID uuid.UUID) (User, error) {
	ctx, span := otel.AddSpan(ctx, "business.userbus.querybyid")
	defer span.End()

	user, err := b.storer.QueryByID(ctx, userID)
	if err != nil {
		return User{}, fmt.Errorf("query: userID[%s]: %w", userID, err)
	}

	return user, nil
}

// QueryByEmail finds the user by a specified user email.
func (b *Business) QueryByEmail(ctx context.Context, email mail.Address) (User, error) {
	ctx, span := otel.AddSpan(ctx, "business.userbus.querybyemail")
	defer span.End()

	user, err := b.storer.QueryByEmail(ctx, email)
	if err != nil {
		return User{}, fmt.Errorf("query: email[%s]: %w", email, err)
	}

	return user, nil
}
