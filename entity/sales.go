package entity

type Sale struct {
	SaleID   int
	Order    Order
	Clothes  Clothes
	Quantity int
}

type ListSale struct {
	Name      string
	Quantity  int
	SalePrice float64
}
