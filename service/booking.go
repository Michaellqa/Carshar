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

//todo: do the same with sql
func (p *BookingProvider) CancelBookings(carId int) error {
	rents, err := p.db.CarRents(carId)
	if err != nil {
		log.Println(err)
		return err
	}

	for _, r := range rents {
		payment, err := p.db.Payment(r.PaymentId)
		if err != nil {
			log.Println(err)
			return err
		}

		reverted := dal.Payment{
			SenderId:   payment.ReceiverId,
			ReceiverId: payment.SenderId,
			Amount:     payment.Amount,
		}

		err = p.user.TransferMoney(reverted.SenderId, reverted.ReceiverId, reverted.Amount)
		if err != nil {
			log.Println(err)
			return err
		}

		_, err = p.db.CreatePayment(reverted)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return p.db.CancelRentsOfCar(carId)
}
