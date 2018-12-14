-- +goose Up

-- +goose StatementBegin
CREATE  FUNCTION t_check_dates() RETURNS TRIGGER
LANGUAGE plpgsql
AS $func$
DECLARE
	rents_count INT;
BEGIN
	rents_count := (SELECT COUNT(*) FROM "Rent" WHERE
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

CREATE TRIGGER "check_dates_of_rent" BEFORE INSERT ON "Rent"
FOR EACH ROW EXECUTE PROCEDURE t_check_dates();

-- +goose Down
DROP TRIGGER IF EXISTS "tg_check_dates";
DROP FUNCTION IF EXISTS t_check_dates;