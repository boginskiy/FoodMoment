package handlers

import (
	"fmt"

	s "mealmate/internal/service"
	"net/http"
	"os"
)

type MealHandler struct {
	servicer s.Servicer
}

func NewMealHandler(s s.Servicer) *MealHandler {
	return &MealHandler{servicer: s}
}

func (f *MealHandler) sendResponse(res http.ResponseWriter, body []byte, status int) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	_, err := res.Write(body)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Problem in MealHandler.sendResponse: %v\n", err)
	}
}

func (f *MealHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// Response
	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusNotFound)
	_, err := fmt.Fprint(res, "Oops! We couldn't find anything here")
	if err != nil {
		fmt.Fprintf(os.Stdout, "%v\n", err)
	}
}

func (f *MealHandler) Create(res http.ResponseWriter, req *http.Request) {
	body, status := f.servicer.Create(req)
	f.sendResponse(res, body, status)
}

func (f *MealHandler) Read(res http.ResponseWriter, req *http.Request) {
	body, status := f.servicer.Read(req)
	f.sendResponse(res, body, status)
}

func (f *MealHandler) Update(res http.ResponseWriter, req *http.Request) {
	body, status := f.servicer.Update(req)
	f.sendResponse(res, body, status)
}

func (f *MealHandler) Delete(res http.ResponseWriter, req *http.Request) {
	body, status := f.servicer.Delete(req)
	f.sendResponse(res, body, status)
}
