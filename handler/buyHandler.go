package handler

import (
	"database/sql"
	"pair-project/entity"
)

func AddProduct(db *sql.DB, clothes entity.Clothes, customer entity.Customer, order entity.Order, sale entity.Sale) error {

	// Add the customer to the customers table
	customerQuery := `
		INSERT INTO customers (AddressID, CustomerName, CustomerEmail, CustomerPassword, CustomerType)
		VALUES (?, ?, ?, ?, ?)
	`

	result, err := db.Exec(
		customerQuery,
		customer.Address.AddressID,
		customer.CustomerName,
		customer.CustomerEmail,
		customer.CustomerPassword,
		customer.CustomerType,
	)

	if err != nil {
		return err
	}


	// Get the last inserted customer ID
	customerID, err := result.LastInsertId()
	if err != nil {
		return err
	}


	// Add the order to the orders table
	orderQuery := `
		INSERT INTO orders (CustomerID, OrderDate, TotalPrice)
		VALUES (?, ?, ?)
	`

	_, err = db.Exec(orderQuery, customerID, order.OrderDate, order.TotalPrice)
	if err != nil {
		return err
	}


	// Get the last inserted order ID
	orderID, err := result.LastInsertId()
	if err != nil {
		return err
	}


	// Add the clothes to the clothes table
	clothesQuery := `
		INSERT INTO clothes (ClothesName, ClothesCategory, ClothesPrice, ClothesStock)
		VALUES (?, ?, ?, ?)
	`

	result, err = db.Exec(
		clothesQuery,
		clothes.ClothesName,
		clothes.ClothesCategory,
		clothes.ClothesPrice,
		clothes.ClothesStock,
	)

	if err != nil {
		return err
	}


	// Get the last inserted clothes ID
	clothesID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Add the sale to the sales table
	saleQuery := `
		INSERT INTO sales (OrderID, ClothesID, Quantity)
		VALUES (?, ?, ?)
	`

	_, err = db.Exec(saleQuery, orderID, clothesID, sale.Quantity)
	if err != nil {
		return err
	}

	return nil
}
