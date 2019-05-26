-- +goose Up
CREATE TABLE "User" (
	"Id" SERIAL PRIMARY KEY,
	"Name" VARCHAR(100) NOT NULL,
	"PhoneNumber" VARCHAR(20) NOT NULL UNIQUE,
	"Password" VARCHAR(50) NOT NULL,
	"BirthDate" VARCHAR(30) NOT NULL,
	"CreditAmount" DECIMAL not NULL default 0
);

CREATE TABLE "Car" (
	"Id" SERIAL PRIMARY KEY,
	"OwnerId" INTEGER NOT NULL,
	"Model" TEXT NOT NULL,
	"Description" TEXT NOT NULL,
	"Year" INTEGER NOT NULL,
	"Mileage" INTEGER NOT NULL default 0,
	"Vin" VARCHAR(17) NOT NULL UNIQUE,
	"ImageUrl" TEXT,
	"Type" VARCHAR(20),
	"Transmission" VARCHAR(10),
	"LocationId" INTEGER
);

CREATE TABLE "Price" (
	"CarId" INTEGER,
	"TimeUnit" VARCHAR(4) NOT NULL UNIQUE,
	"Amount" DECIMAL NOT NULL
);

CREATE TABLE "Availability" (
  "Id" INTEGER PRIMARY KEY,
	"CarId" INTEGER,
	"TimeStart" VARCHAR(30),
	"TimeEnd" VARCHAR(30)
);

CREATE TABLE "Reservation" (
	"Id" SERIAL PRIMARY KEY,
	"RenterId" INTEGER,
	"CarId" INTEGER,
	"StartDateTime" VARCHAR(30),
	"EndDateTime" VARCHAR(30),
	"TotalPrice" DECIMAL,
	"TransactionId" INT
);

CREATE TABLE "Transaction" (
  "Id" SERIAL PRIMARY Key,
  "Timestamp" TIMESTAMP WITH TIME ZONE,
  "SenderId" INTEGER,
  "ReceiverId" INTEGER
);

CREATE TABLE "Location" (
  "Id" SERIAL PRIMARY KEY,
  "City" VARCHAR(40),
  "Latitude" FLOAT,
  "Longitude" FLOAT
);

ALTER TABLE "Car"
ADD CONSTRAINT fk_car_user FOREIGN KEY ("OwnerId") REFERENCES "User" ("Id");

ALTER TABLE "Car"
ADD CONSTRAINT fk_car_location FOREIGN KEY ("LocationId") REFERENCES "Location" ("Id");

ALTER TABLE "Reservation"
ADD CONSTRAINT fk_rent_car FOREIGN KEY ("CarId") REFERENCES "Car" ("Id");

ALTER TABLE "Reservation"
ADD CONSTRAINT fk_rent_user FOREIGN KEY ("RenterId") REFERENCES "User" ("Id");

ALTER TABLE "Reservation"
ADD CONSTRAINT fk_rent_transaction FOREIGN KEY ("TransactionId") REFERENCES "Location" ("Id");

ALTER TABLE "Transaction"
ADD CONSTRAINT fk_trans_sender FOREIGN KEY ("SenderId") REFERENCES "User" ("Id");

ALTER TABLE "Transaction"
ADD CONSTRAINT fk_trans_receiver FOREIGN KEY ("ReceiverId") REFERENCES "User" ("Id");

ALTER TABLE "Price"
ADD CONSTRAINT fk_price_car FOREIGN KEY ("CarId") REFERENCES "Car" ("Id");

ALTER TABLE "Availability"
ADD CONSTRAINT fk_date_car FOREIGN KEY ("CarId") REFERENCES "Car" ("Id");


-- +goose StatementBegin
CREATE  FUNCTION t_check_dates() RETURNS TRIGGER
LANGUAGE plpgsql
AS $func$
DECLARE
	rents_count INT;
BEGIN
	rents_count := (SELECT COUNT(*) FROM "Reservation" WHERE
		("StartDate" BETWEEN NEW."StartDate" AND NEW."EndDate")
		OR ("EndDate" BETWEEN NEW."StartDate" AND NEW."EndDate"));
	IF (rents_count != 0) THEN
		RAISE EXCEPTION 'This period of time conflicts with existing rent';
	END IF;
	RETURN NEW;
END;
$func$
;
-- +goose StatementEnd

CREATE TRIGGER "check_dates_of_rent" BEFORE INSERT ON "Reservation"
FOR EACH ROW EXECUTE PROCEDURE t_check_dates();


-- +goose Down
