package model

type Food struct {
	ID   int
	Name string

	Ingredients
	Cost
}

type Product struct {
	ID   int
	Name string
	Unit string
	Cost float64
}
