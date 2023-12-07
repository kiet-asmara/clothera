package handler

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

func ShowProductsByCategoryCostumes(db *sql.DB) {
	// Implementasi untuk menampilkan semua produk berdasarkan category yang user klik
	// ...

	var categoryToDisplay string
	fmt.Print("Enter the category to display (type category name): ")
	fmt.Scan(&categoryToDisplay)

	// Gunakan db.Query untuk mendapatkan data dari database
	rows, err := db.Query("SELECT * FROM Costumes WHERE CostumeCategory = ?", categoryToDisplay)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Printf("Showing products in the '%s' category...\n", categoryToDisplay)
	fmt.Printf("%-10s %-30s %-15s %-15s %-15s\n", "ID", "Name", "Category", "Price", "Stock")
	fmt.Println(strings.Repeat("-", 80))

	for rows.Next() {
		var id int
		var name, category string
		var price float64
		var stock int

		err := rows.Scan(&id, &name, &category, &price, &stock)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%-10d %-30s %-15s %-15.2f %-15d\n", id, name, category, price, stock)
	}
}

func DeleteProductCostumes(db *sql.DB) {
	// Implementasi untuk menghapus produk berdasarkan input user
	// ...

	var idToDelete int
	fmt.Print("Enter the ID of the product to delete: ")
	fmt.Scan(&idToDelete)

	// Mengecek apakah ada penjualan terkait produk yang akan dihapus
	var rentsCount int
	err := db.QueryRow("SELECT COUNT(*) FROM Rents WHERE CostumeID = ?", idToDelete).Scan(&rentsCount)
	if err != nil {
		log.Fatal(err)
	}

	// Jika ada penjualan terkait, tampilkan pesan dan batalkan penghapusan
	if rentsCount > 0 {
		fmt.Println("Cannot delete the product. There are rents records associated with it.")
		return
	}

	// Gunakan db.Exec untuk menghapus data dari database jika tidak ada penjualan terkait
	result, err := db.Exec("DELETE FROM Costumes WHERE CostumeID = ?", idToDelete)

	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d product(s) deleted.\n", rowsAffected)
}
