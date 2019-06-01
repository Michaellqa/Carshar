package service

import (
	"Carshar/dal"
	"log"
)

type BookingProvider struct {
	user *dal.UserDb
	db   *dal.RentDb
	car  *CarManager
}

func NewBookingProvider(user *dal.UserDb, db *dal.RentDb, car *CarManager) *BookingProvider {
	return &BookingProvider{
		user: user,
		db:   db,
		car:  car,
	}
}

func (p *BookingProvider) CreateBooking(rent dal.Rent) (int, error) {

	car, err := p.car.Get(rent.CarId)
	if err != nil {
		log.Println(err)
		return -1, err
	}

	err = p.user.TransferMoney(rent.RenterId, car.OwnerId, rent.CalculatedTotal)
	if err != nil {
		log.Println(err)
		return -1, err
	}

	payment := dal.Payment{
		Amount:     rent.CalculatedTotal,
		SenderId:   rent.RenterId,
		ReceiverId: car.OwnerId,
	}

	paymentId, err := p.db.CreatePayment(payment)
	if err != nil {
		log.Println(err)
		return -1, err
	}

	rent.PaymentId = paymentId

	rentId, err := p.db.CreateRent(rent)
	if err != nil {
		log.Println(err)
		return -1, err
	}

	return rentId, nil
}

func (p *BookingProvider) CancelBooking() {

}
