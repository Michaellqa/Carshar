package car

import (
	"Carshar/dal"
	"encoding/json"
	"log"
	"net/http"
)

type AddCarHandler struct {
	db dal.CarsharRepository
}

func NewAddCarHandler(db dal.CarsharRepository) AddCarHandler {
	return AddCarHandler{db: db}
}

func (h AddCarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json charset=utf-8")

	var car dal.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}

	err := h.db.CreateCar(car)
	if err != nil {
		w.WriteHeader(502)
	}
}
