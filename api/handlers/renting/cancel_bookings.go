package renting

import (
	"Carshar/api/handlers/csurl"
	"Carshar/service"
	"net/http"
)

type CancelBookingsHandler struct {
	book service.BookingProvider
}

func NewCancelBookingsHandler() CancelBookingsHandler {
	return CancelBookingsHandler{}
}

func (h CancelBookingsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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
