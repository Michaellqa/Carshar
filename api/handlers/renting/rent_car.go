package renting

import (
	"Carshar/api/handlers/auth"
	"Carshar/api/handlers/csurl"
	"Carshar/dal"
	"Carshar/service"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"
)

type RentHandler struct {
	car  *service.CarManager
	rent *service.BookingProvider
}

func NewRentHandler(
	car *service.CarManager,
	rent *service.BookingProvider,
) RentHandler {
	return RentHandler{car: car, rent: rent}
}

func (h RentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("rent:", r.URL)

	uid, err := auth.UserToken(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(403)
		return
	}
	carId, ok := csurl.IntIdParam(r)
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

	var total float64
	if total, err = h.totalPrice(carId, rent.StartTime, rent.EndTime); err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}

	if total != rent.CalculatedTotal {
		log.Println(total, "!=", rent.CalculatedTotal)
		w.WriteHeader(409)
		return
	}

	bookId, err := h.rent.CreateBooking(rent)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}

	resp := struct {
		Id int `json:"id"`
	}{Id: bookId}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}
}

func (h RentHandler) totalPrice(carId int, start, end time.Time) (float64, error) {
	prices, err := h.car.GetPrices(carId)
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
