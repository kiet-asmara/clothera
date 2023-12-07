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
	fmt.Println("1 -> Total Revenue per Customer")
	fmt.Println("2 -> Total Orders & Quantity per Customer")
	fmt.Println("3 -> Show Both")
	fmt.Println("4 -> Back to Main Menu")
	fmt.Println("")

	var choice int
	fmt.Print("Choice: ")
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
	Customers.CustomerID,
    Customers.CustomerName,
    SUM(Orders.TotalPrice) AS revenue
	FROM Customers
	JOIN Orders ON Customers.CustomerID = Orders.CustomerID
    WHERE CustomerType = 'user'
	GROUP BY Orders.CustomerID
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
	fmt.Printf("\n%-15s| %-20s| %-15s\n", "CustomerID", "Customer Name", "Total Spent")
	fmt.Println(strings.Repeat("-", 53))

	for _, c := range customers {
		fmt.Printf("%-15d| %-20s| %-15.2f\n", c.ID, c.Name, c.Revenue)
	}
	fmt.Println("")

	return nil
}

func OrdersCustomer(db *sql.DB) error {
	query := `SELECT
	Customers.CustomerID,
    Customers.CustomerName,
    COUNT(Orders.OrderID) AS OrderCount
	FROM Orders
	JOIN Customers ON Customers.CustomerID = Orders.CustomerID
    WHERE CustomerType = 'user'
	GROUP BY Orders.CustomerID`

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
	fmt.Printf("\n%-15s| %-20s| %-15s\n", "CustomerID", "Customer Name", "Total Orders")
	fmt.Println(strings.Repeat("-", 53))

	for _, c := range customers {
		fmt.Printf("%-15d| %-20s| %-15d\n", c.ID, c.Name, c.OrderCount)
	}
	fmt.Println("")

	return nil
}
