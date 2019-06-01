package car

import (
	"Carshar/api/handlers/auth"
	"Carshar/dal"
	"Carshar/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type AddPriceHandler struct {
	car *service.CarManager
}

func NewAddPriceHandler(car *service.CarManager) AddPriceHandler {
	return AddPriceHandler{car: car}
}

func (h AddPriceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json charset=utf-8")

	id, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(400)
		return
	}
	carId, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	_, err = auth.UserToken(r)
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
	fmt.Println(price)

	err = h.car.AddPrice(int(carId), price)
	if err != nil {
		w.WriteHeader(502)
	}
}
