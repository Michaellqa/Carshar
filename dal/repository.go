package dal

import "time"

type AuthRepository interface {
	CreateUser(user User) (bool, error)
	//returned bool indicates if user was found
	FindUser(phone string) (User, bool, error)
}

type CarsharRepository interface {
	AvailableCarsForDate(uid int, start, end time.Time) ([]CarShortDescription, error)
	FindCar(cid int) (CarFullDescription, error)
	UserCars(uid int) ([]CarRentingStatus, error)
	UserRentedCars(uid int) ([]CarRentingStatus, error)
	CreateRent(r Rent) error

	CreateCar(c Car) (bool, error)
	CreateDate(carId int, d AvailableDate) error
	CreatePrice(carId int, d PriceItem) error

	CarDates(carId int) ([]AvailableDate, error)
	CarPrices(carId int) ([]PriceItem, error)
}
