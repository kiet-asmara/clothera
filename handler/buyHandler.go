package handler

import (
	"database/sql"
	"pair-project/entity"
)

func AddProduct(db *sql.DB, clothes entity.Clothes, customer entity.Customer, order entity.Order, sale entity.Sale) error {

	// Get the last inserted customer ID
	customerID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Get the last inserted order ID
	orderID, err := result.LastInsertId()
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
