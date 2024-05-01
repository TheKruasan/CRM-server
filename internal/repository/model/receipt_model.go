package models

import "time"

type Receipt struct {
	Id          int       `json:"-"`
	Staff_id    int       `json:"staff_id"`
	Customer_id int       `json:"customer_id"`
	State_id    int       `json:"state_id"`
	Created_at  time.Time `json:"created_at"`
	Started_at  time.Time `json:"started_at"`
	Finish_at   time.Time `json:"finish_at"`
}
