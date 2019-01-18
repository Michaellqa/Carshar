-- +goose Up

ALTER TABLE "Date" ADD COLUMN "StartTimeT" TIMESTAMP with time ZONE NULL;
ALTER TABLE "Date" ADD COLUMN "EndTimeT" TIMESTAMP with time ZONE NULL;

UPDATE "Date" Set "StartTimeT" = "StartTime"::TIMESTAMP with time ZONE;
UPDATE "Date" Set "EndTimeT" = "EndTime"::TIMESTAMP with time ZONE;

ALTER TABLE "Date" ALTER COLUMN "StartTime" TYPE TIMESTAMP WITH TIME ZONE USING "StartTimeT";
ALTER TABLE "Date" ALTER COLUMN "EndTime" TYPE TIMESTAMP WITH TIME ZONE USING "EndTimeT";

ALTER TABLE "Date" DROP COLUMN "StartTimeT";
ALTER TABLE "Date" DROP COLUMN "EndTimeT";


-- +goose Down
