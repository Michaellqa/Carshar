package renting

import (
	"Carshar/api/handlers/auth"
	"Carshar/dal"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"
)

type RentHandler struct {
	db dal.CarsharRepository
}

func NewRentHandler(db dal.CarsharRepository) RentHandler {
	return RentHandler{db: db}
}

/* Accept json:
{
	"startTime"
	"endTime"
	"total": float64
}
*/

func (h RentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json charset=utf-8")
	fmt.Println("rent:", r.URL)

	uid, err := auth.UserToken(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(403)
		return
	}
	carId, ok := carIdParam(r)
	if !ok {
		w.WriteHeader(400)
		return
	}

	rent := dal.Rent{}
	if err := json.NewDecoder(r.Body).Decode(&rent); err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}
	rent.RenterId = uid
	rent.CarId = carId

	//fmt.Println(rent)

	var total float64
	if total, err = h.totalPrice(carId, rent.StartTime, rent.EndTime); err != nil {
		log.Println(err)
		w.WriteHeader(409)
		return
	}

	if total != rent.Total {
		log.Println(total, "!=", rent.Total)
		return
	}

	if err = h.db.CreateRent(rent); err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}
}

func (h RentHandler) totalPrice(carId int, start, end time.Time) (float64, error) {
	prices, err := h.db.CarPrices(carId)
	if err != nil || len(prices) == 0 {
		return 0, err
	}

	duration := end.Sub(start)
	hours := int(math.Ceil(duration.Hours()))
	days := hours / 24
	weeks := days / 7

	hourRate := math.MaxFloat64
	dayRate := math.MaxFloat64
	weekRate := math.MaxFloat64

	for _, p := range prices {
		switch p.TimeUnit {
		case "hour":
			hourRate = p.Price
		case "day":
			dayRate = p.Price
		case "week":
			weekRate = p.Price
		}
	}

	dayRate = math.Min(dayRate, hourRate*24)
	weekRate = math.Min(weekRate, dayRate*7)

	hours %= 24
	days %= 7

	if dayRate < float64(hours)*hourRate {
		hours = 0
		days += 1
	}
	if weekRate < float64(days)*dayRate {
		days = 0
		weeks += 1
	}

	total := float64(weeks)*weekRate + float64(days)*dayRate + float64(hours)*hourRate
	return total, nil
}
