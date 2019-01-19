package renting

import (
	"Carshar/api/handlers/auth"
	"Carshar/dal"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

type TotalPriceHandler struct {
	db dal.CarsharRepository
}

func NewTotalHandler(db dal.CarsharRepository) TotalPriceHandler {
	return TotalPriceHandler{db: db}
}

func (h TotalPriceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json charset=utf-8")

	_, err := auth.UserToken(r)
	if err != nil {
		w.WriteHeader(403)
		return
	}

	carId, idOk := carIdParam(r)
	startTime, endTime, datesOk := dateParams(r)
	if !idOk || !datesOk {
		w.WriteHeader(400)
		return
	}

	fmt.Println("total:", carId, startTime, endTime)

	str := strconv.FormatInt(322, 10)
	w.Write([]byte(str))
}

func carIdParam(r *http.Request) (int, bool) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		return -1, false
	}
	carId, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return -1, false
	}
	return int(carId), true
}

func dateParams(r *http.Request) (start, end time.Time, ok bool) {
	s, ok := mux.Vars(r)["start-date"]
	if !ok {
		return time.Time{}, time.Time{}, false
	}
	e, ok := mux.Vars(r)["end-date"]
	if !ok {
		return time.Time{}, time.Time{}, false
	}
	if len(s) == 0 || len(e) == 0 {
		return time.Time{}, time.Time{}, false
	}

	timeFormat := "2006-01-02T15:04Z"
	start, err := time.Parse(timeFormat, s)
	if err != nil {
		log.Println(err)
		return time.Time{}, time.Time{}, false
	}
	end, err = time.Parse(timeFormat, e)
	if err != nil {
		log.Println(err)
		return time.Time{}, time.Time{}, false
	}
	return start, end, true
}
