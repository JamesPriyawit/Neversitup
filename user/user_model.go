package user

import (
	"time"
)

type User struct {
	Id			string `json:"id"`
	Username    string `json:"username" validate:"required,min=6,max=50"`
	Password    string `json:"password" validate:"required,min=6,max=50"`
	RePassword  string `json:"rePassword" validate:"required,min=6,max=50"`
	Firstname   string `json:"firstname" validate:"required,max=50"`
	Lastname    string `json:"lastname" validate:"required,max=50"`
	CreatedDate    time.Time  `json:"createdDate"`
}