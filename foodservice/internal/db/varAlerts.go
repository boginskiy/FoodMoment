package db

import (
	"errors"
	"mealmate/internal/alerts/warnings"
)

var someErr = errors.New("")                                                      // Переменная для создания будущих error
var notFoundFoodWarn = warnings.New("not found food, check fields 'name'")        // Не найден объект для обновления
var notUnicFoodWarn = warnings.New("added food is not unic, check fields 'name'") // Добавляемый объект не уникален, нужно изменить 'name'
