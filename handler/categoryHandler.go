package handler

import "fmt"

func CategoryCostume() {
	fmt.Println("Costume categories:")
	fmt.Println("1. Cosplay")
	fmt.Println("2. Formal")
	fmt.Println("Choose one (1/2).")

}

// func ListCostumes() error {
// 	query := `SELECT
// 	costumes.CostumeCategory
// FROM costumes
// GROUP BY costumes.CostumeCategory`

// }
