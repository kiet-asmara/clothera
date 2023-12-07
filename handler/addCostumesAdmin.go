package handler

import (
	"database/sql"
	"fmt"
	"log"
	"pair-project/entity"
)

func FetchAllCategoriesFromDatabaseCostumes(db *sql.DB) []string {
	rows, err := db.Query("SELECT DISTINCT CostumeCategory FROM Costumes")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	var categories []string
	for rows.Next() {
		var category string
		rows.Scan(&category)
		categories = append(categories, category)
	}
	return categories
}

func GetSelectedCategoryFromUserCostumes(categories []string) string {
	var selectedCategory string
	fmt.Print("Select Category (Type category name): ")
	fmt.Scan(&selectedCategory)
	return selectedCategory
}

func GetProductDetailsFromAdminCostumes(category string) entity.Costume {
	var newProduct entity.Costume
	newProduct.CostumeCategory = category

	fmt.Print("Enter Product Name: ")
	fmt.Scan(&newProduct.CostumeName)

	fmt.Print("Enter Product Price: ")
	fmt.Scan(&newProduct.CostumePrice)

	fmt.Print("Enter Product Stock: ")
	fmt.Scan(&newProduct.CostumeStock)

	return newProduct
}

func InsertProductIntoDatabaseCostumes(db *sql.DB, costume entity.Costume) error {
	_, err := db.Exec("INSERT INTO Costumes (CostumeName, CostumeCategory, CostumePrice, CostumeStock) VALUES (?, ?, ?, ?)",

		costume.CostumeName, costume.CostumeCategory, costume.CostumePrice, costume.CostumeStock)
	return err
}

func PrintCategoriesCostumes(categories []string) {
	fmt.Println("List of Categories:")
	num := 1
	for _, category := range categories {
		fmt.Printf("%d. %s\n", num, category)
		num++
	}
	fmt.Println()
}
