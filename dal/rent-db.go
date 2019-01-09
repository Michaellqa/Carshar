package dal

import (
	"database/sql"
	"log"
)

/*
	RemoveDate(did int) error
	RentHistory(uid int) (History, error)
*/

const (
	SqlAvailableCars = `
-- SELECT "Car"."Id", "Model", "Year"
-- FROM "Car" RIGHT JOIN "Date"
-- ON "Car"."Id" = "Date"."CarId"
-- WHERE "Car"."OwnerId" <> $1;
SELECT "Id", "Model", "Year"
from "Car" WHERE "OwnerId" <> $1
`
	SqlFindCar = `
SELECT "Id", "Model", "Year", "Image", "Mileage" FROM "Car"
WHERE "Id" = $1;
`
	SqlCarPrices = `
SELECT "TimeUnit", "Price" FROM "Price"
WHERE "CarId" = $1;
`
	SqlCarDates = `
SELECT "DayOfWeek", "TimeStart", "TimeEnd" FROM "Date"
WHERE "CarId" = $1;
`
	SqlCreateRent = `
INSERT INTO "Rent" ("CarId", "RenterId", "TimeStart", "TimeEnd", "TotalPrice") VALUES 
($1, $2, $3, $4, $5);
`
	SqlRentHistory = `
SELECT "CarId", "TimeStart", "TimeEnd", "TotalPrice" FROM "Rent"
WHERE "UserId" = $1;
`
	SqlCreateCar = `
INSERT INTO "Car" ("OwnerId", "Model", "Year", "Image", "Mileage", "Vin") VALUES 
($1, $2, $3, $4, $5, $6)
`
	SqlCreatePrice = `
INSERT INTO "Price" ("CarId", "TimeUnit", "Price") VALUES 
($1, $2, $3)
`
	SqlCreateDate = `
INSERT INTO "Date" ("CarId", "DayOfWeek", "StartTime", "EndTime") VALUES 
($1, $2, $3, $4)
`
)

type RentDb struct {
	db *sql.DB
}

func NewRentDb(db *sql.DB) *RentDb {
	return &RentDb{db: db}
}

func (r *RentDb) AvailableCars(uid int) ([]CarShortDescription, error) {
	rows, err := r.db.Query(SqlAvailableCars, uid)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var (
		car  CarShortDescription
		cars []CarShortDescription
	)

	for rows.Next() {
		err := rows.Scan(&car.Id, &car.Model, &car.Year)
		if err != nil {
			log.Println(err)
			continue
		}
		cars = append(cars, car)
	}

	return cars, err
}

func (r *RentDb) FindCar(id int) (car CarFullDescription, err error) {

	row := r.db.QueryRow(SqlFindCar, id)

	err = row.Scan(&car.Id, &car.Model, &car.Year, &car.Image, &car.Mileage)
	if err != nil {
		log.Println(err)
		return CarFullDescription{}, err
	}

	rows, err := r.db.Query(SqlCarDates, id)
	if err != nil {
		log.Println(err)
		return CarFullDescription{}, err
	}
	for rows.Next() {
		date := AvailableDate{}
		err := rows.Scan(&date.DayOfWeek, &date.StartTime, &date.EndTime)
		if err != nil {
			log.Println(err)
			continue
		}
		car.Dates = append(car.Dates, date)
	}

	rows, err = r.db.Query(SqlCarPrices, id)
	if err != nil {
		log.Println(err)
		return CarFullDescription{}, err
	}
	for rows.Next() {
		p := PriceItem{}
		err := rows.Scan(&p.TimeUnit, &p.Price)
		if err != nil {
			log.Println(err)
			continue
		}
		car.Prices = append(car.Prices, p)
	}

	return car, nil
}

func (r *RentDb) CreateCar(car Car) error {
	_, err := r.db.Exec(SqlCreateCar, car.OwnerId, car.Model, car.Year, car.Image, car.Mileage, car.Vin)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *RentDb) CreateRent(rent Rent) error {
	_, err := r.db.Exec(SqlCreateRent, rent.CarId, rent.RenterId, rent.StartTime, rent.EndTime, rent.Total)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *RentDb) CreatePrice(carId int, p PriceItem) error {
	_, err := r.db.Exec(SqlCreatePrice, carId, p.TimeUnit, p.Price)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *RentDb) CreateDate(carId int, d AvailableDate) error {
	_, err := r.db.Exec(SqlCreateDate, carId, d.DayOfWeek, d.StartTime, d.EndTime)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *RentDb) RentHistory(uid int) ([]Rent, error) {
	rows, err := r.db.Query(SqlRentHistory, uid)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	rents := make([]Rent, 0)
	for rows.Next() {
		r := Rent{}
		err := rows.Scan(&r.CarId, &r.RenterId, &r.StartTime, &r.EndTime, &r.Total)
		if err != nil {
			log.Println(err)
			continue
		}
		rents = append(rents, r)
	}

	return rents, nil
}
