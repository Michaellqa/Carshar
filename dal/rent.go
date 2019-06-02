package dal

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"log"
	"time"
)

const (
	SqlAvailableCarsForDate = `
SELECT ac."Id", "Model", "Year", "Latitude", "Longitude"
FROM "Location" RIGHT JOIN (
	SELECT "Car"."Id", "Model", "Year"
	FROM "Car" INNER JOIN "Availability" ON "Car"."Id" = "CarId"
  WHERE "OwnerId" <> $1
    AND ($2 BETWEEN "TimeStart" AND "TimeEnd")
    AND ($3 BETWEEN "TimeStart" AND "TimeEnd")
) as ac ON ac."Id" = "Location"."CarId";
;
`
	SqlUserCars = `
SELECT "Id", "Model"
from "Car" WHERE "OwnerId" = $1
`
	SqlUserRentedCars = `
SELECT "Car"."Id", "Model" FROM "Car" INNER JOIN "Reservation"
ON "Reservation"."CarId" = "Car"."Id"
WHERE "Reservation"."RenterId" = $1;
`
	SqlCarPrices = `
SELECT "Hour", "Day", "Week" FROM "Price"
WHERE "CarId" = $1;
`
	SqlCarDates = `
SELECT "TimeStart", "TimeEnd" FROM "Availability"
WHERE "CarId" = $1;
`
	SqlRentHistory = `
SELECT "CarId", "TimeStart", "TimeEnd", "TotalPrice" FROM "Reservation"
WHERE "UserId" = $1;
`
	SqlCreateCar = `
INSERT INTO "Car" ("OwnerId", "Model", "Year", "Mileage", "Vin") VALUES 
($1, $2, $3, $4, $5)
`

	//	SqlCreatePrice = `
	//INSERT INTO "Price" ("CarId", "TimeUnit", "Price") VALUES
	//($1, $2, $3)
	//`

)

type RentDb struct {
	db *sql.DB
}

func NewRentDb(db *sql.DB) *RentDb {
	return &RentDb{db: db}
}

func (r *RentDb) AvailableCarsForDate(uid int, start, end time.Time) ([]CarShortDescription, error) {
	fmt.Println("AvailableCarsForDate: ", uid, start, end)

	rows, err := r.db.Query(SqlAvailableCarsForDate, uid, start, end)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var (
		car  CarShortDescription
		cars []CarShortDescription
	)

	var lat, long sql.NullFloat64

	for rows.Next() {
		err := rows.Scan(&car.Id, &car.Model, &car.Year, &lat, &long)
		if err != nil {
			log.Println(err)
			continue
		}
		if lat.Valid && long.Valid {
			car.Coordinates.Latitude = lat.Float64
			car.Coordinates.Longitude = long.Float64
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

//todo: move to car service
func (r *RentDb) FindCar(id int) (car CarFullDescription, err error) {
	SqlFindCar := `SELECT "Id", "OwnerId", "Model", "Year", "ImageUrl" FROM "Car"WHERE "Id" = $1;`
	row := r.db.QueryRow(SqlFindCar, id)

	var imageUrl sql.NullString
	err = row.Scan(&car.Id, &car.OwnerId, &car.Model, &car.Year, &imageUrl)
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
		var hour, day, week sql.NullFloat64
		err := rows.Scan(&hour, &day, &week)
		if err != nil {
			log.Println(err)
			continue
		}

		if hour.Valid {
			car.Prices = append(car.Prices, PriceItem{id, "hour", hour.Float64})
		}
		if day.Valid {
			car.Prices = append(car.Prices, PriceItem{id, "day", hour.Float64})
		}
		if week.Valid {
			car.Prices = append(car.Prices, PriceItem{id, "week", hour.Float64})
		}
	}

	return car, nil
}

func (r *RentDb) CreateCar(car Car) (bool, error) {
	_, err := r.db.Exec(SqlCreateCar, car.OwnerId, car.Model, car.Year, car.Mileage, car.Vin)
	if err != nil {
		if err.(*pq.Error).Code == "23505" {
			return false, nil
		}
		log.Println(err)
		return false, err
	}
	return true, nil
}

func (r *RentDb) DeleteCar(id int) error {
	SqlDeleteCar := `DELETE FROM "Car" WHERE "Id" = $1;`
	_, err := r.db.Exec(SqlDeleteCar, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *RentDb) RentsOfTheCar(carId int) ([]Rent, error) {

	query := `SELECT "RenterId", "CalculatedTotalPrice" FROM "Reservation" WHERE "CarId" = 2;`
	var rents []Rent

	rows, err := r.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {
		r := Rent{}
		err = rows.Scan(r.RenterId, r.CalculatedTotal)
		if err != nil {
			log.Println("car was missed", err)
			continue
		}
		rents = append(rents, r)
	}
	return rents, nil
}

func (r *RentDb) DeleteAvailabilityForCar(carId int) error {
	query := `DELETE FROM "Availability" WHERE "CarId" = $1`
	_, err := r.db.Exec(query, carId)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *RentDb) DeletePricesForCar(carId int) error {
	SqlDeleteCar := `DELETE FROM "Price" WHERE "CarId" = $1;`
	_, err := r.db.Exec(SqlDeleteCar, carId)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *RentDb) CreateRent(rent Rent) (int, error) {
	SqlCreateRent := `
INSERT INTO "Reservation" ("RenterId", "CarId", "StartDate", "EstimatedEndDate", "CalculatedTotalPrice", "PaymentId") VALUES 
($1, $2, $3, $4, $5, $6) RETURNING "Id";`
	var rentId int
	err := r.db.QueryRow(SqlCreateRent,
		rent.RenterId,
		rent.CarId,
		rent.StartTime,
		rent.EndTime,
		rent.CalculatedTotal,
		rent.PaymentId).Scan(&rentId)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return rentId, nil
}

func (r *RentDb) CancelRent(id int) error {
	sqlCancelPayment := `UPDATE "Reservation" SET "Status" = "cancelled" WHERE "Id" = $1;`
	_, err := r.db.Exec(sqlCancelPayment, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func (r *RentDb) CancelRentsOfCar(carId int) error {
	sqlCancelRents := `UPDATE "Reservation" SET "Status" = "cancelled" WHERE "CarId" = $1`
	_, err := r.db.Exec(sqlCancelRents, carId)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func (r *RentDb) CreatePrice(carId int, p CarPrices) error {
	query := `INSERT INTO "Price" ("CarId", "%s", "%s", "%s") VALUES ($1, $2, $3, $4);`
	_, err := r.db.Exec(query, carId, p.Hour, p.Day, p.Week)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *RentDb) CreateDate(carId int, d AvailableDate) error {
	SqlCreateDate := `INSERT INTO "Availability" ("CarId", "TimeStart", "TimeEnd") VALUES($1, $2, $3)`
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
	var prices []PriceItem
	for rows.Next() {
		var hour, day, week sql.NullFloat64
		if err := rows.Scan(&hour, &day, &week); err != nil {
			log.Println(err)
			continue
		}

		if hour.Valid {
			prices = append(prices, PriceItem{carId, "hour", hour.Float64})
		}
		if day.Valid {
			prices = append(prices, PriceItem{carId, "day", day.Float64})
		}
		if week.Valid {
			prices = append(prices, PriceItem{carId, "week", week.Float64})
		}
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
		err := rows.Scan(&r.CarId, &r.RenterId, &r.StartTime, &r.EndTime, &r.CalculatedTotal)
		if err != nil {
			log.Println(err)
			continue
		}
		rents = append(rents, r)
	}

	return rents, nil
}

func (r *RentDb) CarRents(id int) ([]Rent, error) {
	SqlCarRents := `SELECT "Id", "RenterId", "CalculatedTotalPrice" FROM "Reservation" WHERE "CarId" = $1;`
	rows, err := r.db.Query(SqlCarRents, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var rents []Rent
	for rows.Next() {
		rent := Rent{}
		if err := rows.Scan(&rent.Id, &rent.RenterId, &rent.CalculatedTotal); err != nil {
			log.Println(err)
			continue
		}
		rents = append(rents, rent)
	}

	return rents, nil
}

func (r *RentDb) CreatePayment(p Payment) (int, error) {
	SqlCreatePayment := `
INSERT INTO "Payment" ("Id", "Amount", "Timestamp", "SenderId", "ReceiverId") VALUES
(DEFAULT, $1, now(), $2, $3) RETURNING "Id";`
	var paymentId int
	err := r.db.QueryRow(SqlCreatePayment, p.Amount, p.SenderId, p.ReceiverId).Scan(&paymentId)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return paymentId, nil
}

func (r *RentDb) Payment(pid int) (Payment, error) {
	payment := Payment{}
	sqlFindPayment := `SELECT "SenderId", "ReceiverId", "Amount" FROM "Payment" WHERE "Id" = $1`
	err := r.db.QueryRow(sqlFindPayment, pid).Scan(&payment.SenderId, &payment.SenderId, &payment.Amount)
	if err != nil {
		log.Println(err)
		return payment, err
	}
	return payment, nil
}
