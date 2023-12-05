package entity

type Rent struct {
	RentID    int
	OrderID   Order
	Costume   Costume
	Quantity  int
	StartDate string
	EndDate   string
}
