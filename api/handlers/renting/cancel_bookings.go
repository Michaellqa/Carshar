package renting

import (
	"Carshar/api/handlers/csurl"
	"Carshar/service"
	"log"
	"net/http"
)

type CancelBookingsHandler struct {
	book *service.BookingProvider
}

func NewCancelBookingsHandler(provider *service.BookingProvider) CancelBookingsHandler {
	return CancelBookingsHandler{book: provider}
}

func (h CancelBookingsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RequestURI)

	carId, ok := csurl.IntIdParam(r)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.book.CancelBookings(carId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
