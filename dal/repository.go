package dal

type AuthRepository interface {
	CreateUser(user User) (bool, error)
	//returned bool indicates if user was found
	FindUser(phone string) (User, bool, error)
}

type CarsharRepository interface {
	AvailableCars(uid int) ([]CarShortDescription, error)
	FindCar(cid int) (CarFullDescription, error)
	CreateRent(r Rent) error

	CreateCar(c Car) error
	CreateDate(carId int, d AvailableDate) error
	CreatePrice(carId int, d PriceItem) error
	//AddAvailableDate(date AvailableDate) error
	//RemoveDate(did int) error

	//RentHistory(uid int) (History, error)
}
