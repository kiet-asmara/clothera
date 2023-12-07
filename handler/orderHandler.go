package handler

import (
	"database/sql"
	"fmt"
	"pair-project/entity"
	"time"
)

func CreateOrder(db *sql.DB, customerID int) (int, error) {
	// create order
	query := `INSERT INTO orders (CustomerID, orderDate) VALUES
	(?,?)`

	orderDate := time.Now().Format("2006-01-02")

	_, err := db.Exec(query, customerID, orderDate)
	if err != nil {
		return 0, fmt.Errorf("CreateOrder: %w", err)
	}

	// get order ID
	query = `SELECT orderID FROM orders ORDER BY orderID DESC LIMIT 1`
	rows, err := db.Query(query)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var orderID int

	for rows.Next() {
		err := rows.Scan(&orderID)
		if err != nil {
			return 0, err
		}
	}

	return orderID, nil
}

func CalcDiscount(totalPrice float64) float64 {

	// calculate discount
	if totalPrice > 500 {
		// diskon 20%
		fmt.Println("Congrats! You qualify for a 20% discount.")
		totalPrice *= 0.8
	} else if totalPrice > 200 {
		// diskon 10%
		fmt.Println("Congrats! You qualify for a 10% discount.")
		totalPrice *= 0.9
	}

	// calculate taxes (PPn = 11%)
	totalPrice *= 1.11

	return totalPrice
}

func InsertTotal(db *sql.DB, total float64, orderID int) error {
	query := `UPDATE orders
	SET TotalPrice = ?
	WHERE OrderID = ?`

	_, err := db.Exec(query, total, orderID)
	if err != nil {
		return fmt.Errorf("calculateTotal: %w", err)
	}

	return nil
}

func ListPesanan(db *sql.DB, orderID int) error {
	err := ListRental(db, orderID)
	if err != nil {
		return fmt.Errorf("ListPesanan: %w", err)
	}

	err = ListSales(db, orderID)
	if err != nil {
		return fmt.Errorf("ListPesanan: %w", err)
	}
	return nil
}

func ListRental(db *sql.DB, orderID int) error {
	query := `SELECT
	costumes.CostumeName,
    rents.Quantity,
    rents.RentPrice
FROM rents
JOIN costumes ON rents.CostumeID = costumes.CostumeID
WHERE rents.OrderID = ?
GROUP BY rents.CostumeID`

	rows, err := db.Query(query, orderID)
	if err != nil {
		return fmt.Errorf("ListRental: %w", err)
	}
	defer rows.Close()

	var rentList []entity.ListRent

	for rows.Next() {
		var r entity.ListRent
		err := rows.Scan(&r.Name, &r.Quantity, &r.RentPrice)
		if err != nil {
			return fmt.Errorf("ListRental: %w", err)
		}
		rentList = append(rentList, r)
	}

	for _, r := range rentList {
		fmt.Printf("%s, Quantity: %d, Price: %.2f\n", r.Name, r.Quantity, r.RentPrice)
	}

	return nil
}

func ListSales(db *sql.DB, orderID int) error {
	query := `SELECT
	clothes.ClothesName,
    (sales.Quantity) AS Quantity,
    (clothes.ClothesPrice * sales.Quantity) AS TotalPrice
	FROM sales
	JOIN clothes ON sales.ClothesID = clothes.ClothesID
	WHERE sales.OrderID = ?
    GROUP BY clothes.ClothesName`

	rows, err := db.Query(query, orderID)
	if err != nil {
		return fmt.Errorf("ListSales: %w", err)
	}
	defer rows.Close()

	var saleList []entity.ListSale

	for rows.Next() {
		var s entity.ListSale
		err := rows.Scan(&s.Name, &s.Quantity, &s.SalePrice)
		if err != nil {
			return fmt.Errorf("ListSales: %w", err)
		}
		saleList = append(saleList, s)
	}

	for _, s := range saleList {
		fmt.Printf("%s, Quantity: %d, Price: %.2f\n", s.Name, s.Quantity, s.SalePrice)
	}

	return nil
}
