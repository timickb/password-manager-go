package errors

type ErrNoSuchKey struct{}
type ErrStoreNotOpened struct{}
type ErrPassManagerNotReady struct{}

type ErrCmdNotFound struct{}
type ErrCmdWrongUsage struct{}

func (e ErrNoSuchKey) Error() string {
	return "no such key"
}

func (e ErrStoreNotOpened) Error() string {
	return "store wasn't opened"
}

func (e ErrPassManagerNotReady) Error() string {
	return "password manager not ready"
}

func (e ErrCmdNotFound) Error() string {
	return "command not found"
}

func (e ErrCmdWrongUsage) Error() string {
	return "wrong command usage"
}
