package handler

import (
	"context"
	"database/sql"
	"errors"
	"pair-project/entity"
	"time"
)

func UpdateCostumeByID(db *sql.DB, costume *entity.Costume) error {
	query := `
		UPDATE Costumes
		SET 
			CostumeName = ?,
			CostumeCategory = ?,
			CostumePrice = ?,
			CostumeStock = ?
		WHERE CostumeID = ?
	`

	args := []any{costume.CostumeName, costume.CostumeCategory, costume.CostumePrice, costume.CostumeStock, costume.CostumeID}

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
