package service

import (
	"errors"
	"mealmate/internal/alerts/warnings"
)

var badBodyErr = errors.New("the body for the update is bad.") // Некорректное тело запроса
var notIdFieldErr = errors.New("missing ID field")             // Отсутствует поле ID
var notValidIdErr = errors.New("invalid ID format")            // Не верный формат данных у поля ID
var newWarn = warnings.New("")                                 // New
var notIdFromPathInfo = "[INFO]: the 'id' is not in url path"  // TODO. Возможно стоит переделать в структуру как warn
