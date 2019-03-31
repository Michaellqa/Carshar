package dal

import (
	"database/sql"
	"github.com/lib/pq"
	"log"
	"time"
)

type AnalyticsDb struct {
	db *sql.DB
}

const (
	SqlNumberOfRents = `
SELECT COUNT(*) FROM "Rent" 
WHERE "CarId" = $1 
AND ("StartDate" BETWEEN $2 AND $3);
`
	SqlAverageRentTime = `
SELECT avg("EndDate" - "StartDate") as "avg"
FROM "Rent" 
WHERE "CarId" = $1
AND "StartDate" BETWEEN $2 AND $3;
`
	SqlPercentRentTime = `
WITH 
"CarDates" as (
SELECT * FROM "Rent" WHERE "CarId" = $1
),
"before" as (
SELECT SUM("EndDate" - $2) as "sum" 
FROM "CarDates"
WHERE ("StartDate" < $2)
AND ("EndDate" BETWEEN $2 AND $3)
), 
"during" as (
SELECT SUM("EndDate" - "StartDate") as "sum" 
FROM "CarDates"
WHERE ("StartDate" BETWEEN $2 AND $3)
AND ("EndDate" BETWEEN $2 AND $3)
), 
"current" as (
SELECT SUM($3 - "StartDate") as "sum" 
FROM "CarDates"
WHERE ("StartDate" BETWEEN $2 AND $3)
AND ("EndDate" > $3)
),
"total" as (
SELECT SUM("sum") as "t" FROM (
	SELECT "sum" FROM "before" 
	UNION  SELECT "sum" FROM "during"
	UNION  SELECT "sum" FROM "current"
) as "res"
)
SELECT EXTRACT(epoch FROM "t") / EXTRACT(epoch FROM $3 - $2) as "percent" 
FROM total;
`
	SqlTimeAvailableForRent = `
WITH 
"CarDates" as (
SELECT * FROM "Date" WHERE "CarId" = $1
),
"before" as (
SELECT SUM("EndTime" - $2) as "sum" 
FROM "CarDates"
WHERE ("StartTime" < $2)
AND ("EndTime" BETWEEN $2 AND $3)
), 
"during" as (
SELECT SUM("EndTime" - "StartTime") as "sum" 
FROM "CarDates"
WHERE ("StartTime" BETWEEN $2 AND $3)
AND ("EndTime" BETWEEN $2 AND $3)
), 
"current" as (
SELECT SUM($3 - "StartTime") as "sum" 
FROM "CarDates"
WHERE ("StartTime" BETWEEN $2 AND $3)
AND ("EndTime" > $3)
)
SELECT SUM("sum") as "ssum" FROM (
	SELECT "sum" FROM "before" 
	UNION  SELECT "sum" FROM "during"
	UNION  SELECT "sum" FROM "current"
) as "res";
`
	SqlAverageRenterAge = `
SELECT avg(age("BirthDate"))  as "avg_age"
FROM "Rent" INNER JOIN "User"
ON "User"."Id" = "Rent"."RenterId"
WHERE "CarId" = $1;
`
)

func NewAnalyticsDb(db *sql.DB) *AnalyticsDb {
	return &AnalyticsDb{db: db}
}

func (a *AnalyticsDb) NumberOfRents(carId int, start, end time.Time) int {
	rents := 0
	if err := a.db.QueryRow(SqlNumberOfRents, carId, start, end).Scan(&rents); err != nil {
		log.Println(err)
		return 0
	}
	return rents
}

func (a *AnalyticsDb) AverageRentTime(carId int, start, end time.Time) time.Time {
	t := pq.NullTime{}
	if err := a.db.QueryRow(SqlAverageRentTime, carId, start, end).Scan(&t); err != nil {
		log.Println(err)
		return t.Time
	}
	log.Println(carId, start, end, t)
	return t.Time
}

func (a *AnalyticsDb) AvailableForRent(carId int, start, end time.Time) time.Time {
	t := pq.NullTime{}
	if err := a.db.QueryRow(SqlTimeAvailableForRent, carId, start, end).Scan(&t); err != nil {
		log.Println(err)
		return t.Time
	}
	return t.Time
}

func (a *AnalyticsDb) PercentRentTime(carId int, start, end time.Time) float64 {
	p := sql.NullFloat64{}
	if err := a.db.QueryRow(SqlPercentRentTime, carId, start, end).Scan(&p); err != nil {
		log.Println(err)
		return 0
	}
	return p.Float64
}

func (a *AnalyticsDb) AverageRenterAge(carId int, start, end time.Time) time.Time {
	t := pq.NullTime{}
	if err := a.db.QueryRow(SqlAverageRenterAge, carId).Scan(&t); err != nil {
		log.Println(err)
		return t.Time
	}
	return t.Time
}
