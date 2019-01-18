package dal

import "time"

type User struct {
	Id        int       `json:"-"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	BirthDate time.Time `json:"birthDate"`
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

//TimeUnit possible values:
// 1 = hour,
// 2 = day
type PriceItem struct {
	CarId    int     `json:"carId"`
	TimeUnit string  `json:"timeUnit"`
	Price    float32 `json:"price"`
}

type Rent struct {
	Id        int       `json:"-"`
	CarId     int       `json:"carId"`
	RenterId  int       `json:"renterId"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Total     float32   `json:"total"`
}

type CarFullDescription struct {
	Id      int
	Model   string
	Year    int
	Mileage int
	Image   string
	Prices  []PriceItem
	Dates   []AvailableDate
}

// list cell representation
type CarShortDescription struct {
	Id    int    `json:"id"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

type CarRentingStatus struct {
	Id     int    `json:"id"`
	Model  string `json:"model"`
	Status int    `json:"status"`
}
