package service

import (
	models "github.com/TheKruasan/CRM-server/internal/repository/model"
	"github.com/TheKruasan/CRM-server/pkg/repository"
)

type Order interface {
	CreateOrder(order models.Order) (int, error)
}

type Service struct {
	Order
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Order: NewOrderService(repository.Order),
	}
}
