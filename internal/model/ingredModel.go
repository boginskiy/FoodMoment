package model

// TODO
// Вставить калорийность на единицу, на 100 гр, на 100 мл
// Общая калорийность ингредиента в продукте

type IngredModel2 struct {
	ID             int
	Name           string
	Unit           string  // Unit is unit of measurement
	Quantity       float64 // Quantity is quantity ingredient in food
	CostOfUnit     float64 // CostOfUnit is cost of unit ingredient
	CostOfQuantity float64 // CostOfQuantity is all cost ingredient
}

type IngredModel struct {
	ID             int
	Name           string
	Unit           string  // Unit is unit of measurement // Now Show for client
	Quantity       float64 // Quantity is quantity ingredient in food
	CostOfUnit     float64 // CostOfUnit is cost of unit ingredient // Now Show for client
	CostOfQuantity float64 // CostOfQuantity is all cost ingredient
}

func NewIngredModel() *IngredModel {
	return &IngredModel{}
}

func (f IngredModel) GetName() string {
	return f.Name
}
