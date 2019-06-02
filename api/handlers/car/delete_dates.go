package car

import (
	"Carshar/api/handlers/csurl"
	"Carshar/service"
	"net/http"
)

type DeleteDatesHandler struct {
	car *service.CarManager
}

func NewDeleteDatesHandler() DeleteDatesHandler {
	return DeleteDatesHandler{}
}

// cars/{id}/dates
func (h DeleteDatesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	carId, ok := csurl.IntIdParam(r)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.car.DeleteDates(carId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
