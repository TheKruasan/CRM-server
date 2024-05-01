package models

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type Order struct {
	Id          int       `json:"-"`
	Staff_id    int       `json:"staff_id"`
	Customer_id int       `json:"customer_id"`
	State_id    int       `json:"state_id"`
	Created_at  time.Time `json:"created_at"`
	Started_at  time.Time `json:"started_at"`
	Finish_at   time.Time `json:"finish_at"`
}

func GET_ORDERS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var orders = myOrders()
	json.NewEncoder(w).Encode(orders)
}

func myOrders() []Order {
	var orders = []Order{}
	db, err := sql.Open("postgres", "postres://postgres:1234@localhost/")

	if err != nil {
		logrus.Fatal(err)
	}
	rows, err := db.Query("SELECT * FROM order")
	if err != nil {
		logrus.Fatal(err)
	}
	for rows.Next() {
		var order Order = Order{}
		err := rows.Scan(&order.Id, &order.Customer_id, &order.State_id, &order.Created_at, &order.Started_at, &order.Finish_at)
		if err != nil {
			logrus.Fatal(err)
		}
		orders = append(orders, order)
	}

	return orders
}
