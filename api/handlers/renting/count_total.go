package renting

import (
	"Carshar/api/handlers/auth"
	"Carshar/dal"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

type TotalPriceHandler struct {
	db *dal.RentDb
}

func NewTotalHandler(db *dal.RentDb) TotalPriceHandler {
	return TotalPriceHandler{db: db}
}

func (h TotalPriceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json charset=utf-8")

	_, err := auth.UserToken(r)
	if err != nil {
		w.WriteHeader(403)
		return
	}

	carId, idOk := intIdParam(r)
	startTime, endTime, datesOk := dateParams(r)
	if !idOk || !datesOk {
		w.WriteHeader(400)
		return
	}

	total, err := h.totalPrice(carId, startTime, endTime)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}

	fmt.Println("total =", total.Value, "|", carId, startTime, endTime)
	if err = json.NewEncoder(w).Encode(total); err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}
}

func intIdParam(r *http.Request) (int, bool) {
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

func (h TotalPriceHandler) totalPrice(carId int, start, end time.Time) (res Total, err error) {
	prices, err := h.db.CarPrices(carId)
	if err != nil || len(prices) == 0 {
		return Total{}, err
	}

	res = findMinPrice(start, end, prices)
	return res, nil
}

func findMinPrice(start, end time.Time, prices []dal.PriceItem) (res Total) {

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
		hours = 0
		days = 0
		weeks += 1
	}

	res.Value = float64(weeks)*weekRate + float64(days)*dayRate + float64(hours)*hourRate
	if weeks > 0 {
		res.Parts = append(res.Parts, Part{Unit: "week", Base: weekRate, Count: weeks})
	}
	if days > 0 {
		res.Parts = append(res.Parts, Part{Unit: "day", Base: dayRate, Count: days})
	}
	if hours > 0 {
		res.Parts = append(res.Parts, Part{Unit: "hour", Base: hourRate, Count: hours})
	}

	return res
}

type Total struct {
	Value float64 `json:"value"`
	Parts []Part  `json:"parts"`
}

type Part struct {
	Unit  string  `json:"unit"`
	Base  float64 `json:"base"`
	Count int     `json:"count"`
}
