package car

import (
	"Carshar/api/handlers/auth"
	"Carshar/dal"
	"encoding/json"
	"log"
	"net/http"
)

type AddCarHandler struct {
	db *dal.RentDb
}

func NewAddCarHandler(db *dal.RentDb) AddCarHandler {
	return AddCarHandler{db: db}
}

func (h AddCarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json charset=utf-8")

	uid, err := auth.UserToken(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(403)
		return
	}

	var car dal.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}
	car.OwnerId = uid

	added, err := h.db.CreateCar(car)
	if err != nil {
		w.WriteHeader(502)
	}
	if !added {
		w.WriteHeader(409)
	}
}
