package car

import (
	"CarShar/api/handlers/auth"
	"CarShar/dal"
	"encoding/json"
	"log"
	"net/http"
)

type AddDateHandler struct {
	db dal.CarsharRepository
}

func NewAddDateHandler(db dal.CarsharRepository) AddDateHandler {
	return AddDateHandler{db: db}
}

func (h AddDateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json charset=utf-8")

	uid, err := auth.UserToken(r)
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

	err = h.db.CreateDate(uid, date)
	if err != nil {
		w.WriteHeader(502)
	}
}
