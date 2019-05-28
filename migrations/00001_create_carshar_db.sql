-- +goose Up
CREATE TABLE "User" (
	"Id" SERIAL PRIMARY KEY,
	"Name" VARCHAR(100) NOT NULL,
	"PhoneNumber" VARCHAR(20) NOT NULL UNIQUE,
	"Password" VARCHAR(50) NOT NULL,
	"BirthDate" DATE NOT NULL,
	"CreditAmount" DECIMAL not NULL default 0
);

CREATE TABLE "Car" (
	"Id" SERIAL PRIMARY KEY,
	"OwnerId" INTEGER NOT NULL,
	"Model" TEXT NOT NULL,
	"Description" TEXT DEFAULT "",
	"Year" INTEGER NOT NULL,
	"Mileage" INTEGER NOT NULL default 0,
	"Vin" VARCHAR(17) NOT NULL UNIQUE,
	"ImageUrl" TEXT DEFAULT "",
	"Type" VARCHAR(20) DEFAULT "",
	"Transmission" VARCHAR(10) DEFAULT "",
	"LocationId" INTEGER
);

CREATE TABLE "Price" (
	"CarId" INTEGER,
	"Hour" DECIMAL,
	"Day" DECIMAL,
	"Week" DECIMAL
);

CREATE TABLE "Availability" (
  "Id" SERIAL PRIMARY KEY,
	"CarId" INTEGER,
	"TimeStart" TIMESTAMP WITH TIME ZONE NOT NULL,
	"TimeEnd" TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE "Reservation" (
	"Id" SERIAL PRIMARY KEY,
	"RenterId" INTEGER,
	"CarId" INTEGER,
	"StartDate" TIMESTAMP WITH TIME ZONE NOT NULL,
	"EstimatedEndDate" TIMESTAMP WITH TIME ZONE NOT NULL,
	"EndDate" TIMESTAMP WITH TIME ZONE,
	"CalculatedTotalPrice" DECIMAL,
	"PaymentId" INT
);

CREATE TABLE "Payment" (
  "Id" SERIAL PRIMARY Key,
  "Amount" DECIMAL NOT NULL,
  "Timestamp" TIMESTAMP WITH TIME ZONE,
  "SenderId" INTEGER,
  "ReceiverId" INTEGER
);

CREATE TABLE "Location" (
  "Id" SERIAL PRIMARY KEY,
  "CarId" INTEGER NOT NULL,
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
ADD CONSTRAINT fk_rent_transaction FOREIGN KEY ("PaymentId") REFERENCES "Payment" ("Id");

ALTER TABLE "Payment"
ADD CONSTRAINT fk_trans_sender FOREIGN KEY ("SenderId") REFERENCES "User" ("Id");

ALTER TABLE "Payment"
ADD CONSTRAINT fk_trans_receiver FOREIGN KEY ("ReceiverId") REFERENCES "User" ("Id");

ALTER TABLE "Price"
ADD CONSTRAINT fk_price_car FOREIGN KEY ("CarId") REFERENCES "Car" ("Id");

ALTER TABLE "Availability"
ADD CONSTRAINT fk_date_car FOREIGN KEY ("CarId") REFERENCES "Car" ("Id");

ALTER TABLE "Location"
ADD CONSTRAINT fk_location_car FOREIGN KEY ("CarId") REFERENCES "Car" ("Id");


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

--Test data
INSERT INTO "public"."User" ("Id", "Name", "PhoneNumber", "Password", "BirthDate", "CreditAmount") VALUES
(DEFAULT, 'Mike', '8937241422', 'admin', '1997-08-31', 295),
(DEFAULT, 'Fox', '19993332425', '123', '2019-05-27', 1928.30);

INSERT INTO "public"."Car" ("Id", "OwnerId", "Model", "Description", "Year", "Mileage", "Vin", "ImageUrl", "Type", "Transmission", "LocationId") VALUES
(DEFAULT, 1, 'tesla', '-', 2011, 10000, 'QWERTYQWERTY12345', null, null, null, null),
(DEFAULT, 2, 'jaguar f-pace', 'cool', 2018, 2000, 'ABCDE12345FGH0987', null, null, null, null);

INSERT INTO "public"."Availability" ("Id", "CarId", "TimeStart", "TimeEnd") VALUES
(DEFAULT, 1, '2019-05-27 18:57:06.552000', '2019-10-27 18:57:12.987000'),
(DEFAULT, 2, '2019-05-27 18:57:06.552000', '2019-10-27 18:57:12.987000');

INSERT INTO "public"."Location" ("Id", "CarId", "City", "Latitude", "Longitude") VALUES
(DEFAULT, 2, 'Penza', 53.19475, 45.013747);

INSERT INTO "public"."Price" ("CarId", "Hour", "Day", "Week") VALUES
(2, 5, 25, 150);


-- +goose Down
