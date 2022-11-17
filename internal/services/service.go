package services

import (
	"github.com/No1ball/avitoTask/internal/repository"
)

type User interface {
	GetUserBalance(id int) (int, error)
	AddCashToUser(userId, cost int) error
	ReserveCash(userId, orderId, serviceId, cost int, serviceName, description string) error
}

type Accounting interface {
	RevenueConfirmation(userId, orderId, serviceId, cost int) error
}
type Service struct {
	User
	Accounting
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		NewUserService(repo.User),
		NewAccountingService(repo.Accounting),
	}
}
