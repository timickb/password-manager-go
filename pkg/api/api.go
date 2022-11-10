package api

import (
	"context"

	"github.com/timickb/password-manager/internal/common"
	"github.com/timickb/password-manager/internal/config"
	"github.com/timickb/password-manager/internal/crypto"
	"github.com/timickb/password-manager/internal/installer"
	"github.com/timickb/password-manager/internal/store"
)

type PasswordManager struct {
	cfg   *config.Config
	st    store.Store
	cr    crypto.Encrypter
	ready bool
}

func New(ctx *context.Context, cfg *config.Config, st store.Store, cr crypto.Encrypter) (*PasswordManager, error) {
	return &PasswordManager{cfg: cfg, st: st, cr: cr, ready: false}, nil
}

func (p *PasswordManager) IsReady() bool {
	return p.ready
}

// Set a new pair "name-secret" to the store.
func (p *PasswordManager) Set(ctx *context.Context, name string, secret string) error {
	if !p.ready {
		return common.ErrPassManagerNotReady
	}
	if err := p.st.SetItem(ctx, name, secret); err != nil {
		return err
	}
	return nil
}

// Read the secret by its name.
func (p *PasswordManager) Read(ctx *context.Context, name string) (string, error) {
	if !p.ready {
		return "", common.ErrPassManagerNotReady

	}

	sec, err := p.st.GetItem(ctx, name)
	if err != nil {
		return "", err
	}

	return sec, nil
}

// Delete the secret by its name.
func (p *PasswordManager) Delete(ctx *context.Context, name string) error {
	if !p.ready {
		return common.ErrPassManagerNotReady
	}

	if err := p.st.RemoveItem(ctx, name); err != nil {
		return err
	}
	return nil
}

// Setup the password manager.
func (p *PasswordManager) Setup(ctx *context.Context, ins installer.Installer) error {
	if err := ins.Install(); err != nil {
		return err
	}

	if err := p.st.Open(ctx); err != nil {
		return err
	}

	p.ready = true

	return nil
}
