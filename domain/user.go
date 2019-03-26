package domain

import "time"

type Users []User

type User struct {
	//gorm.Model
	ID               int    `json:"id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	PhoneNumber      string `json:"phone_num"`
	ActivationStatus uint   `json:"activation_status"`
	UserAddress      string `json:"user_address"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time `sql:"index"`
}
