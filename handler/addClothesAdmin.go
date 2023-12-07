package handler

import (
	"database/sql"
	"fmt"
	"log"
	"pair-project/entity"
)

func FetchAllCategoriesFromDatabase(db *sql.DB) []string {
	rows, err := db.Query("SELECT DISTINCT ClothesCategory FROM Clothes")
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

func GetSelectedCategoryFromUser(categories []string) string {
	var selectedCategory string
	fmt.Print("Select Category (Type category name): ")
	fmt.Scan(&selectedCategory)
	return selectedCategory
}

func GetProductDetailsFromAdmin(category string) entity.Clothes {
	var newProduct entity.Clothes
	newProduct.ClothesCategory = category

	fmt.Print("Enter Product Name: ")
	fmt.Scan(&newProduct.ClothesName)

	fmt.Print("Enter Product Price: ")
	fmt.Scan(&newProduct.ClothesPrice)

	fmt.Print("Enter Product Stock: ")
	fmt.Scan(&newProduct.ClothesStock)

	return newProduct
}

func InsertProductIntoDatabase(db *sql.DB, clothes entity.Clothes) error {
	_, err := db.Exec("INSERT INTO Clothes (ClothesName, ClothesCategory, ClothesPrice, ClothesStock) VALUES (?, ?, ?, ?)",
		clothes.ClothesName, clothes.ClothesCategory, clothes.ClothesPrice, clothes.ClothesStock)
	return err
}

func PrintCategoriesClothes(categories []string) {
	fmt.Println("List of Categories:")
	num := 1
	for _, category := range categories {
		fmt.Printf("%d. %s\n", num, category)
		num++
	}
	fmt.Println()
}
