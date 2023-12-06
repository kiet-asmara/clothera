package handler

import (
	"context"
	"database/sql"
	"errors"
	"pair-project/entity"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func UpdateCustomerByID(db *sql.DB, customer *entity.Customer) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(customer.CustomerPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `
		UPDATE Customers
		SET 
			CustomerName = ?,
			CustomerEmail = ?,
			CustomerPassword = ?
		WHERE CustomerID = ?
	`

	args := []any{customer.CustomerName, customer.CustomerEmail, passwordHash, customer.CustomerID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrorRecordNotFound
		default:
			return err
		}
	}

	customer.CustomerPassword = string(passwordHash)
	return nil
}
