-- +goose Up
ALTER table "Car" add column "latitude" float;
ALTER table "Car" add column "longitude" float;

-- +goose Down
