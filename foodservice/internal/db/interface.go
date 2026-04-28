package db

import (
	w "mealmate/internal/alerts/warnings"
	m "mealmate/internal/model"
)

// interface для взаимодействия слоев "Репозитория" с "Таблицами BD"
type Storer interface {
	getRecord(string) (m.Modeler, bool)
	setRecord(string, m.Modeler) bool
	getEmptyRecord() m.Modeler
	deleteRecord(string)
	getStore() any
}

// interface для взаимодействия слоев "Бизнес Логики" с "Репозиторием"
type Repository interface {
	UpdateRecord(string, map[string]any) (m.Modeler, w.Warning)
	DeleteRecord(string) (m.Modeler, w.Warning)
	TakeRecord(string) (m.Modeler, w.Warning)
	PutRecord(m.Modeler) w.Warning
	TakeAllStore() any
}
