package handler

import (
	"context"
	"database/sql"
	"errors"
	"pair-project/entity"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func UpdateCustomerByID(db *sql.DB, customer *entity.Customer, isUpdatePassword bool) error {
	var passwordhash = []byte(customer.CustomerPassword)

	var err error
	if isUpdatePassword {
		passwordhash, err = bcrypt.GenerateFromPassword([]byte(passwordhash), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
	}

	query := `
		UPDATE Customers
		SET 
			CustomerName = ?,
			CustomerEmail = ?,
			CustomerPassword = ?
		WHERE CustomerID = ?
	`

	args := []any{customer.CustomerName, customer.CustomerEmail, passwordhash, customer.CustomerID}

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

	customer.CustomerPassword = string(passwordhash)
	return nil
}
