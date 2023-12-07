package handler

import (
	"context"
	"database/sql"
	"fmt"
	"pair-project/entity"
	"slices"
	"time"
)

func GetCategoriesProduct(db *sql.DB, product entity.Product) ([]string, error) {
	query := fmt.Sprintf("SELECT DISTINCT %s FROM %s", product.Category, product.Types)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []string
	for rows.Next() {
		var category string

		err := rows.Scan(&category)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	slices.Sort(categories)
	return categories, nil
}
