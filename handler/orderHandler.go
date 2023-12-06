
package handler

import (
	"database/sql"
	"fmt"
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

	fmt.Println("Order created")

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

	fmt.Println("Updated price")

	return nil
}
