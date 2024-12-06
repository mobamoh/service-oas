package commands

import (
	"context"
	"fmt"
	"github.com/mobamoh/service-oas/business/sdk/migrate"
	"github.com/mobamoh/service-oas/business/sdk/sqldb"
	"time"
)

// Seed loads test data into the database.
func Seed(cfg sqldb.Config) error {
	ctx := context.Background()
	db, err := sqldb.Open(ctx, cfg)
	if err != nil {
		return fmt.Errorf("connect database: %w", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := migrate.Seed(ctx, db); err != nil {
		return fmt.Errorf("seed database: %w", err)
	}

	fmt.Println("seed data complete")
	return nil
}
