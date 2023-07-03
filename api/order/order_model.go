package order

import (
	"time"
)

type Order struct {
	Id          string    `json:"id"`
	ProductId   string    `json:"productId"`
	ProductName string    `json:"productName`
	ProductDesc string    `json:"productDesc`
	Status      string    `json:"status"`
	CreatedDate time.Time `json:"createdDate"`
}
