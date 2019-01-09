package car

import (
	"Carshar/dal"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type CarListHandler struct {
	db dal.CarsharRepository
}

func NewCarListHandler(db dal.CarsharRepository) CarListHandler {
	return CarListHandler{db: db}
}

func (h CarListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json charset=utf-8")

	idStr := r.Header.Get("Authorization")
	uid, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}

	cars, err := h.db.AvailableCars(int(uid))
	if err != nil {
		log.Println(err)
		w.WriteHeader(502)
		return
	}

	if err := json.NewEncoder(w).Encode(cars); err != nil {
		log.Println(err)
		w.WriteHeader(502)
	}
}
