package entity

type Rent struct {
	RentID    int
	OrderID   int
	CostumeID int
	Quantity  int
	StartDate string
	EndDate   string
}

type ListRent struct {
	Name      string
	Quantity  int
	RentPrice float64
}
