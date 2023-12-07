package handler

import (
	"database/sql"
	"fmt"
	"pair-project/entity"
	"strings"
)

// features:
// revenue per user
// orders & quantity per user

func UserReportMenu(db *sql.DB) error {
	fmt.Println("\n1 -> Total Revenue per Customer")
	fmt.Println("2 -> Total Orders & Quantity per Customer")
	fmt.Println("3 -> Show Both")
	fmt.Println("4 -> Back to Main Menu")

	var choice int
	fmt.Println("Choice: ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		return fmt.Errorf("UserReportMenu: %w", err)
	}
	fmt.Println("")

	switch choice {
	case 1:
		err = RevenueCustomer(db)
		if err != nil {
			return fmt.Errorf("UserReportMenu: %w", err)
		}
	case 2:
		err = OrdersCustomer(db)
		if err != nil {
			return fmt.Errorf("UserReportMenu: %w", err)
		}
	case 3:
		err = RevenueCustomer(db)
		if err != nil {
			return fmt.Errorf("UserReportMenu: %w", err)
		}
		err = OrdersCustomer(db)
		if err != nil {
			return fmt.Errorf("UserReportMenu: %w", err)
		}
	case 4:
		return nil
	}
	return nil
}

func RevenueCustomer(db *sql.DB) error {
	query := `SELECT
	customers.CustomerID,
    customers.CustomerName,
    SUM(orders.totalPrice) AS revenue
	FROM customers
	JOIN orders ON customers.CustomerID = orders.CustomerID
    WHERE CustomerType = 'user'
	GROUP BY orders.CustomerID
	ORDER BY revenue DESC`

	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("RevenueCustomer: %w", err)
	}
	defer rows.Close()

	var customers []entity.CustomerRevenue

	for rows.Next() {
		var c entity.CustomerRevenue
		err := rows.Scan(&c.ID, &c.Name, &c.Revenue)
		if err != nil {
			return fmt.Errorf("RevenueCustomer: %w", err)
		}
		customers = append(customers, c)
	}

	fmt.Println("Showing Total Revenue by Customer...")
	fmt.Printf("\n%-15s %-20s %-15s\n", "CustomerID", "Customer Name", "Total Spent")
	fmt.Println(strings.Repeat("-", 53))

	for _, c := range customers {
		fmt.Printf("%-15d %-20s%-15.2f\n", c.ID, c.Name, c.Revenue)
	}
	fmt.Println("")

	return nil
}

func OrdersCustomer(db *sql.DB) error {
	query := `SELECT
	customers.CustomerID,
    customers.CustomerName,
    COUNT(orders.OrderID) AS OrderCount
	FROM orders
	JOIN customers ON customers.CustomerID = orders.CustomerID
    WHERE CustomerType = 'user'
	GROUP BY orders.CustomerID`

	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("OrdersCustomer: %w", err)
	}
	defer rows.Close()

	var customers []entity.CustomerOrders

	for rows.Next() {
		var c entity.CustomerOrders
		err := rows.Scan(&c.ID, &c.Name, &c.OrderCount)
		if err != nil {
			return fmt.Errorf("OrdersCustomer: %w", err)
		}
		customers = append(customers, c)
	}

	fmt.Println("Showing Total Orders by Customer...")
	fmt.Printf("\n%-15s %-20s %-15s\n", "CustomerID", "Customer Name", "Total Orders")
	fmt.Println(strings.Repeat("-", 53))

	for _, c := range customers {
		fmt.Printf("%-15d %-20s%-15d\n", c.ID, c.Name, c.OrderCount)
	}
	fmt.Println("")

	return nil
}
