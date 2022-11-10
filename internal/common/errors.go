package common

import "errors"

var ErrNoSuchKey = errors.New("no such key")
var ErrStoreNotOpened = errors.New("store wasn't opened")
var ErrPassManagerNotReady = errors.New("password manager not ready")

var ErrCmdNotFound = errors.New("command not found")
var ErrCmdWrongUsage = errors.New("wrong commands usage")
