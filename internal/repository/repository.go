package repository

import (
	"github.com/No1ball/avitoTask/internal/models"
	"github.com/jmoiron/sqlx"
)

type User interface {
	GetUserBalance(id int) (int, error)
	AddCashToUser(userId, cost int) error
	ReserveCash(userId, orderId, serviceId, cost int, serviceName, description string) error
	GetUser(id int) (models.User, error)
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
