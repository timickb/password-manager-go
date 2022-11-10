package memory

import (
	"context"

	"github.com/timickb/password-manager/internal/common"
)

type Store struct {
	data map[string]string
}

func New() *Store {
	return &Store{}
}

func (s *Store) Open(ctx *context.Context) error {
	s.data = make(map[string]string)
	return nil
}

func (s *Store) Close(ctx *context.Context) error {
	if s.data == nil {
		return common.ErrStoreNotOpened
	}
	return nil
}

func (s *Store) SetItem(ctx *context.Context, key string, value string) error {
	s.data[key] = value
	return nil
}

func (s *Store) GetItem(ctx *context.Context, key string) (string, error) {
	if val, ok := s.data[key]; ok {
		return val, nil
	}
	return "", common.ErrNoSuchKey
}

func (s *Store) RemoveItem(ctx *context.Context, key string) error {
	if _, ok := s.data[key]; ok {
		delete(s.data, key)
		return nil
	}
	return common.ErrNoSuchKey
}
