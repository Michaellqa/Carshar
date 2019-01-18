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

type AddDateHandler struct {
	db dal.CarsharRepository
}

func NewAddDateHandler(db dal.CarsharRepository) AddDateHandler {
	return AddDateHandler{db: db}
}

func (h AddDateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	//b, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//fmt.Println(string(b))

	var date dal.AvailableDate
	if err := json.NewDecoder(r.Body).Decode(&date); err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}

	fmt.Println(carId, date)

	err = h.db.CreateDate(int(carId), date)
	if err != nil {
		w.WriteHeader(502)
	}
}
