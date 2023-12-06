package handler

import (
	"context"
	"database/sql"
	"fmt"
	"pair-project/entity"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func InsertCustomer(db *sql.DB, param *entity.Customer) error {
	query := `
		INSERT INTO customers(AddressID, CustomerName, CustomerEmail, CustomerPassword, CustomerType)
		VALUES (?, ?, ?, ?, ?)
	`
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(param.CustomerPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	args := []any{param.Address.AddressID, param.CustomerName, param.CustomerEmail, passwordHash, param.CustomerType}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "Duplicate entry"):
			return ErrorDuplicateEntry
		default:
			return err
		}
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error getting id from result: %v", err)
	}


	param.CustomerID = int(id)
	param.CustomerPassword = string(passwordHash)

	return nil
}
