package handler

import (
	"database/sql"
	"fmt"
)

var Categories []string

func ListCategory(db *sql.DB) {
	query := "SELECT DISTINCT ClothesCategory FROM Clothes"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error querying the database:", err)
		return
	}
	defer rows.Close()

	var category string
	Categories = nil

	fmt.Println("\n---Clothing Categories---")
	num := 1
	for rows.Next() {
		err := rows.Scan(&category)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		fmt.Printf("%d. %s\n", num, category)
		num++
		Categories = append(Categories, category)
	}

	fmt.Println()

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
		return
	}
}

func DisplayClothesByCategory(db *sql.DB, selectedCategory string) []string {
	query := "SELECT ClothesName FROM Clothes WHERE ClothesCategory = ?"
	rows, err := db.Query(query, selectedCategory)
	if err != nil {
		fmt.Println("Error querying the database:", err)
		return nil
	}
	defer rows.Close()

	var result []string

	fmt.Printf("\n---Clothes in Category '%s'---\n", selectedCategory)
	num := 1
	for rows.Next() {
		var clothesName string
		err := rows.Scan(&clothesName)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil
		}
		fmt.Printf("%d. %s\n", num, clothesName)
		num++
		result = append(result, clothesName)
	}

	fmt.Println()

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
		return nil
	}

	return result
}
