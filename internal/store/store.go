package store

import (
	"context"
)

type Store interface {
	Open(*context.Context) error
	Close(*context.Context) error
	SetItem(*context.Context, string, string) error
	GetItem(*context.Context, string) (string, error)
	RemoveItem(*context.Context, string) error
}
