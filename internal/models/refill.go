package models

import "time"

type Refill struct {
	id     int       `json:"id" db:"id"`
	cost   int       `json:"cost"`
	userId int       `json:"userId"`
	date   time.Time `json:"date"`
}
