package analytics

import (
	"Carshar/dal"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

type AnalyticsHandler struct {
	db *dal.AnalyticsDb
}

func NewAnalyticsHandler(db *dal.AnalyticsDb) AnalyticsHandler {
	return AnalyticsHandler{db: db}
}

func (a AnalyticsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(400)
		return
	}
	carId_, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}
	carId := int(carId_)
	now := time.Now()
	monthAgo := now.AddDate(0, -1, 0)

	fmt.Println("analytics:", carId, monthAgo, now)

	rents := a.db.NumberOfRents(carId, monthAgo, now)
	avTime := a.db.AvailableForRent(carId, monthAgo, now)
	avgRentTime := a.db.AverageRentTime(carId, monthAgo, now)
	percentRentTime := a.db.PercentRentTime(carId, monthAgo, now)
	avgAge := a.db.AverageRenterAge(carId, monthAgo, now)

	report := Report{
		Rents:            rents,
		TimeAvailable:    avTime.String(),
		AverageRentTime:  avgRentTime.String(),
		PercentRentTime:  percentRentTime,
		AverageRenterAge: avgAge.Year(),
	}

	if err := json.NewEncoder(w).Encode(report); err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}
}

type Report struct {
	Rents            int     `json:"rents"`
	AverageRentTime  string  `json:"average_time_rent"`
	TimeAvailable    string  `json:"time_available"`
	PercentRentTime  float64 `json:"percent_rent_time"`
	AverageRenterAge int     `json:"average_renter_age"`
}
