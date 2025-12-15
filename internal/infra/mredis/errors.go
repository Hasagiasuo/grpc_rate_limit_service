package mredis

import "errors"

var (
	ErrCannotInitRedis = errors.New("cannot init redis client")
)
