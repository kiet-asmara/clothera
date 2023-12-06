// handler.go

package handler

import (
	"database/sql"
	"fmt"
)

var Categories []string // Global variable to store categories

func ListCategory(db *sql.DB) {
	query := "SELECT DISTINCT ClothesCategory FROM Clothes"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error querying the database:", err)
		return
	}
	defer rows.Close()

	var category string
	Categories = nil // Reset Categories slice before populating it again

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

func DisplayClothesByCategory(db *sql.DB, selectedCategory string) {
	query := "SELECT ClothesName FROM Clothes WHERE ClothesCategory = ?"
	rows, err := db.Query(query, selectedCategory)
	if err != nil {
		fmt.Println("Error querying the database:", err)
		return
	}
	defer rows.Close()

	var clothesName string

	fmt.Printf("\n---Clothes in Category '%s'---\n", selectedCategory)
	num := 1
	for rows.Next() {
		err := rows.Scan(&clothesName)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		fmt.Printf("%d. %s\n", num, clothesName)
		num++
	}

	fmt.Println()

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
		return
	}
}
