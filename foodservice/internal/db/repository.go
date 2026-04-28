package db

import (
	w "mealmate/internal/alerts/warnings"
	m "mealmate/internal/model"
	"mealmate/pkg"
	"strings"
	"sync"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Repos struct {
	exrefl pkg.ExReflecter
	muR    sync.RWMutex
	mu     sync.Mutex
	store  Storer
}

func NewRep(exreflecter pkg.ExReflecter, storer Storer) *Repos {
	return &Repos{exrefl: exreflecter, store: storer}
}

func (r *Repos) checkForUnic(id string) bool {
	r.muR.RLock()
	defer r.muR.RUnlock()
	if _, ok := r.store.getRecord(id); ok {
		return false
	}
	return true
}

func (r *Repos) TakeAllStore() any {
	return r.store.getStore()
}

func (r *Repos) PutRecord(value m.Modeler) w.Warning {
	// Формируем ключ Food
	key := strings.TrimSpace(value.GetName())
	// Check for inic
	isUnic := r.checkForUnic(key)
	if !isUnic {
		return notUnicFoodWarn
	}
	// Делаем новую запись с Food
	r.mu.Lock()
	r.store.setRecord(key, value)
	r.mu.Unlock()
	return nil
}

func (r *Repos) TakeRecord(id string) (m.Modeler, w.Warning) {
	r.muR.RLock()
	defer r.muR.RUnlock()
	tookRecord, ok := r.store.getRecord(id)
	if !ok {
		return r.store.getEmptyRecord(), notFoundFoodWarn
	}
	return tookRecord, nil
}

func (r *Repos) UpdateRecord(id string, partForUpdate map[string]any) (m.Modeler, w.Warning) {
	r.mu.Lock()
	defer r.mu.Unlock()

	tookRecord, ok := r.store.getRecord(id)
	if !ok {
		return r.store.getEmptyRecord(), notFoundFoodWarn
	}

	// Update old Record on new
	tookRecord = r.exrefl.CrossUpdateStructs(tookRecord, partForUpdate).(m.Modeler)
	return tookRecord, nil
}

func (r *Repos) DeleteRecord(id string) (m.Modeler, w.Warning) {
	// TODO. Временно так
	titleCase := cases.Title(language.Russian)
	id = titleCase.String(id)

	r.mu.Lock()
	defer r.mu.Unlock()

	if deletingRecord, ok := r.store.getRecord(id); ok {
		r.store.deleteRecord(id)
		return deletingRecord, nil
	}
	return r.store.getEmptyRecord(), notFoundFoodWarn
}
