package model

type FoodModel struct {
	ID          int
	Name        string
	Type        string
	Category    string
	TotalPrice  float64
	Composition []Modeler
}

func NewFoodModel() *FoodModel {
	return &FoodModel{}
}

func (f FoodModel) GetName() string {
	return f.Name
}
