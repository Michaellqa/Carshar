package renting

import (
	"Carshar/api/handlers/auth"
	"Carshar/api/handlers/csurl"
	"Carshar/dal"
	"encoding/json"
	"log"
	"net/http"
)

type CarBookingsHandler struct {
	db *dal.RentDb
}

func NewCarBookingsHandler(db *dal.RentDb) CarBookingsHandler {
	return CarBookingsHandler{db: db}
}

func (h CarBookingsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	_, err := auth.UserToken(r)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	carId, ok := csurl.IntIdParam(r)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bookings, err := h.db.CarRents(carId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(bookings) == 0 {
		return
	}

	err = json.NewEncoder(w).Encode(bookings)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
