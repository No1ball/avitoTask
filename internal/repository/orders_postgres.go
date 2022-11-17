package repository

import (
	"github.com/jmoiron/sqlx"
)

type OrdersDB struct {
	db *sqlx.DB
}

func NewOrdersDB(db *sqlx.DB) *OrdersDB {
	return &OrdersDB{db: db}
}
