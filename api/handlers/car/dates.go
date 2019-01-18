package car

import (
	"Carshar/api/handlers/auth"
	"Carshar/dal"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type DatesHandler struct {
	db dal.CarsharRepository
}

func NewDatesHandler(db dal.CarsharRepository) DatesHandler {
	return DatesHandler{db: db}
}

func (h DatesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	cars, err := h.db.CarDates(int(carId))
	if err != nil {
		log.Println(err)
		w.WriteHeader(502)
		return
	}

	fmt.Println("car list", r.URL, cars)

	if err := json.NewEncoder(w).Encode(cars); err != nil {
		log.Println(err)
		w.WriteHeader(502)
	}
}