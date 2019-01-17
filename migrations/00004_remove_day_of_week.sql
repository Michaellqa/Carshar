-- +goose Up
ALTER TABLE "Date" DROP COLUMN "DayOfWeek";

-- +goose Down
