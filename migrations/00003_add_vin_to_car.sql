-- +goose Up
ALTER TABLE "Car" ADD COLUMN "Vin" VARCHAR(17);
ALTER TABLE "Car" ADD CONSTRAINT unique_car_vin UNIQUE("Vin");

-- +goose Down
ALTER TABLE "Car" DROP CONSTRAINT unique_car_vin;
ALTER TABLE "Car" DROP COLUMN IF EXISTS "Vin";