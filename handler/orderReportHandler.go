package handler

import (
	"database/sql"
	"fmt"
	"pair-project/entity"
)

func OrderReportMenu(db *sql.DB) error {
	fmt.Println("\n1 -> Total Revenue & Quantity Sold")
	fmt.Println("2 -> Rental Revenue by Costume")
	fmt.Println("3 -> Sales Revenue by Clothes")
	fmt.Println("4 -> Back to Main Menu")

	var choice int
	fmt.Println("Choice: ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		return fmt.Errorf("OrderReportMenu: %w", err)
	}

	switch choice {
	case 1:
		err = AllRevenue(db)
		if err != nil {
			return fmt.Errorf("OrderReportMenu: %w", err)
		}
		err = TotalQuantity(db)
		if err != nil {
			return fmt.Errorf("OrderReportMenu: %w", err)
		}
	case 2:
		err = RentalRevenueByCostume(db)
		if err != nil {
			return fmt.Errorf("OrderReportMenu: %w", err)
		}
	case 3:
		err = SalesRevenueByClothes(db)
		if err != nil {
			return fmt.Errorf("OrderReportMenu: %w", err)
		}
	case 4:
		return nil
	}
	return nil

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

func AllRevenue(db *sql.DB) error {
	// get values
	totalRev, err := TotalRevenue(db)
	if err != nil {
		return fmt.Errorf("AllRevenue: %w", err)
	}

	rentRev, err := TotalRentRevenue(db)
	if err != nil {
		return fmt.Errorf("AllRevenue: %w", err)
	}

	salesRev, err := TotalSalesRevenue(db)
	if err != nil {
		return fmt.Errorf("AllRevenue: %w", err)
	}

	// print
	fmt.Printf("\nTotal Revenue | Rental Revenue | Sales Revenue \n")
	fmt.Println("-----------------------------------------------")
	fmt.Printf("   %.2f    |     %.2f     |    %.2f \n\n", totalRev, rentRev, salesRev)

	return nil
}

func TotalRevenue(db *sql.DB) (float64, error) {
	query := `SELECT SUM(orders.totalPrice) AS totalRevenue FROM orders`

	rows, err := db.Query(query)
	if err != nil {
		return 0, fmt.Errorf("TotalRevenue: %w", err)
	}
	defer rows.Close()

	var totalRev float64

	for rows.Next() {
		err := rows.Scan(&totalRev)
		if err != nil {
			return 0, fmt.Errorf("TotalRevenue: %w", err)
		}
	}
	return totalRev, nil
}

func TotalRentRevenue(db *sql.DB) (float64, error) {
	query := `SELECT SUM(rents.rentPrice) AS totalRentRevenue FROM rents`

	rows, err := db.Query(query)
	if err != nil {
		return 0, fmt.Errorf("TotalRentRevenue: %w", err)
	}
	defer rows.Close()

	var totalRev float64

	for rows.Next() {
		err := rows.Scan(&totalRev)
		if err != nil {
			return 0, fmt.Errorf("TotalRentRevenue: %w", err)
		}
	}
	return totalRev, nil
}

func TotalSalesRevenue(db *sql.DB) (float64, error) {
	query := `SELECT SUM(sales.Quantity*clothes.ClothesPrice) AS totalSalesRevenue
				FROM sales JOIN clothes ON sales.ClothesID = clothes.ClothesID`

	rows, err := db.Query(query)
	if err != nil {
		return 0, fmt.Errorf("TotalSalesRevenue: %w", err)
	}
	defer rows.Close()

	var totalRev float64

	for rows.Next() {
		err := rows.Scan(&totalRev)
		if err != nil {
			return 0, fmt.Errorf("TotalSalesRevenue: %w", err)
		}
	}
	return totalRev, nil
}

func RentalRevenueByCostume(db *sql.DB) error {
	query := `SELECT 
	costumes.CostumeName, 
	SUM(rents.Quantity) AS Quantity, 
	SUM(rents.RentPrice) AS TotalRentPrice
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

func SalesRevenueByClothes(db *sql.DB) error {
	query := `SELECT 
    clothes.ClothesName, 
	(sales.Quantity) AS Quantity, 
	SUM(sales.Quantity*clothes.ClothesPrice) AS TotalSalesPrice
FROM sales
JOIN clothes ON sales.ClothesID = clothes.ClothesID
GROUP BY sales.ClothesID
ORDER BY TotalSalesPrice DESC`

	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("SalesRevenueByClothes: %w", err)
	}
	defer rows.Close()

	var sales []entity.RevenueByClothes

	for rows.Next() {
		var s entity.RevenueByClothes
		err := rows.Scan(&s.Name, &s.Quantity, &s.TotalRevenue)
		if err != nil {
			return fmt.Errorf("SalesRevenueByClothes: %w", err)
		}
		sales = append(sales, s)
	}

	fmt.Printf("\nClothes Name | Quantity | Total Revenue\n")
	fmt.Println("----------------------------------------")

	for _, s := range sales {
		fmt.Printf("%s   |   %d   |   %.2f\n", s.Name, s.Quantity, s.TotalRevenue)
	}
	fmt.Println("")

	return nil
}

// stock
// which items out of stock, which items low on stock < 5
