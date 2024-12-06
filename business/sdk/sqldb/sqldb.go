package sqldb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/url"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Set of error variables for CRUD operations.
var (
	ErrDBNotFound        = sql.ErrNoRows
	ErrDBDuplicatedEntry = errors.New("duplicated entry")
	ErrUndefinedTable    = errors.New("undefined table")
)

// Config is the required properties to use the database.
type Config struct {
	User         string
	Password     string
	Host         string
	Name         string
	Schema       string
	MaxIdleConns int
	MaxOpenConns int
	DisableTLS   bool
	MigrateSeed  bool
}

type Postgres struct {
	DB *pgxpool.Pool
}

var (
	pgInstance *Postgres
	pgOnce     sync.Once
)

// Open knows how to open a database connection based on the configuration.
func Open(ctx context.Context, cfg Config) (*Postgres, error) {

	pgOnce.Do(func() {
		// db, err := pgxpool.New(ctx, u.String())
		db, err := GetDbPool(ctx, cfg)
		if err != nil {
			log.Fatal("Failed to create a config, error: ", err)
		}

		pgInstance = &Postgres{db}
	})

	return pgInstance, nil
}

func (pg *Postgres) Ping(ctx context.Context) error {
	return pg.DB.Ping(ctx)
}

func (pg *Postgres) Close() {
	pg.DB.Close()
}

func MakeConnectionString(cfg Config) string {
	sslMode := "require"
	if cfg.DisableTLS {
		sslMode = "disable"
	}

	q := make(url.Values)
	q.Set("sslmode", sslMode)
	q.Set("timezone", "utc")
	if cfg.Schema != "" {
		q.Set("search_path", cfg.Schema)
	}

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.User, cfg.Password),
		Host:     cfg.Host,
		Path:     cfg.Name,
		RawQuery: q.Encode(),
	}
	return u.String()
}

func GetDbPool(ctx context.Context, cfg Config) (*pgxpool.Pool, error) {
	// Set up a new pool with the custom types and the config.
	config, err := pgxpool.ParseConfig(MakeConnectionString(cfg))
	if err != nil {
		return nil, err
	}
	dbPool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	// If enabled, don't register custom types on migrations & seeds.
	if !cfg.MigrateSeed {
		// Collect the custom data types once, store them in memory, and register them for every future connection.
		customTypes, err := getCustomDataTypes(ctx, dbPool)
		if err != nil {
			return nil, err
		}

		config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
			for _, t := range customTypes {
				conn.TypeMap().RegisterType(t)
			}
			return nil
		}
	}

	// Immediately close the old pool and open a new one with the new config.
	dbPool.Close()
	dbPool, err = pgxpool.NewWithConfig(ctx, config)
	return dbPool, err
}

// Any custom DB types made with CREATE TYPE need to be registered with pgx.
// https://github.com/kyleconroy/sqlc/issues/2116
// https://stackoverflow.com/questions/75658429/need-to-update-psql-row-of-a-composite-type-in-golang-with-jack-pgx
// https://pkg.go.dev/github.com/jackc/pgx/v5/pgtype
func getCustomDataTypes(ctx context.Context, pool *pgxpool.Pool) ([]*pgtype.Type, error) {
	// Get a single connection just to load type information.
	conn, err := pool.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return nil, err
	}

	dataTypeNames := []string{
		//"enum_type",
	}

	var typesToRegister []*pgtype.Type
	for _, typeName := range dataTypeNames {
		dataType, err := conn.Conn().LoadType(ctx, typeName)
		if err != nil {
			return nil, fmt.Errorf("failed to load type %s: %v", typeName, err)
		}
		// You need to register only for this connection too, otherwise the array type will look for the register element type.
		conn.Conn().TypeMap().RegisterType(dataType)
		typesToRegister = append(typesToRegister, dataType)
	}
	return typesToRegister, nil
}
