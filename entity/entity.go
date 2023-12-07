package entity

type Product struct {
	Types    string
	Category string
}

var ProductClothes = Product{
	Types:    "Clothes",
	Category: "ClothesCategory",
}

var ProductCostume = Product{
	Types:    "Costumes",
	Category: "CostumeCategory",
}
