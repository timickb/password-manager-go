package memory_test

import (
	"context"
	"testing"

	"github.com/timickb/password-manager/internal/errors"
	"github.com/timickb/password-manager/internal/store/memory"
)

func TestNew(t *testing.T) {
	st := memory.New()
	if st == nil {
		t.Fatalf("store is nil")
	}
}

func TestOpen(t *testing.T) {
	st := memory.New()
	ctx := context.Background()
	err := st.Open(&ctx)

	if err != nil {
		t.Fatalf("store open err: %s", err.Error())
	}
}
func TestClose(t *testing.T) {
	ctx := context.Background()
	st := memory.New()

	_ = st.Open(&ctx)
	if err := st.Close(&ctx); err != nil {
		t.Fatalf("store close err: %s", err.Error())
	}

}

func TestCloseNotOpened(t *testing.T) {
	ctx := context.Background()
	st := memory.New()

	err := st.Close(&ctx)

	if err == nil {
		t.Fatalf("error expected")
	}
	if _, ok := err.(errors.ErrStoreNotOpened); !ok {
		t.Fatalf("wrong error type")
	}
}

func TestSetItem(t *testing.T) {
	ctx := context.Background()
	st := memory.New()

	_ = st.Open(&ctx)

	if err := st.SetItem(&ctx, "service", "value"); err != nil {
		t.Fatalf("item set error: %s", err.Error())
	}
}

func TestGetItem(t *testing.T) {
	ctx := context.Background()
	st := memory.New()

	_ = st.Open(&ctx)

	item := "secret"

	_ = st.SetItem(&ctx, "service", item)

	ret, err := st.GetItem(&ctx, "service")

	if err != nil {
		t.Fatalf("get item error: %s", err.Error())
	}

	if ret != item {
		t.Fatalf("items are different")
	}
}

func TestRemoveItem(t *testing.T) {
	ctx := context.Background()
	st := memory.New()

	_ = st.Open(&ctx)

	item := "secret"

	_ = st.SetItem(&ctx, "service", item)

	if err := st.RemoveItem(&ctx, "service"); err != nil {
		t.Fatalf("remove item error: %s", err.Error())
	}

	_, err := st.GetItem(&ctx, "service")

	if err == nil {
		t.Fatalf("error expected")
	}

	if _, ok := err.(errors.ErrNoSuchKey); !ok {
		t.Fatalf("wrong error type")
	}
}
