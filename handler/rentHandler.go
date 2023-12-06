package handler

import (
	"database/sql"
	"fmt"
	"pair-project/entity"
	"time"
)

func Rent(db *sql.DB, orderID int) (float64, error) {

	// show menu & take input
	rentDetail, mssg, err := RentInput(db, orderID)
	if err != nil {
		return 0, fmt.Errorf("Rent: %w", err)
	} else if mssg != "" {
		fmt.Println(mssg)
		fmt.Println("")
		return 0, nil
	}

	// reduce stock
	err = ReduceStock(db, rentDetail.CostumeID, rentDetail.Quantity)
	if err != nil {
		return 0, fmt.Errorf("Rent: %w", err)
	}

	// insert data
	err = RentInsert(db, rentDetail)
	if err != nil {
		return 0, fmt.Errorf("Rent: %w", err)
	}

	// calculate price
	price, err := RentPrice(db, rentDetail)
	if err != nil {
		return 0, fmt.Errorf("Rent: %w", err)
	}

	return price, nil
}

func RentInput(db *sql.DB, orderID int) (entity.Rent, string, error) {

	// get input
	var costumeID, quantity int

	fmt.Println("Choose costume ID:")
	fmt.Scan(&costumeID)
	fmt.Println("How many costumes:")
	fmt.Scan(&quantity)

	// check stock
	stock, err := CostumeStock(db, costumeID, quantity)
	if err != nil {
		return entity.Rent{}, "", fmt.Errorf("rentMenu: %w", err)
	}
	if stock == 1 {
		return entity.Rent{}, "Item out of stock", nil
	} else if stock == 2 {
		return entity.Rent{}, "Item stock is not enough", nil
	}

	var start, end string

	fmt.Println("Insert rental date (format: 2023-05-09).")
	fmt.Println("Start:")
	fmt.Scan(&start)
	fmt.Println("End:")
	fmt.Scan(&end)

	// check date
	days := DaysBetween(start, end)
	if days == 0 {
		return entity.Rent{}, "", fmt.Errorf("rentMenu: %w", err)
	}

	rental := entity.Rent{
		OrderID:   orderID,
		CostumeID: costumeID,
		Quantity:  quantity,
		StartDate: start,
		EndDate:   end,
	}

	return rental, "", nil
}

func RentInsert(db *sql.DB, rent entity.Rent) error {
	query := `INSERT INTO rents (orderid,costumeid,quantity,startdate,enddate) VALUES
	(?,?,?,?,?)`

	_, err := db.Exec(query, rent.OrderID, rent.CostumeID, rent.Quantity, rent.StartDate, rent.EndDate)
	if err != nil {
		return fmt.Errorf("RentInsert: %w", err)
	}

	fmt.Println("Rental order created")
	return nil
}

func RentPrice(db *sql.DB, rent entity.Rent) (float64, error) {
	// get price per day
	query := `SELECT
	costumes.CostumePrice
FROM costumes
JOIN rents ON costumes.CostumeID = rents.CostumeID
WHERE costumes.costumeID = ?`

	rows, err := db.Query(query, rent.CostumeID)
	if err != nil {
		return 0, fmt.Errorf("RentPrice: %w", err)
	}
	defer rows.Close()

	var priceDay float64

	for rows.Next() {
		err := rows.Scan(&priceDay)
		if err != nil {
			return 0, fmt.Errorf("RentPrice: %w", err)
		}
	}

	// get number of days
	days := DaysBetween(rent.StartDate, rent.EndDate)
	if days == 0 {
		return 0, nil
	}

	// calculate price
	totalPrice := priceDay * float64(days) * float64(rent.Quantity)

	return totalPrice, nil
}

func DaysBetween(start, end string) int {
	timeFormat := "2006-01-02"

	startDate, err := time.Parse(timeFormat, start)
	if err != nil {
		fmt.Println("Invalid date format.")
		return 0
	}

	endDate, err := time.Parse(timeFormat, end)
	if err != nil {
		fmt.Println("Invalid date format.")
		return 0
	}

	if startDate.After(endDate) {
		fmt.Println("invalid date: start date is after end date.")
		return 0
	} else if startDate.Equal(endDate) {
		return 1
	}

	days := int(endDate.Sub(startDate).Hours() / 24)

	return days
}

func CostumeStock(db *sql.DB, costumeID int, quantity int) (int, error) {
	query := `SELECT
	costumes.CostumeStock
FROM costumes
WHERE costumes.CostumeID = ?`

	// get costume stock
	rows, err := db.Query(query, costumeID)
	if err != nil {
		return 0, fmt.Errorf("CostumeStock: %w", err)
	}
	defer rows.Close()

	var stock int

	for rows.Next() {
		err := rows.Scan(&stock)
		if err != nil {
			return 0, fmt.Errorf("CostumeStock: %w", err)
		}
	}

	// check stock
	if stock == 0 {
		return 1, nil
	} else if stock < quantity {
		return 2, nil
	}
	return 0, nil
}

func ReduceStock(db *sql.DB, costumeID int, quantity int) error {
	query := `UPDATE costumes
	SET CostumeStock = CostumeStock - ?
	WHERE CostumeID = ?`

	_, err := db.Exec(query, quantity, costumeID)
	if err != nil {
		return fmt.Errorf("ReduceStock: %w", err)
	}

	fmt.Println("Stock reduced")
	return nil
}
