package entity

type Clothes struct {
	ClothesID       int
	ClothesName     string
	ClothesCategory string
	ClothesPrice    float64
	ClothesStock    int
}

type RevenueByClothes struct {
	Name         string
	Quantity     int
	TotalRevenue float64
}