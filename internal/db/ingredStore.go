package db

import m "mealmate/internal/model"

type IngredStore struct {
	// view table for Ingredient
	Store map[string]*m.IngredModel
}

func NewIngredStore() *IngredStore {
	return &IngredStore{make(map[string]*m.IngredModel, 5)}
}

func (i *IngredStore) getEmptyRecord() m.Modeler {
	return m.IngredModel{}
}

func (i *IngredStore) getRecord(id string) (m.Modeler, bool) {
	value, ok := i.Store[id]
	return value, ok
}

func (i *IngredStore) getStore() any {
	return i.Store
}

func (i *IngredStore) setRecord(id string, value m.Modeler) (ok bool) {
	i.Store[id] = value.(*m.IngredModel)
	return true
}

func (i *IngredStore) deleteRecord(id string) {
	delete(i.Store, id)
}
