// Package migrate contains the database schema, migrations and seeding data.
package migrate

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/mobamoh/service-oas/business/sdk/migrate/migrations"
	"github.com/mobamoh/service-oas/business/sdk/sqldb"
	"io/fs"

	"github.com/jackc/pgx/v5/stdlib"

	"github.com/pressly/goose/v3"
)

var (
	//go:embed seeds/seed.sql
	seedDoc string
)

// Migrate attempts to bring the database up to date with the migrations
// defined in this package.
func Migrate(ctx context.Context, pool *sqldb.Postgres) error {

	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("status check database: %w", err)
	}

	return migrateFS(pool, migrations.FS, ".")
}

func migrateFS(pool *sqldb.Postgres, fs fs.FS, dir string) error {
	goose.SetBaseFS(fs)
	defer func() {
		goose.SetBaseFS(nil)
	}()

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("migrate database: %w", err)
	}

	db := stdlib.OpenDBFromPool(pool.DB)

	if err := goose.Up(db, dir); err != nil {
		return fmt.Errorf("migrate database: %w", err)
	}
	return nil
}

// Seed runs the seed document defined in this package against db. The queries
// are run in a transaction and rolled back if any fail.
func Seed(ctx context.Context, pool *sqldb.Postgres) (err error) {
	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("status check database: %w", err)
	}

	tx, err := pool.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// tx, err := db.Begin()
	// if err != nil {
	// 	return err
	// }

	// defer func() {
	// 	if errTx := tx.Rollback(ctx); errTx != nil {
	// 		if errors.Is(errTx, sql.ErrTxDone) {
	// 			return
	// 		}

	// 		err = fmt.Errorf("rollback: %w", errTx)
	// 		return
	// 	}
	// }()

	if _, err := tx.Exec(ctx, seedDoc); err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit: %w", err)
	}

	return nil
}
