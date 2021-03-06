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
	db *dal.RentDb
}

func NewFindCarHandler(db *dal.RentDb) FindCarHandler {
	return FindCarHandler{db: db}
}

func (h FindCarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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
		return
	}

	if err := json.NewEncoder(w).Encode(desc); err != nil {
		log.Println(err)
		w.WriteHeader(502)
	}
}
