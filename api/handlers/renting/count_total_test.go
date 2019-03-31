package renting

import (
	"Carshar/dal"
	"testing"
	"time"
)

func TestCountTotal(t *testing.T) {
	t1, _ := time.Parse("15:04 02-01-2006", "14:00 01-03-2019")
	t2, _ := time.Parse("15:04 02-01-2006", "18:00 09-03-2019")
	t3, _ := time.Parse("15:04 02-01-2006", "18:00 11-03-2019")

	cases := []struct {
		start, end time.Time
		units      []dal.PriceItem
		total      Total
	}{
		{
			start: t1,
			end:   t2,
			units: []dal.PriceItem{
				{Price: 3, TimeUnit: dal.PricePerHour},
				{Price: 15, TimeUnit: dal.PricePerDay},
				{Price: 42, TimeUnit: dal.PricePerWeek},
			},
			total: Total{Value: 69, Parts: []Part{
				{Unit: dal.PricePerWeek, Count: 1, Base: 42},
				{Unit: dal.PricePerDay, Count: 1, Base: 15},
				{Unit: dal.PricePerHour, Count: 4, Base: 3},
			}},
		},
		{
			start: t1,
			end:   t3,
			units: []dal.PriceItem{
				{Price: 3, TimeUnit: dal.PricePerHour},
				{Price: 15, TimeUnit: dal.PricePerDay},
				{Price: 42, TimeUnit: dal.PricePerWeek},
			},
			total: Total{Value: 84, Parts: []Part{
				{Unit: dal.PricePerWeek, Count: 2, Base: 42},
			}},
		},
	}

	for _, c := range cases {
		res := findMinPrice(c.start, c.end, c.units)

		if res.Value != c.total.Value {
			t.Error("Total:", res.Value, "!=", c.total.Value)
		}

		if len(res.Parts) != len(c.total.Parts) {
			t.Error("Parts count:", len(res.Parts), "!=", len(c.total.Parts), res.Parts)
			return
		}

		for i := range res.Parts {
			if res.Parts[i] != c.total.Parts[i] {
				t.Error("Part #", i, res.Parts[i], "!=", c.total.Parts[i])
			}
		}
	}
}
