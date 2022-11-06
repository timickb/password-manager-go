package main

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/timickb/password-manager/internal/config"
	"github.com/timickb/password-manager/internal/crypto"
	"github.com/timickb/password-manager/internal/delivery/cli"
	"github.com/timickb/password-manager/internal/installer"
	"github.com/timickb/password-manager/internal/store/memory"
	"github.com/timickb/password-manager/pkg/api"
)

func main() {
	log := logrus.StandardLogger()
	ctx := context.Background()

	if err := mainNoExit(&ctx, log); err != nil {
		log.Fatalf("global error: %s", err.Error())
	}
}

func mainNoExit(ctx *context.Context, log *logrus.Logger) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	st := memory.New()
	cr := crypto.StubEncrypter{}

	pm, err := api.New(ctx, cfg, st, &cr)
	if err != nil {
		return err
	}

	ins := installer.StubInstaller{}

	c, err := cli.New(ctx, pm, &ins)
	if err != nil {
		return err
	}

	if len(os.Args) == 1 {
		c.Execute("help")
		return nil
	}

	c.Execute(os.Args[1], os.Args[2:]...)

	return nil
}
