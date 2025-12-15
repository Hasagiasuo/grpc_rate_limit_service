package models

type Service struct {
	Token string `db:"token"`
	Balance int `db:"balance"`
}