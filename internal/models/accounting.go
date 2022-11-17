package models

import "time"

type Accounting struct {
	id          int       `json:"id" db:"id"`
	name        string    `json:"name"`
	description string    `json:"description"`
	date        time.Time `json:"date"`
}
