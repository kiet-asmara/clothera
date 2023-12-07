package entity

const (
	User  string = "user"
	Admin string = "admin"
)

type Customer struct {
	CustomerID       int
	Address          Address
	CustomerName     string
	CustomerEmail    string
	CustomerPassword string
	CustomerType     string
}

type CustomerRevenue struct {
	ID      int
	Name    string
	Revenue float64
}

type CustomerOrders struct {
	ID         int
	Name       string
	OrderCount int
}
