package service

import (
	models "github.com/TheKruasan/CRM-server/internal/repository/model"
	"github.com/TheKruasan/CRM-server/pkg/repository"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (s *OrderService) CreateOrder(order models.Order) (int, error) {
	return s.repo.CreateOrder(order)
}
