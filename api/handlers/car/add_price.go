package car

import (
	"Carshar/api/handlers/auth"
	"Carshar/dal"
	"encoding/json"
	"log"
	"net/http"
)

type AddPriceHandler struct {
	db dal.CarsharRepository
}

func NewAddPriceHandler(db dal.CarsharRepository) AddPriceHandler {
	return AddPriceHandler{db: db}
}

func (h AddPriceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json charset=utf-8")

	_, err := auth.UserToken(r)
	if err != nil {
		w.WriteHeader(403)
		return
	}

	var price dal.PriceItem
	if err := json.NewDecoder(r.Body).Decode(&price); err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}

	err = h.db.CreatePrice(price.CarId, price)
	if err != nil {
		w.WriteHeader(502)
	}
}
