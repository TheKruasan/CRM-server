package repository

import (
	models "github.com/TheKruasan/CRM-server/internal/repository/model"
	_ "github.com/jmoiron/sqlx"
)

type Order interface {
	CreateOrder(order models.Order) (int, error)
}

type Repository struct {
	Order
}

func NewRepository() *Repository {
	return &Repository{}
}

// func NewRepository(db *sqlx.DB) *Repository {
// 	return &Repository{
// 		Order: NewOrderPostgres(db),
// 	}
// }
