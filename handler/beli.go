package handler

import (
	"database/sql"
	"fmt"
	"pair-project/entity"
)

func AddProduct(db *sql.DB, clothes entity.Clothes) error {
	query := `INSERT INTO products (ClothesName, ClothesCategory, ClothesPrice, ClothesStock) VALUES (?, ?, ?, ?)`

	_, err := db.Exec(query, clothes.ClothesName, clothes.ClothesCategory, clothes.ClothesPrice, clothes.ClothesPrice)
	if err != nil {
		return err
	}

	fmt.Println("Product added successfully")
	return nil

}
