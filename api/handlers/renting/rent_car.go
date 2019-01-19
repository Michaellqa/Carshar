package renting

import (
	"Carshar/api/handlers/auth"
	"Carshar/dal"
	"encoding/json"
	"log"
	"net/http"
)

type RentHandler struct {
	db dal.CarsharRepository
}

func NewRentHandler(db dal.CarsharRepository) RentHandler {
	return RentHandler{db: db}
}

func (h RentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json charset=utf-8")

	uid, err := auth.UserToken(r)
	if err != nil {
		w.WriteHeader(403)
		return
	}
	carId, ok := carIdParam(r)
	if !ok {
		w.WriteHeader(400)
		return
	}

	rent := dal.Rent{}
	if err := json.NewDecoder(r.Body).Decode(&rent); err != nil {
		w.WriteHeader(400)
		return
	}
	rent.RenterId = uid
	rent.CarId = carId

	//TODO check change price

	if err = h.db.CreateRent(rent); err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}
}
