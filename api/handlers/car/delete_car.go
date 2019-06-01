package car

import (
	"Carshar/dal"
	"net/http"
)

type DeleteCarHandler struct {
	db *dal.RentDb
}

//money are kept by service till booking ends, after that scheduler
//check all unfinished transactions and pay to the renter

// cars/{id}/delete
func (h *DeleteCarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//find all bookings of the car[id]

	//get that amount back to renters

	//cancel all rents

	//delete car

}
