package handler

import (
	"context"
	"database/sql"
	"errors"
	"pair-project/entity"
	"time"
)

func UpdateAdressByID(db *sql.DB, address *entity.Address) error {
	query := `
		UPDATE address
		SET 
			AddressCountry = ?,
			AddressCity = ?,
			AddressStreet = ?
		WHERE 
			addressID = ?
	`

	args := []any{address.AddressCountry, address.AddressCity, address.AddressStreet, address.AddressID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrorRecordNotFound
		default:
			return err
		}
	}
	return nil
}
