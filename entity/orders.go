package entity

type Order struct {
	OrderID    int
	Customer   Customer
	OrderDate  string
	TotalPrice float64
}
