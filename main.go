package main

import (
	"Carshar/api"
	"Carshar/api/handlers/analytics"
	"Carshar/api/handlers/auth"
	"Carshar/api/handlers/car"
	"Carshar/api/handlers/renting"
	"Carshar/dal"
	"Carshar/service"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"log"
)

const (
	postgres     = "postgres"
	pgConnection = "postgres://localhost/Carshar?sslmode=disable"
)

func main() {
	log.SetFlags(log.Lshortfile)

	doneSignal := make(chan struct{})
	port := 8080

	server := provideServer(port, doneSignal)
	go server.Start()
	fmt.Println("Started on ", port)

	<-doneSignal
}

func provideServer(port int, done chan struct{}) *api.Server {
	pDb := dbConnection()

	goose.SetTableName("Carshar_db_version")
	err := goose.Up(pDb, "migrations/")
	if err != nil {
		panic(err)
	}

	authDb := dal.NewAuthDb(pDb)
	carDb := dal.NewRentDb(pDb)
	analyticsDb := dal.NewAnalyticsDb(pDb)

	carManager := service.NewCarManager(carDb)
	bookingProvider := service.NewBookingProvider(authDb, carDb, carManager)

	authHandler := auth.NewAuthHandler(authDb)
	userHandler := auth.NewCreateUserHandler(authDb)

	addCarHandler := car.NewAddCarHandler(carDb)
	carListHandler := car.NewCarListHandler(carDb)
	findCarHandler := car.NewFindCarHandler(carDb)
	userCarsHandler := car.NewUserCarsHandler(carDb)
	userRentedCarsHandler := car.NewUserRentedCarsHandler(carDb)

	addDateHandler := car.NewAddDateHandler(carDb)
	addPriceHandler := car.NewAddPriceHandler(carManager)

	datesHandler := car.NewDatesHandler(carDb)
	pricesHandler := car.NewPricesHandler(carDb)

	totalHandler := renting.NewTotalHandler(carDb)
	rentHandler := renting.NewRentHandler(carManager, bookingProvider)

	carRentsHandler := renting.NewCarBookingsHandler(carDb)

	analyticsHandler := analytics.NewAnalyticsHandler(analyticsDb)

	mx := api.NewMux(
		userHandler,
		authHandler,
		addCarHandler,
		carListHandler,
		findCarHandler,
		userCarsHandler,
		userRentedCarsHandler,
		addDateHandler,
		addPriceHandler,
		datesHandler,
		pricesHandler,
		totalHandler,
		rentHandler,
		carRentsHandler,
		analyticsHandler,
	)

	return api.NewServer(port, mx, done)
}

func dbConnection() *sql.DB {
	db, err := sql.Open(postgres, pgConnection)
	if err != nil {
		panic(err)
	}
	return db
}
