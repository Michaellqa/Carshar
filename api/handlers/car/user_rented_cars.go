package car

import (
	"Carshar/dal"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type UserRentedCarsHandler struct {
	db dal.CarsharRepository
}

func NewUserRentedCarsHandler(db dal.CarsharRepository) UserRentedCarsHandler {
	return UserRentedCarsHandler{db: db}
}

func (h UserRentedCarsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json charset=utf-8")

	id, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(400)
		return
	}

	uid, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	cars, err := h.db.UserRentedCars(int(uid))
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
	}

	fmt.Println("car list", r.URL, cars)

	if err := json.NewEncoder(w).Encode(cars); err != nil {
		log.Println(err)
		w.WriteHeader(502)
	}
}
