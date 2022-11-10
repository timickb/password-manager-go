package api_test

import (
	"context"
	"errors"
	"testing"

	"github.com/timickb/password-manager/internal/common"
	"github.com/timickb/password-manager/internal/config"
	"github.com/timickb/password-manager/internal/crypto"
	"github.com/timickb/password-manager/internal/installer"
	"github.com/timickb/password-manager/internal/store/memory"
	"github.com/timickb/password-manager/pkg/api"
)

func TestNew(t *testing.T) {
	cfg, err := config.New()
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	store := memory.New()
	crypto := crypto.StubEncrypter{}
	ctx := context.Background()

	pm, err := api.New(&ctx, cfg, store, &crypto)
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	if pm.IsReady() {
		t.Fatalf("Ready is true")
	}
}

func TestSetup(t *testing.T) {
	cfg, err := config.New()
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	store := memory.New()
	crypto := crypto.StubEncrypter{}
	ctx := context.Background()

	pm, err := api.New(&ctx, cfg, store, &crypto)
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	ins := installer.StubInstaller{}

	err = pm.Setup(&ctx, &ins)
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	if !pm.IsReady() {
		t.Fatalf("Ready is false")
	}
}

func TestSetAndRead(t *testing.T) {
	cfg, err := config.New()
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	store := memory.New()
	crypto := crypto.StubEncrypter{}
	ctx := context.Background()

	pm, err := api.New(&ctx, cfg, store, &crypto)
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	ins := installer.StubInstaller{}

	err = pm.Setup(&ctx, &ins)
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	err = pm.Set(&ctx, "key", "value")
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	val, err := pm.Read(&ctx, "key")
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	if val != "value" {
		t.Fatalf("value is invalid: %s", val)
	}
}

func TestReadDoesNotExist(t *testing.T) {
	cfg, err := config.New()
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	store := memory.New()
	crypto := crypto.StubEncrypter{}
	ctx := context.Background()

	pm, err := api.New(&ctx, cfg, store, &crypto)
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	ins := installer.StubInstaller{}

	err = pm.Setup(&ctx, &ins)
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	_, err = pm.Read(&ctx, "key")
	if !errors.Is(err, common.ErrNoSuchKey) {
		t.Fatalf("wrong error type")
	}
}

func TestDelete(t *testing.T) {
	cfg, err := config.New()
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	store := memory.New()
	crypto := crypto.StubEncrypter{}
	ctx := context.Background()

	pm, err := api.New(&ctx, cfg, store, &crypto)
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	ins := installer.StubInstaller{}

	err = pm.Setup(&ctx, &ins)
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	err = pm.Set(&ctx, "key", "value")
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	err = pm.Delete(&ctx, "key")
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}
}

func TestDeleteDoesNotExist(t *testing.T) {
	cfg, err := config.New()
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	store := memory.New()
	crypto := crypto.StubEncrypter{}
	ctx := context.Background()

	pm, err := api.New(&ctx, cfg, store, &crypto)
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	ins := installer.StubInstaller{}

	err = pm.Setup(&ctx, &ins)
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	err = pm.Delete(&ctx, "key")
	if !errors.Is(err, common.ErrNoSuchKey) {
		t.Fatalf("wrong error type")
	}
}

func TestIsReady(t *testing.T) {
	cfg, err := config.New()
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	store := memory.New()
	crypto := crypto.StubEncrypter{}
	ctx := context.Background()

	pm, err := api.New(&ctx, cfg, store, &crypto)
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	if pm.IsReady() {
		t.Fatalf("unexpected ready")
	}

	ins := installer.StubInstaller{}

	err = pm.Setup(&ctx, &ins)
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	if !pm.IsReady() {
		t.Fatalf("expected ready")
	}
}
