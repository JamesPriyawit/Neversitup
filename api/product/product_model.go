package product

import(
	"time"
)

type Product struct {
	Id				string `json:"productId"`
	ProductName    string `json:"productName" validate:"required,max=255"`
	ProductDesc    string `json:"productDesc" validate:"max=1000"`
	ProductType		string `json:"productType" validate:"required,max=50"`
	Price  			int `json:"price" validate:"numeric"`
	CreatedDate    time.Time  `json:"createdDate"`
}