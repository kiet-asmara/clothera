package entity

type Order struct {
	OrderID    int
	Customer   Customer
	OrderDate  string
	TotalPrice float64
}

type RevenueByCostume struct {
	CostumeName  string
	Quantity     int
	TotalRevenue float32
}
