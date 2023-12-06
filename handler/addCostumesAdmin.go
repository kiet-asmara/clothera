package handler

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"pair-project/entity"
// )

// func FetchAllCategoriesFromDatabaseCostumes(db *sql.DB) []string {
// 	rows, err := db.Query("SELECT DISTINCT CostumeCategory FROM Costumes")
// 	defer rows.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var categories []string
// 	for rows.Next() {
// 		var category string
// 		rows.Scan(&category)
// 		categories = append(categories, category)
// 	}
// 	return categories
// }

// func GetSelectedCategoryFromUserCostumes(categories []string) string {
// 	var selectedCategory string
// 	fmt.Print("Select Category: ")
// 	fmt.Scan(&selectedCategory)
// 	return selectedCategory
// }

// func GetProductDetailsFromAdminCostumes(category string) entity.Costume {
// 	var newProduct entity.Costume
// 	newProduct.ClothesCategory = category

// 	fmt.Print("Enter Product Name: ")
// 	fmt.Scan(&newProduct.ClothesName)

// 	fmt.Print("Enter Product Price: ")
// 	fmt.Scan(&newProduct.ClothesPrice)

// 	fmt.Print("Enter Product Stock: ")
// 	fmt.Scan(&newProduct.ClothesStock)

// 	return newProduct
// }

// func InsertProductIntoDatabaseCostumes(db *sql.DB, clothes entity.Clothes) error {
// 	_, err := db.Exec("INSERT INTO Clothes (ClothesName, ClothesCategory, ClothesPrice, ClothesStock) VALUES (?, ?, ?, ?)",
// 		clothes.ClothesName, clothes.ClothesCategory, clothes.ClothesPrice, clothes.ClothesStock)
// 	return err
// }
