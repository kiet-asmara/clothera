package handler

import (
	"context"
	"database/sql"
	"errors"
	"pair-project/entity"
	"time"
)

func UpdateClotheByID(db *sql.DB, clothe *entity.Clothes) error {
	query := `
		UPDATE Clothes
		SET 
			ClothesName = ?,
			ClothesCategory = ?,
			ClothesPrice = ?,
			ClothesStock = ?
		WHERE ClothesID = ?
	`

	args := []any{clothe.ClothesName, clothe.ClothesCategory, clothe.ClothesPrice, clothe.ClothesStock, clothe.ClothesID}

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
