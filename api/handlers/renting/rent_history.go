package renting

import (
	"Carshar/dal"
	"net/http"
	"time"
)

type RentHistoryHandler struct {
	db dal.CarsharRepository
}

func NewRentHistoryHandler(db dal.CarsharRepository) RentHistoryHandler {
	return RentHistoryHandler{db: db}
}

func (h RentHistoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json charset=utf-8")

	//uid, err := auth.UserToken(r)
	//if err != nil {
	//	w.WriteHeader(403)
	//	return
	//}

	//TODO
}

type History struct {
	Model     string
	DateStart time.Time
	Total     float32
}
