package entity

type Product struct {
	Types    string
	Category string
}

var ProductClothes = Product{
	Types:    "clothes",
	Category: "ClothesCategory",
}

var ProductCostume = Product{
	Types:    "costumes",
	Category: "CostumeCategory",
}
