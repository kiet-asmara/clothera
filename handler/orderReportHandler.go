package handler

import (
	"database/sql"
	"fmt"
)

func OrderReportMenu() {
	fmt.Println("\n1 -> Total Quantity Sold")
	fmt.Println("2 -> Total Revenue")
	fmt.Println("3 -> Back to Main Menu\n")
}

func TotalQuantity(db *sql.DB) error {
	query := `SELECT
	SUM(rents.Quantity) AS rentQuantity,
    SUM(sales.Quantity) AS saleQuantity
FROM orders
JOIN rents ON orders.OrderID = rents.OrderID
JOIN sales ON orders.OrderID = sales.OrderID`

	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("TotalOrders: %w", err)
	}
	defer rows.Close()

	var rentQuantity, saleQuantity int

	for rows.Next() {
		err := rows.Scan(&rentQuantity, &saleQuantity)
		if err != nil {
			return fmt.Errorf("RentPrice: %w", err)
		}
	}

	fmt.Printf("\nTotal Products Sold | Total Rentals | Total Sales \n")
	fmt.Println("--------------------------------------------------")
	fmt.Printf("         %d         |       %d      |      %d \n\n", (rentQuantity + saleQuantity), rentQuantity, saleQuantity)

	return nil
}

// order
// per year -> total products sold, orders made, how many rentals & sells
// per month

// stock
// which items out of stock, which items low on stock < 5
// choose clothes, rentals -> clothes category ->
