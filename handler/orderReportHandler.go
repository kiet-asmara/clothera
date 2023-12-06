package handler

import (
	"database/sql"
	"fmt"
	"pair-project/entity"
)

func OrderReportMenu() {
	fmt.Println("\n1 -> Total Quantity Sold")
	fmt.Println("2 -> Total Revenue")
	fmt.Println("3 -> Back to Main Menu")
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

func RentalRevenueByCostume(db *sql.DB) error {
	query := `SELECT 
	costumes.CostumeName, 
	SUM(rents.Quantity) AS Quantity, 
	(rents.Quantity*costumes.CostumePrice) AS TotalRentPrice
FROM rents
JOIN costumes ON rents.CostumeID = costumes.CostumeID
JOIN orders ON rents.OrderID = orders.OrderID
GROUP BY costumes.CostumeName`

	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("RentalRevenueByCostume: %w", err)
	}
	defer rows.Close()

	var rentals []entity.RevenueByCostume

	for rows.Next() {
		var r entity.RevenueByCostume
		err := rows.Scan(&r.CostumeName, &r.Quantity, &r.TotalRevenue)
		if err != nil {
			return fmt.Errorf("RentalRevenueByCostume: %w", err)
		}
		rentals = append(rentals, r)
	}

	fmt.Printf("\nCostume Name | Quantity | Total Revenue\n")
	fmt.Println("----------------------------------------")

	for _, r := range rentals {
		fmt.Printf("%s   |   %d   |   %.2f\n", r.CostumeName, r.Quantity, r.TotalRevenue)
	}
	fmt.Println("")

	return nil
}

// order
// per year -> total products sold, orders made, how many rentals & sells
// per month

// stock
// which items out of stock, which items low on stock < 5
// choose clothes, rentals -> clothes category ->
