// Package userdb contains user related CRUD functionality.
package userdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/mobamoh/service-oas/business/domain/common/db/gen"
	"github.com/mobamoh/service-oas/business/domain/userbus"
	"github.com/mobamoh/service-oas/business/sdk/sqldb"
	"github.com/mobamoh/service-oas/foundation/logger"
	"net/mail"
)

// Store manages the set of APIs for prospect database access.
type Store struct {
	log     *logger.Logger
	pool    *sqldb.Postgres
	queries *gen.Queries
}

// NewStore constructs the api for data access.
func NewStore(log *logger.Logger, pool *sqldb.Postgres) *Store {
	return &Store{
		log:     log,
		pool:    pool, // pool needed to implement transactions, ex: tx, err := s.pool.DB.Begin(ctx)
		queries: gen.New(pool.DB),
	}
}

func (s *Store) Create(ctx context.Context, usr userbus.User) error {
	if err := s.queries.CreateUser(ctx, toDBUser(usr)); err != nil {
		return fmt.Errorf("db: %w", err)
	}
	return nil
}

func (s *Store) Update(ctx context.Context, usr userbus.User) error {
	//TODO implement me
	panic("implement me")
}

func (s *Store) Delete(ctx context.Context, userID uuid.UUID) error {
	if err := s.queries.DeleteUser(ctx, userID); err != nil {
		return fmt.Errorf("db: %w", err)
	}
	return nil
}

func (s *Store) QueryByID(ctx context.Context, userID uuid.UUID) (userbus.User, error) {
	dbUsr, err := s.queries.QueryUserByID(ctx, userID)

	if err != nil {
		if errors.Is(err, sqldb.ErrDBNotFound) {
			return userbus.User{}, fmt.Errorf("db: %w", userbus.ErrNotFound)
		}
		return userbus.User{}, fmt.Errorf("db: %w", err)
	}
	return toBusUser(dbUsr)
}

// QueryByEmail gets the specified user from the database by email.
func (s *Store) QueryByEmail(ctx context.Context, email mail.Address) (userbus.User, error) {
	dbUsr, err := s.queries.QueryUserByEmail(ctx, email.Address)

	if err != nil {
		if errors.Is(err, sqldb.ErrDBNotFound) {
			return userbus.User{}, fmt.Errorf("db: %w", userbus.ErrNotFound)
		}
		return userbus.User{}, fmt.Errorf("db: %w", err)
	}
	return toBusUser(dbUsr)
}
