package car

import (
	"Carshar/api/handlers/auth"
	"Carshar/dal"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type AddDateHandler struct {
	db dal.CarsharRepository
}

func NewAddDateHandler(db dal.CarsharRepository) AddDateHandler {
	return AddDateHandler{db: db}
}

func (h AddDateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json charset=utf-8")

	_, err := auth.UserToken(r)
	if err != nil {
		w.WriteHeader(403)
		return
	}

	var date dal.AvailableDate
	if err := json.NewDecoder(r.Body).Decode(&date); err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}

	err = h.db.CreateDate(date.CarId, date)
	if err != nil {
		w.WriteHeader(502)
	}
}

type Date struct {
	Id        int       `json:"-"`
	DayOfWeek int       `json:"DayOfWeek"`
	TimeStart time.Time `json:"StartTim"`
	TimeEnd   time.Time `json:"TimeEnd"`
}
