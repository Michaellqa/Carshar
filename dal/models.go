package dal

import "time"

const (
	PricePerHour = "hour"
	PricePerDay  = "day"
	PricePerWeek = "week"

	RentStatusActive    = "active"
	RentStatusCancelled = "cancelled"
)

type User struct {
	Id        int       `json:"-"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	BirthDate time.Time `json:"birthDate"`
	Balance   float64   `json:"balance"`
}

type Car struct {
	Id      int    `json:"-"`
	OwnerId int    `json:"ownerId"`
	Model   string `json:"model"`
	Year    int    `json:"year"`
	Mileage int    `json:"mileage"`
	Image   string `json:"image"`
	Vin     string `json:"vin"`
}

type AvailableDate struct {
	CarId     int       `json:"carId"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

type PriceItem struct {
	CarId    int     `json:"carId"`
	TimeUnit string  `json:"timeUnit"`
	Price    float64 `json:"price"`
}

type CarPrices struct {
	CarId int     `json:"carId"`
	Hour  float64 `json:"hour"`
	Day   float64 `json:"day"`
	Week  float64 `json:"week"`
}

type Rent struct {
	Id              int       `json:"-"`
	CarId           int       `json:"carId"`
	RenterId        int       `json:"renterId"`
	StartTime       time.Time `json:"startTime"`
	EndTime         time.Time `json:"endTime"`
	CalculatedTotal float64   `json:"total"`
	PaymentId       int       `json:"payment_id"`
}

type CarFullDescription struct {
	Id      int             `json:"-"`
	OwnerId int             `json:"owner_id"`
	Model   string          `json:"model"`
	Year    int             `json:"year"`
	Mileage int             `json:"mileage"`
	Image   string          `json:"image"`
	Prices  []PriceItem     `json:"-"`
	Dates   []AvailableDate `json:"-"`
}

type CarShortDescription struct {
	Id          int        `json:"id"`
	Model       string     `json:"model"`
	Year        int        `json:"year"`
	Coordinates Coordinate `json:"coordinates"`
}

type CarRentingStatus struct {
	Id     int    `json:"id"`
	Model  string `json:"model"`
	Status int    `json:"status"`
}

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Payment struct {
	Id         int       `json:"id"`
	Amount     float64   `json:"amount"`
	Date       time.Time `json:"date"`
	SenderId   int       `json:"sender_id"`
	ReceiverId int       `json:"receiver_id"`
}
