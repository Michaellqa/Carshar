package dal

import "time"

type User struct {
	Id        int       `json:"-"`
	Name      string    `json:"Name"`
	Phone     string    `json:"Phone"`
	Password  string    `json:"Password"`
	BirthDate time.Time `json:"BirthDate"`
}

type Car struct {
	Id      int    `json:"-"`
	OwnerId int    `json:"OwnerId"`
	Model   string `json:"Model"`
	Year    int    `json:"Year"`
	Mileage int    `json:"Mileage"`
	Image   string `json:"Image"`
}

type AvailableDate struct {
	CarId     int       `json:"CarId"`
	DayOfWeek int       `json:"DayOfWeek"`
	StartTime time.Time `json:"StartTime"`
	EndTime   time.Time `json:"EndTime"`
}

//TimeUnit possible values:
// 1 = hour,
// 2 = day
type PriceItem struct {
	CarId    int     `json:"CarId"`
	TimeUnit string  `json:"TimeUnit"`
	Price    float32 `json:"Price"`
}

type Rent struct {
	Id        int       `json:"-"`
	CarId     int       `json:"CarId"`
	RenterId  int       `json:"RenterId"`
	StartTime time.Time `json:"StartTime"`
	EndTime   time.Time `json:"EndTime"`
	Total     float32   `json:"Total"`
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
	Id    int
	Model string
	Year  int
}
