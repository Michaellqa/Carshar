package car

import (
	"Carshar/api/handlers/csurl"
	"Carshar/service"
	"net/http"
)

type DeleteCarHandler struct {
	car *service.CarManager
}

func NewDeleteCarHandler(car *service.CarManager) DeleteCarHandler {
	return DeleteCarHandler{car: car}
}

//money are kept by service till booking ends, after that scheduler
//check all unfinished transactions and pay to the renter

// cars/{id} [Delete]
func (h DeleteCarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	carId, ok := csurl.IntIdParam(r)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.car.DeleteCar(carId); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}
