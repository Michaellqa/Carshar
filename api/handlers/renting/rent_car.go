package renting

import (
	"Carshar/dal"
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

	//uid, err := auth.UserToken(r)
	//if err != nil {
	//	w.WriteHeader(403)
	//	return
	//}

	//TODO use procedure
}
