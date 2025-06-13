package database

import (
	"context"
	"errors"
	"fmt"

	"payd/config"
	"payd/infrastructure/trace/embedded"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	semconv "go.opentelemetry.io/otel/semconv/v1.32.0"
)

type postgres struct {
	embedded.Monitor
	client *pgxpool.Pool
}

func NewPostgres(config *config.Config) DatabaseService {
	db := config.Database

	pool, err := pgxpool.New(context.Background(), db.String())
	if err != nil {
		panic(err)
	}

	pg := &postgres{
		client: pool,
	}

	pg.Init(
		semconv.DBSystemNamePostgreSQL,
		semconv.ServerAddressKey.String(db.URL),
		semconv.ServerPortKey.String(fmt.Sprintf("%d", db.Port)),
	)

	return pg
}

func (db *postgres) Insert(ctx context.Context, sql string, args ...any) (int, error) {
	var id int

	_, span := db.Trace(ctx, "Database.Insert")
	defer span.End()

	row, err := db.QueryOne(ctx, sql, args)
	if err != nil {
		return 0, err
	}

	row.(pgx.Row).Scan(&id)

	return id, nil
}

func (db *postgres) QueryOne(ctx context.Context, sql string, args ...any) (any, error) {
	_, span := db.Trace(ctx, "Database.QueryOne")
	defer span.End()

	row := db.client.QueryRow(ctx, sql, args)
	return row, nil
}

func (db *postgres) QueryMany(ctx context.Context, sql string, args ...any) (any, error) {
	_, span := db.Trace(ctx, "Database.QueryMany")
	defer span.End()

	rows, err := db.client.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (db *postgres) Exec(ctx context.Context, sql string, args ...any) (any, error) {
	_, span := db.Trace(ctx, "Database.Exec")
	defer span.End()

	res, err := db.client.Exec(ctx, sql, args)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (db *postgres) Lock(ctx context.Context, id int) error {
	const sql = `select pg_try_advisory_lock=$1`

	locked, err := db.mutex(ctx, sql, id)
	if err != nil {
		return err
	}

	if !locked {
		return errors.New("cannot acquire lock")
	}

	return nil
}

func (db *postgres) Unlock(ctx context.Context, id int) error {
	const sql = `select pg_advisory_unlock=$1`

	locked, err := db.mutex(ctx, sql, id)
	if err != nil {
		return err
	}

	if !locked {
		return errors.New("cannot release lock")
	}

	return nil
}

func (db *postgres) Close() {
	db.client.Close()
}

func (db *postgres) mutex(ctx context.Context, sql string, id int) (bool, error) {
	var locked bool

	res, err := db.QueryOne(ctx, sql, id)
	if err != nil {
		return false, err
	}

	if err := res.(pgx.Row).Scan(&locked); err != nil {
		return false, err
	}

	return locked, nil
}
