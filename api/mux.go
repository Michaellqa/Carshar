package api

import (
	"Carshar/api/handlers/auth"
	"Carshar/api/handlers/car"
	"github.com/gorilla/mux"
	"net/http"
)

func NewMux(
	createUserHandler auth.CreateUserHandler,
	authHandler auth.AuthorizeHandler,
	addCarHandler car.AddCarHandler,
	carListHandler car.CarListHandler,
	findCarHandler car.FindCarHandler,
	addDateHandler car.AddDateHandler,
	addPriceHandler car.AddPriceHandler,
) http.Handler {
	mx := mux.NewRouter()

	mx.Handle("/users", createUserHandler).Methods(http.MethodPost)
	mx.Handle("/user", authHandler).Methods(http.MethodGet)

	mx.Handle("/cars", addCarHandler).Methods(http.MethodPost)
	mx.Handle("/cars", carListHandler).Methods(http.MethodGet)
	mx.Handle("/cars/{id}", findCarHandler).Methods(http.MethodGet)

	mx.Handle("/cars/{id}/dates", addDateHandler).Methods(http.MethodPost)
	mx.Handle("/cars/{id}/prices", addPriceHandler).Methods(http.MethodPost)

	return mx
}

//type Auth struct {
//
//}
//
//func (a Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//
//}
