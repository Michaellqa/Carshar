package api

import (
	"Carshar/api/handlers/analytics"
	"Carshar/api/handlers/auth"
	"Carshar/api/handlers/car"
	"Carshar/api/handlers/renting"
	"github.com/gorilla/mux"
	"net/http"
)

func NewMux(
	createUserHandler auth.CreateUserHandler,
	authHandler auth.AuthorizeHandler,

	addCarHandler car.AddCarHandler,
	carListHandler car.CarListHandler,
	findCarHandler car.FindCarHandler,
	userCarsHandler car.UserCarsHandler,
	userRentedCarsHandler car.UserRentedCarsHandler,
	addDateHandler car.AddDateHandler,
	addPriceHandler car.AddPriceHandler,
	dateHandler car.DatesHandler,
	priceHandler car.PricesHandler,

	totalHandler renting.TotalPriceHandler,
	rentHandler renting.RentHandler,

	analyticsHandler analytics.AnalyticsHandler,
) http.Handler {
	mx := mux.NewRouter()

	mx.Handle("/users", createUserHandler).Methods(http.MethodPost)
	mx.Handle("/user", authHandler).Methods(http.MethodGet)

	mx.Handle("/cars", addCarHandler).Methods(http.MethodPost)
	mx.Handle("/cars", carListHandler).Methods(http.MethodGet)
	mx.Handle("/cars/{id}", findCarHandler).Methods(http.MethodGet)
	mx.Handle("/users/{id}/cars/my", userCarsHandler).Methods(http.MethodGet)
	mx.Handle("/users/{id}/cars/rented", userRentedCarsHandler).Methods(http.MethodGet)

	mx.Handle("/cars/{id}/dates", addDateHandler).Methods(http.MethodPost)
	mx.Handle("/cars/{id}/prices", addPriceHandler).Methods(http.MethodPost)

	mx.Handle("/cars/{id}/dates", dateHandler).Methods(http.MethodGet)
	mx.Handle("/cars/{id}/prices", priceHandler).Methods(http.MethodGet)

	mx.Handle("/cars/{id}/{start-date}/{end-date}/total", totalHandler).Methods(http.MethodGet)
	mx.Handle("/cars/{id}/rent", rentHandler).Methods(http.MethodPost)

	mx.Handle("/cars/{id}/analytics", analyticsHandler).Methods(http.MethodGet)

	return mx
}
