package handler

import (
	"context"
	"database/sql"
	"fmt"
	"pair-project/entity"
	"time"
)

func InsertAddress(db *sql.DB, param *entity.Address) error {
	query := `
		INSERT INTO address(AddressCountry, AddressCity, AddressStreet)
		VALUES (?, ?, ?)
	`

	args := []any{param.AddressCountry, param.AddressCity, param.AddressStreet}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error getting id from result: %v", err)
	}

	param.AddressID = int(id)
	return nil
}
