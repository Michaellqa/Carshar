package service

import (
	"Carshar/dal"
	"errors"
	"log"
)

type CarManager struct {
	db *dal.RentDb
}

func NewCarManager(db *dal.RentDb) *CarManager {
	return &CarManager{db: db}
}

func (m *CarManager) Get(id int) (dal.CarFullDescription, error) {
	return m.db.FindCar(id)
}

func (m *CarManager) AddCar(car dal.Car) (bool, error) {
	return m.db.CreateCar(car)
}

func (m *CarManager) DeleteCar(carId int) error {
	return m.db.DeleteCar(carId)
}

func (m *CarManager) AddPrice(carId int, p dal.PriceItem) error {
	prices, err := m.GetPrices(carId)
	if err != nil {
		log.Println(err)
		return err
	}

	for _, p := range prices {
		if p.TimeUnit == p.TimeUnit {
			return errors.New("already exist")
		}
	}

	item := dal.CarPrices{CarId: carId}
	switch p.TimeUnit {
	case "hour":
		item.Hour = p.Price
	case "day":
		item.Day = p.Price
	case "week":
		item.Week = p.Price
	}
	return m.db.CreatePrice(carId, item)
}

func (m *CarManager) GetPrices(carId int) ([]dal.PriceItem, error) {
	return m.db.CarPrices(carId)
}

func (m *CarManager) DeletePrices(carId int) error {
	return m.db.DeletePricesForCar(carId)
}

func (m *CarManager) AddDates(carId int, dates dal.AvailableDate) error {
	return m.db.CreateDate(carId, dates)
}

func (m *CarManager) DeleteDates(carId int) error {
	return m.db.DeletePricesForCar(carId)
}
