package db

import m "mealmate/internal/model"

type FoodStore struct {
	// view table for Food
	Store map[string]*m.FoodModel
}

func NewFoodStore() *FoodStore {
	return &FoodStore{make(map[string]*m.FoodModel, 5)}
}

func (f *FoodStore) getEmptyRecord() m.Modeler {
	return m.FoodModel{}
}

func (f *FoodStore) getRecord(id string) (m.Modeler, bool) {
	value, ok := f.Store[id]
	return value, ok
}

func (f *FoodStore) getStore() any {
	return f.Store
}

func (f *FoodStore) setRecord(id string, value m.Modeler) (ok bool) {
	f.Store[id] = value.(*m.FoodModel)
	return true
}

func (f *FoodStore) deleteRecord(id string) {
	delete(f.Store, id)
}
