package service

import "Carshar/dal"

type BookingProvider struct {
	db  dal.RentDb
	car CarManager
}

func (p *BookingProvider) CreateBooking() {

}

func (p *BookingProvider) CancelBooking() {

}
