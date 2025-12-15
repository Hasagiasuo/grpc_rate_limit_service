package models

type Script struct {
	Rule string `redis:"rule"`
	Cost int    `redis:"cost"`
}
