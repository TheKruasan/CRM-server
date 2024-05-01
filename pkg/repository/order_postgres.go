package repository

import (
	"fmt"

	models "github.com/TheKruasan/CRM-server/internal/repository/model"
	"github.com/jmoiron/sqlx"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{
		db: db,
	}
}

func (r *OrderPostgres) CreateOrder(order models.Order) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO order (staff_id, customer_id, state_id, created_at, started_at, finish_at)  values ($1 $2 $3 $4 $5 $6) RETURNING Id")
	row := r.db.QueryRow(query, order.Staff_id, order.Customer_id, order.State_id, order.Created_at, order.Started_at, order.Finish_at)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
