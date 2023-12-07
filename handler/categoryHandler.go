package handler

import (
	"database/sql"
	"fmt"
	"pair-project/entity"
)

func CategoryCostume() string {
	fmt.Println("\n---Costume categories---")
	fmt.Println("1. Cosplay")
	fmt.Println("2. Formal")
	fmt.Println("Pilih kategori (1/2).")

	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil || (choice != 1 && choice != 2) {
		return "Invalid choice"
	}

	if choice == 1 {
		return "Cosplay"
	} else {
		return "Formal"
	}
}

func ListCostumes(db *sql.DB, category string) error {
	query := `
		SELECT * 
		FROM costumes 
		WHERE CostumeCategory = ?
	`

	rows, err := db.Query(query, category)
	if err != nil {
		return fmt.Errorf("ListCostumes: %w", err)
	}
	defer rows.Close()

	var costumes []entity.Costume

	for rows.Next() {
		var c entity.Costume
		err := rows.Scan(&c.CostumeID, &c.CostumeName, &c.CostumeCategory, &c.CostumePrice, &c.CostumeStock)
		if err != nil {
			return fmt.Errorf("ListCostumes: %w", err)
		}
		costumes = append(costumes, c)
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("ListCostumes: %w", err)
	}

	// list costumes
	fmt.Printf("\n---%s Costumes---\n", category)

	for _, c := range costumes {
		fmt.Printf("%d. %s, Price: %.2f, Stock: %d\n", c.CostumeID, c.CostumeName, c.CostumePrice, c.CostumeStock)
	}
	return nil
}
