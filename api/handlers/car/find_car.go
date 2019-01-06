package car

import (
	"Carshar/dal"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type FindCarHandler struct {
	db dal.CarsharRepository
}

func NewFindCarHandler(db dal.CarsharRepository) FindCarHandler {
	return FindCarHandler{db: db}
}

func (h FindCarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	desc, err := h.db.FindCar(int(carId))
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
	}

	if err := json.NewEncoder(w).Encode(desc); err != nil {
		log.Println(err)
		w.WriteHeader(502)
	}
}
