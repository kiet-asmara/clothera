package handler

import (
	"context"
	"database/sql"
	"errors"
	"pair-project/entity"
	"time"
)

func GetAddressByID(db *sql.DB, inputID int) (*entity.Address, error) {
	query := `
	SELECT AddressID, AddressCountry, AddressCity, AddressStreet
	FROM Address
	WHERE AddressID = ?
	LIMIT 1
`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var address entity.Address
	err := db.QueryRowContext(ctx, query, inputID).Scan(
		&address.AddressID,
		&address.AddressCountry,
		&address.AddressCity,
		&address.AddressStreet,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrorRecordNotFound
		default:
			return nil, err
		}
	}

	return &address, nil

}
