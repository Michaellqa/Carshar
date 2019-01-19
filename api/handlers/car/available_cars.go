package car

import (
	"Carshar/dal"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
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

	cars := make([]dal.CarShortDescription, 0)

	start, end, ok := dateParams(r)
	if ok {
		log.Println("AvailableCarsForDate")
		cars, err = h.db.AvailableCarsForDate(int(uid), start, end)
	} else {
		log.Println("AvailableCars")
		cars, err = h.db.AvailableCars(int(uid))
	}

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

func dateParams(r *http.Request) (start, end time.Time, ok bool) {
	s, ok := r.URL.Query()["start"]
	if !ok {
		return time.Time{}, time.Time{}, false
	}
	e, ok := r.URL.Query()["end"]
	if !ok {
		return time.Time{}, time.Time{}, false
	}
	if len(s) == 0 || len(e) == 0 {
		return time.Time{}, time.Time{}, false
	}
	start, err := time.Parse("2006-01-02T15:04Z", s[0])
	if err != nil {
		log.Println(err)
		return time.Time{}, time.Time{}, false
	}
	end, err = time.Parse("2006-01-02T15:04Z", e[0])
	if err != nil {
		log.Println(err)
		return time.Time{}, time.Time{}, false
	}
	return start, end, true
}
