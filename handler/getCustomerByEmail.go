package handler

import (
	"context"
	"database/sql"
	"errors"
	"pair-project/entity"
	"time"
)

func GetCustomerByEmail(db *sql.DB, email string) (*entity.Customer, error) {
	query := `
		SELECT CustomerID, AddressID, CustomerName, CustomerEmail, CustomerPassword, CustomerType
		FROM Customers
		WHERE CustomerEmail = ?
		LIMIT 1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var customer entity.Customer
	var passwordHash []byte

	err := db.QueryRowContext(ctx, query, email).Scan(
		&customer.CustomerID,
		&customer.Address.AddressID,
		&customer.CustomerName,
		&customer.CustomerEmail,
		&passwordHash,
		&customer.CustomerType,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrorRecordNotFound
		default:
			return nil, err
		}
	}

	customer.CustomerPassword = string(passwordHash)
	return &customer, nil
}
