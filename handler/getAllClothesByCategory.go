package handler

import (
	"context"
	"database/sql"
	"pair-project/entity"
	"time"
)

func GetAllClothesByCategory(db *sql.DB, category string) ([]*entity.Clothes, error) {
	query := `
		SELECT 
			ClothesID,
			ClothesName,
			ClothesCategory,
			ClothesPrice,
			ClothesStock
		FROM
			clothes
		WHERE 
			ClothesCategory = ?
		
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, query, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clothes []*entity.Clothes
	for rows.Next() {
		var clothe = &entity.Clothes{}

		err := rows.Scan(
			&clothe.ClothesID,
			&clothe.ClothesName,
			&clothe.ClothesCategory,
			&clothe.ClothesPrice,
			&clothe.ClothesStock,
		)
		if err != nil {
			return nil, err
		}

		clothes = append(clothes, clothe)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return clothes, nil
}
