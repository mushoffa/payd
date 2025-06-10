package database

import (
	"context"
	"errors"
	"fmt"

	"payd/config"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
	client *pgxpool.Pool
}

func NewPostgres(config *config.Config) DatabaseService {
	db := config.Database

	pool, err := pgxpool.New(context.Background(), db.String())
	if err != nil {
		panic(err)
	}

	return &postgres{pool}
}

func (db *postgres) Insert(ctx context.Context, sql string, args ...any) (int, error) {
	var id int

	row, err := db.QueryOne(ctx, sql, args)
	if err != nil {
		return 0, err
	}

	row.(pgx.Row).Scan(&id)

	return id, nil
}

func (db *postgres) QueryOne(ctx context.Context, sql string, args ...any) (any, error) {
	row := db.client.QueryRow(ctx, sql, args)
	return row, nil
}

func (db *postgres) QueryMany(ctx context.Context, sql string, args ...any) (any, error) {
	rows, err := db.client.Query(ctx, sql, args...)
	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}
	return rows, nil
}

func (db *postgres) Exec(ctx context.Context, sql string, args ...any) (any, error) {
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
	db.Close()
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
