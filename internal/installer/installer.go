package installer

type Installer interface {
	Install() error
}

type StubInstaller struct{}

func (i *StubInstaller) Install() error {
	return nil
}
