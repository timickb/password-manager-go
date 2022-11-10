package cli_test

import (
	"context"
	"errors"
	"testing"

	"github.com/timickb/password-manager/internal/common"
	"github.com/timickb/password-manager/internal/config"
	"github.com/timickb/password-manager/internal/crypto"
	"github.com/timickb/password-manager/internal/delivery/cli"
	"github.com/timickb/password-manager/internal/installer"
	"github.com/timickb/password-manager/internal/store/memory"
	"github.com/timickb/password-manager/pkg/api"
)

func TestNew(t *testing.T) {
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	store := memory.New()
	crypto := crypto.StubEncrypter{}

	pm, err := api.New(&ctx, cfg, store, &crypto)
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	ins := installer.StubInstaller{}

	if _, err := cli.New(&ctx, pm, &ins); err != nil {
		t.Fatalf("error: %s", err.Error())
	}
}

func TestExecuteWrongCmd(t *testing.T) {
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	store := memory.New()
	crypto := crypto.StubEncrypter{}

	pm, err := api.New(&ctx, cfg, store, &crypto)
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	ins := installer.StubInstaller{}

	cli, err := cli.New(&ctx, pm, &ins)
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	err = cli.Execute("unexpected_command")
	if !errors.Is(err, common.ErrCmdNotFound) {
		t.Fatalf("wrong error type")
	}
}

func TestExecute(t *testing.T) {
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	store := memory.New()
	crypto := crypto.StubEncrypter{}

	pm, err := api.New(&ctx, cfg, store, &crypto)
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	ins := installer.StubInstaller{}

	cli, err := cli.New(&ctx, pm, &ins)
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	// help right usage case
	if err = cli.Execute("help"); err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	// setup right usage case
	if err = cli.Execute("setup"); err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	// get right usage case
	if err = cli.Execute("get", "secret_name"); err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	// get wrong usage case
	err = cli.Execute("get")
	if !errors.Is(err, common.ErrCmdWrongUsage) {
		t.Fatalf("wrong error type")
	}

	// set right usage case
	if err = cli.Execute("set", "secret_name", "secret_value"); err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	// set wrong usage case 1
	err = cli.Execute("set", "secret_name")
	if !errors.Is(err, common.ErrCmdWrongUsage) {
		t.Fatalf("wrong error type")
	}

	// set wrong usage case 2
	err = cli.Execute("set")
	if !errors.Is(err, common.ErrCmdWrongUsage) {
		t.Fatalf("wrong error type")
	}

	// delete right usage case
	if err = cli.Execute("delete", "secret_name"); err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	// delete wrong usage case 1
	err = cli.Execute("delete")
	if !errors.Is(err, common.ErrCmdWrongUsage) {
		t.Fatalf("wrong error type")
	}
}
