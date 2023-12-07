package handler

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

func ShowProductsByCategory(db *sql.DB) {
	for {
		var categoryToDisplay string
		fmt.Print("Enter the category to display (type category name): ")
		fmt.Scan(&categoryToDisplay)

		// Gunakan db.Query untuk mendapatkan data dari database
		rows, err := db.Query("SELECT * FROM Clothes WHERE ClothesCategory = ?", categoryToDisplay)

		if err != nil {
			log.Fatal(err)
		}

		// Periksa apakah ada hasil yang ditemukan
		if !rows.Next() {
			fmt.Printf("No products found in the '%s' category.\n", categoryToDisplay)
			continue // Lanjut ke iterasi berikutnya jika tidak ada produk yang ditemukan
		}

		// Tampilkan produk jika ada hasil yang ditemukan
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

		// Tanya user apakah ingin melihat kategori lainnya
		fmt.Print("Do you want to view products in another category? (y/n): ")
		var input string
		fmt.Scan(&input)
		if strings.ToLower(input) != "y" {
			break // Keluar dari loop jika user tidak ingin melihat kategori lainnya
		}
	}
}

func DeleteProduct(db *sql.DB) {
	// Implementasi untuk menghapus produk berdasarkan input user
	// ...

	var idToDelete int
	fmt.Print("Enter the ID of the product to delete: ")
	fmt.Scan(&idToDelete)

	// Mengecek apakah ada penjualan terkait produk yang akan dihapus
	var salesCount int
	err := db.QueryRow("SELECT COUNT(*) FROM Sales WHERE ClothesID = ?", idToDelete).Scan(&salesCount)
	if err != nil {
		log.Fatal(err)
	}

	// Jika ada penjualan terkait, tampilkan pesan dan batalkan penghapusan
	if salesCount > 0 {
		fmt.Println("Cannot delete the product. There are sales records associated with it.")
		return
	}

	// Gunakan db.Exec untuk menghapus data dari database jika tidak ada penjualan terkait
	result, err := db.Exec("DELETE FROM Clothes WHERE ClothesID = ?", idToDelete)

	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d product(s) deleted.\n", rowsAffected)
}
