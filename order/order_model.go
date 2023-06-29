package order

import(
	"time"
)

type Order struct {
	Id				string `json:"productId"`
	UserId    string `json:"user_id" validate:"required,max=50"`
	ProductId    string `json:"product_id" validate:"required,max=50"`
	Status		string `json:"status"`
	CreatedDate    time.Time  `json:"createdDate"`
}