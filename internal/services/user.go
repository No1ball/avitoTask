package services

import (
	"github.com/No1ball/avitoTask/internal/repository"
)

type UserService struct {
	userRepo repository.User
}

func NewUserService(userRepo repository.User) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUserBalance(id int) (int, error) {
	return s.userRepo.GetUserBalance(id)
}

func (s *UserService) AddCashToUser(userId, cost int) error {
	if _, err := s.userRepo.CheckUser(userId); err != nil {
		return s.userRepo.AddCashToUser(userId, cost)
	}
	return s.userRepo.AddCashToUserWithUpdate(userId, cost)
}

func (s *UserService) ReserveCash(userId, orderId, serviceId, cost int, serviceName, description string) error {
	return s.userRepo.ReserveCash(userId, orderId, serviceId, cost, serviceName, description)
}
