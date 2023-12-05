package entity

type Sale struct {
	SaleID   int
	Order    Order
	Clothes  Clothes
	Quantity int
}
