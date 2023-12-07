package handler

import (
	"database/sql"
	"fmt"
	"pair-project/entity"
)

func GetCostumeByID(db *sql.DB, costumeID int) (*entity.Costume, error) {
	query := `
		SELECT CostumeID, CostumeName, CostumeCategory, CostumePrice, CostumeStock
		FROM costumes
		WHERE CostumeID = ?
	`

	var costume entity.Costume
	err := db.QueryRow(query, costumeID).Scan(
		&costume.CostumeID,
		&costume.CostumeName,
		&costume.CostumeCategory,
		&costume.CostumePrice,
		&costume.CostumeStock,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get costume by ID: %v", err)
	}

	return &costume, nil
}
