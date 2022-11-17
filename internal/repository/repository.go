package repository

import (
	"github.com/jmoiron/sqlx"
)

type User interface {
	GetUserBalance(id int) (int, error)
	AddCashToUser(userId, cost int) error
	ReserveCash(userId, orderId, serviceId, cost int, serviceName, description string) error
	CheckUser(id int) (int, error)
	AddCashToUserWithUpdate(userId, cost int) error
}

type Accounting interface {
	RevenueConfirmation(userId, orderId, serviceId, cost int) error
}

type Repository struct {
	User
	Accounting
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		NewUserBD(db),
		NewAccountingDB(db),
	}
}
