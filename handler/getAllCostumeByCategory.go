package handler

import (
	"context"
	"database/sql"
	"pair-project/entity"
	"time"
)

func GetAllCostumeByCategory(db *sql.DB, category string) ([]*entity.Costume, error) {
	query := `
		SELECT 
			CostumeID,
			CostumeName,
			CostumeCategory,
			CostumePrice,
			CostumeStock
		FROM
			costumes
		WHERE 
			CostumeCategory = ?
		
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, query, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var costumes []*entity.Costume
	for rows.Next() {
		var costume = &entity.Costume{}

		err := rows.Scan(
			&costume.CostumeID,
			&costume.CostumeName,
			&costume.CostumeCategory,
			&costume.CostumePrice,
			&costume.CostumeStock,
		)
		if err != nil {
			return nil, err
		}

		costumes = append(costumes, costume)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return costumes, nil
}
