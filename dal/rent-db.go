package dal

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"log"
	"time"
)

/*
	RemoveDate(did int) error
	RentHistory(uid int) (History, error)
*/

const (
	SqlAllAvailableCars = `
SELECT "Id", "Model", "Year"
from "Car" WHERE "OwnerId" <> $1
`
	SqlAvailableCarsForDate = `
SELECT "Id", "Model", "Year"
FROM "Car" INNER JOIN "Date" ON "Car"."Id" = "CarId"
WHERE "OwnerId" <> $1 
AND ($2 BETWEEN "StartTime" AND "EndTime")
AND ($3 BETWEEN "StartTime" AND "EndTime");
`
	SqlUserCars = `
SELECT "Id", "Model"
from "Car" WHERE "OwnerId" = $1
`
	SqlUserRentedCars = `
SELECT "Model" FROM "Car" INNER JOIN "Rent"
ON "Rent"."CarId" = "Car"."Id"
WHERE "Rent"."RenterId" = $1;
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
SELECT "StartTime", "EndTime" FROM "Date"
WHERE "CarId" = $1;
`
	SqlCreateRent = `
INSERT INTO "Rent" ("CarId", "RenterId", "StartDate", "EndDate", "TotalPrice") VALUES 
($1, $2, $3, $4, $5);
`
	SqlRentHistory = `
SELECT "CarId", "StartTime", "EndTime", "TotalPrice" FROM "Rent"
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
INSERT INTO "Date" ("CarId", "StartTime", "EndTime") VALUES 
($1, $2, $3)
`
)

type RentDb struct {
	db *sql.DB
}

func NewRentDb(db *sql.DB) *RentDb {
	return &RentDb{db: db}
}

func (r *RentDb) AvailableCars(uid int) ([]CarShortDescription, error) {
	rows, err := r.db.Query(SqlAllAvailableCars, uid)
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

func (r *RentDb) AvailableCarsForDate(uid int, start, end time.Time) ([]CarShortDescription, error) {
	fmt.Println("searching: ", uid, start, end)
	rows, err := r.db.Query(SqlAvailableCarsForDate, uid, start, end)
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

func (r *RentDb) UserCars(uid int) ([]CarRentingStatus, error) {
	rows, err := r.db.Query(SqlUserCars, uid)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var (
		car  CarRentingStatus
		cars []CarRentingStatus
	)

	for rows.Next() {
		err := rows.Scan(&car.Id, &car.Model)
		if err != nil {
			log.Println(err)
			continue
		}
		cars = append(cars, car)
	}

	return cars, err
}

func (r *RentDb) UserRentedCars(uid int) ([]CarRentingStatus, error) {
	rows, err := r.db.Query(SqlUserRentedCars, uid)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var (
		car  CarRentingStatus
		cars []CarRentingStatus
	)

	for rows.Next() {
		err := rows.Scan(&car.Id, &car.Model)
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
		err := rows.Scan(&date.StartTime, &date.EndTime)
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

func (r *RentDb) CreateCar(car Car) (bool, error) {
	_, err := r.db.Exec(SqlCreateCar, car.OwnerId, car.Model, car.Year, car.Image, car.Mileage, car.Vin)
	if err != nil {
		if err.(*pq.Error).Code == "23505" {
			return false, nil
		}
		log.Println(err)
		return false, err
	}
	return true, nil
}

func (r *RentDb) CreateRent(rent Rent) error {

	//TODO use procedure

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
	_, err := r.db.Exec(SqlCreateDate, carId, d.StartTime, d.EndTime)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *RentDb) CarDates(carId int) ([]AvailableDate, error) {
	rows, err := r.db.Query(SqlCarDates, carId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var (
		date  AvailableDate
		dates []AvailableDate
	)
	for rows.Next() {
		if err := rows.Scan(&date.StartTime, &date.EndTime); err != nil {
			log.Println(err)
			continue
		}
		dates = append(dates, date)
	}
	return dates, nil
}

func (r *RentDb) CarPrices(carId int) ([]PriceItem, error) {
	rows, err := r.db.Query(SqlCarPrices, carId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var (
		price  PriceItem
		prices []PriceItem
	)
	for rows.Next() {
		if err := rows.Scan(&price.TimeUnit, &price.Price); err != nil {
			log.Println(err)
			continue
		}
		prices = append(prices, price)
	}
	return prices, nil
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
