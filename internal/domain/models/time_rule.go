package models

import "time"

type TimeRule struct {
	Key          string
	ServiceToken string
	Limit        int64
	Window       time.Duration
	Cost         int64
}
