package services

import "github.com/No1ball/avitoTask/internal/repository"

type AccountingService struct {
	repo repository.Accounting
}

func NewAccountingService(repo repository.Accounting) *AccountingService {
	return &AccountingService{repo: repo}
}

func (s *AccountingService) RevenueConfirmation(userId, orderId, serviceId, cost int) error {
	return s.repo.RevenueConfirmation(userId, orderId, serviceId, cost)
}
