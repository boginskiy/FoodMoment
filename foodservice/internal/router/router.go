package router

import (
	"context"

	"github.com/go-chi/chi"
)

type Router struct {
	*chi.Mux
}

func NewRouter(ctx context.Context) *Router {
	return &Router{
		chi.NewRouter(),
	}
}

func (r *Router) Root() {
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/food", handler1)
		r.Route("/menu", handler2)
		r.Route("/user", handler3)
	})
}

// TODO ...
// Какие маршруты нужны ?
// food - вся еда
// category - фильтрация по категории
// популярная еду по фильтру выше ...
//
//
//

// Сервис Лайки
//
//

// Сервис склад:
// учета расходов продуктов
// Информирование о необходимости дозаказа

// Сервис Pay:
// Оплата на сайте

// Сервис Доставка:
// оповещение о доставке, курьер

// Сервис Info:
// Информирование об акциях, увеличение кешбека, персональные скидки, ДР
