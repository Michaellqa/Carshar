package renting

import (
	"CarShar/api/handlers/auth"
	"CarShar/dal"
	"net/http"
)

type TotalPriceHandler struct {
	db dal.CarsharRepository
}

func NewTotalHandler(db dal.CarsharRepository) TotalPriceHandler {
	return TotalPriceHandler{db: db}
}

func (h TotalPriceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json charset=utf-8")

	uid, err := auth.UserToken(r)
	if err != nil {
		w.WriteHeader(403)
		return
	}

	// TODO
}
