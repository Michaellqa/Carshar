-- +goose Up
CREATE TABLE "User" (
	"Id" SERIAL PRIMARY KEY,
	"Phone" VARCHAR(20) NOT NULL UNIQUE,
	"Password" VARCHAR(50) NOT NULL,
	"Name" VARCHAR(100) NOT NULL,
	"BirthDate" VARCHAR(30) NOT NULL
);

CREATE TABLE "Car" (
	"Id" SERIAL PRIMARY KEY,
	"OwnerId" INTEGER NOT NULL,
	"Model" TEXT NOT NULL,
	"Year" INTEGER NOT NULL,
	"Mileage" INTEGER,
	"Image" TEXT
);

CREATE TABLE "Price" (
	"CarId" INTEGER,
	"TimeUnit" INTEGER, -- 1=minute 2=hour 3=day 4=week
	"Price" DECIMAL
);

CREATE TABLE "Date" (
	"CarId" INTEGER,
	"DayOfWeek" INTEGER, -- 1..7 = Mon..Sun; 0 = any
	"TimeStart" VARCHAR(30),
	"TimeEnd" VARCHAR(30)
);

CREATE TABLE "Rent" (
	"Id" SERIAL PRIMARY KEY,
	"RenterId" INTEGER,
	"CarId" INTEGER,
	"TimeStart" VARCHAR(30),
	"TimeEnd" VARCHAR(30),
	"TotalPrice" DECIMAL
);


-- ALTER TABLE "User"
-- ADD CONSTRAINT UNIQUE

ALTER TABLE "Car"
ADD CONSTRAINT fk_car_user FOREIGN KEY ("OwnerId") REFERENCES "User" ("Id");

ALTER TABLE "Price"
ADD CONSTRAINT fk_price_car FOREIGN KEY ("CarId") REFERENCES "Car" ("Id");

ALTER TABLE "Date"
ADD CONSTRAINT fk_date_car FOREIGN KEY ("CarId") REFERENCES "Car" ("Id");

ALTER TABLE "Rent"
ADD CONSTRAINT fk_rent_car FOREIGN KEY ("CarId") REFERENCES "Car" ("Id");

ALTER TABLE "Rent"
ADD CONSTRAINT fk_rent_user FOREIGN KEY ("RenterId") REFERENCES "User" ("Id");

-- +goose Down
