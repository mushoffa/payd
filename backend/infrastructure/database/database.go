package database

import (
	"context"
)

type DatabaseService interface {
	Insert(context.Context, string, ...any) (int, error)
	QueryOne(context.Context, string, ...any) (any, error)
	QueryMany(context.Context, string, ...any) (any, error)
	Exec(context.Context, string, ...any) (any, error)
	Lock(context.Context, int) error
	Unlock(context.Context, int) error
	Close()
}
