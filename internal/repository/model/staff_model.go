package models

import "time"

type Staff struct {
	Id          int       `json:"-"`
	Role_id     int       `json:"role_id"`
	First_name  string    `json:"first_name"`
	Second_name string    `json:"second_name"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	Address     string    `json:"address"`
	Created_at  time.Time `json:"created_at"`
	Started_at  time.Time `json:"started_at"`
}

//binding: "required"
