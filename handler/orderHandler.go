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
