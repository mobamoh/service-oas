package commands

import (
	"context"
	"errors"
	"fmt"
	"github.com/mobamoh/service-oas/business/sdk/migrate"
	"github.com/mobamoh/service-oas/business/sdk/sqldb"
	"time"
)

// ErrHelp provides context that help was given.
var ErrHelp = errors.New("provided help")

// Migrate creates the schema in the database.
func Migrate(cfg sqldb.Config) error {
	// cfg.Migrate = true
	ctx := context.Background()
	db, err := sqldb.Open(ctx, cfg)
	if err != nil {
		return fmt.Errorf("connect database: %w", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := migrate.Migrate(ctx, db); err != nil {
		return fmt.Errorf("migrate database: %w", err)
	}

	fmt.Println("migrations complete")
	return nil
}
