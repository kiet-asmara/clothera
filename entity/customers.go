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
