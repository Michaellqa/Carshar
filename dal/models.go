package dal

import "time"

type User struct {
	Id        int       `json:"Id"`
	Name      string    `json:"Name"`
	Phone     string    `json:"Phone"`
	Password  string    `json:"Password"`
	BirthDate time.Time `json:"BirthDate"`
}

type Car struct {
	Id      int `json:"Id"`
	OwnerId int
	Model   string
	Year    int
	Mileage int
	Image   string
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

type AvailableDate struct {
	DayOfWeek int
	TimeStart time.Time
	TimeEnd   time.Time
}

type PriceItem struct {
	TimeUnit string //hour, day
	Price    float32
}

type Rent struct {
	CarId     int
	RenterId  int
	TimeStart time.Time
	TimeEnd   time.Time
	Total     float32
}

// list cell representation
type CarShortDescription struct {
	Id    int
	Model string
	Year  int
}

type History struct {
	Model     string
	DateStart time.Time
	Total     float32
}
