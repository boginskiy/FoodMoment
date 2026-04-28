package service

import (
	a "mealmate/internal/alerts"
	"mealmate/internal/db"
	m "mealmate/internal/model"
	"mealmate/pkg"
	"net/http"
)

const ID = "Name"

type FoodServ struct {
	exReflecter pkg.ExReflecter
	exEncoder   pkg.ExEncoder
	exFuncer    ExFuncer
	alerter     a.Alerter
	repos       db.Repository
}

func NewFoodServ(exR pkg.ExReflecter, exE pkg.ExEncoder, a a.Alerter, r db.Repository) *FoodServ {
	exF := NewExtraFunc(exE)
	return &FoodServ{
		exReflecter: exR,
		exEncoder:   exE,
		exFuncer:    exF,
		alerter:     a,
		repos:       r}
}

func (f *FoodServ) Create(req *http.Request) ([]byte, int) {
	// Read Body
	body, err := f.exFuncer.ReadRequestBody(req)
	if err != nil {
		return f.alerter.HandleAlert(req, err, http.StatusBadRequest)
	}

	// Deserialization
	newFood := m.NewFoodModel()
	err = f.exEncoder.Deserialization(body, newFood)
	if err != nil {
		return f.alerter.HandleAlert(req, err, http.StatusBadRequest)
	}

	// Write in DB
	warn := f.repos.PutRecord(newFood)
	if warn != nil {
		return f.alerter.HandleAlert(req, warn, http.StatusUnprocessableEntity)
	}
	return body, http.StatusOK
}

/*
Params:

	?id=2              # Идентификатор блюда (по умолчанию: просмотр всех блюд)
	?category=Fastfood # Категория блюда (пример: Fastfood)
	?name=Burger       # Название блюда (пример: Burger)
	?totalPrice=10.25  # Стоимость блюда (пример: 10.25)
	?type=Vegetarian   # Тип блюда (пример: Vegetarian)

Default behavior: View all available foods.
Todo: Set up pagination.
*/
func (f *FoodServ) Read(req *http.Request) ([]byte, int) {
	needFood := make([]m.FoodModel, 0, 10) // needFood
	queryParams := req.URL.Query()         // Params of URL

	// Приведение типа "any" к "map[string]*m.FoodModel"
	tmpStore := f.repos.TakeAllStore()

	for _, food := range tmpStore.(map[string]*m.FoodModel) {
		// Filter about queryParams
		if ok := f.exFuncer.NeedShow(*food, queryParams); ok {
			needFood = append(needFood, *food)
		}
	}

	// Serialization
	tmpByte, err := f.exEncoder.Serialization(needFood)
	if err != nil {
		return f.alerter.HandleAlert(req, err, http.StatusBadRequest)
	}
	return tmpByte, http.StatusOK
}

func (f *FoodServ) Update(req *http.Request) ([]byte, int) {
	// Read Body
	body, err := f.exFuncer.ReadRequestBody(req)
	if err != nil {
		return f.alerter.HandleAlert(req, err, http.StatusBadRequest)
	}

	// Deserialization to map
	var foodForUpdate map[string]any
	err = f.exEncoder.Deserialization(body, &foodForUpdate)
	if err != nil {
		return f.alerter.HandleAlert(req, err, http.StatusBadRequest)
	}

	// Check that is ID, like 'Name'
	id, err := f.exFuncer.GetModelerID(foodForUpdate, ID)
	if err != nil {
		return f.alerter.HandleAlert(req, err, http.StatusBadRequest)
	}

	// Give updating newFood into db
	updatedFood, warn := f.repos.UpdateRecord(id, foodForUpdate)
	if warn != nil {
		return f.alerter.HandleAlert(req, warn, http.StatusBadRequest)
	}

	// Serialization
	tmpByte, err := f.exEncoder.Serialization(updatedFood)
	if err != nil {
		return f.alerter.HandleAlert(req, err, http.StatusBadRequest)
	}
	return tmpByte, http.StatusOK
}

func (f *FoodServ) Delete(req *http.Request) ([]byte, int) {
	// Пробуем достать ID == Name
	nameFood, err := f.exFuncer.TakeIDFromPath(req, ID)

	if err != nil {
		// Пробуем достать ID из Body
		nameFood, err = f.exFuncer.TakeIDFromBody(req, ID)
		if err != nil {
			return f.alerter.HandleAlert(req, err, http.StatusBadRequest)
		}
	}

	// Delete Food into db
	deletedFood, warn := f.repos.DeleteRecord(nameFood)
	if warn != nil {
		return f.alerter.HandleAlert(req, warn, http.StatusBadRequest)
	}

	// Serialization
	tmpByte, err := f.exEncoder.Serialization(deletedFood)
	if err != nil {
		f.alerter.HandleAlert(req, err, http.StatusBadRequest)
	}
	return tmpByte, http.StatusOK
}
