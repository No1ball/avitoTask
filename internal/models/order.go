package models

type Order struct {
	id   int `json:"id" db:"id"`
	cost int `json:"cost"`
}
