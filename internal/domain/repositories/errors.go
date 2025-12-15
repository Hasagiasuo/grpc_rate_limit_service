package repositories

import "errors"

var (
	ErrUndefinedTimeType   = errors.New("time type is undefined")
	ErrUncorrectRuleFormat = errors.New("uncorrect rule format")
)
