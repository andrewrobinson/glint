package util

import (
	"testing"
	"time"
)

func TestDateInAugust2020(t *testing.T) {

	t.Run("a mixture of dates either side of August 2020", func(t *testing.T) {

		var date time.Time

		date, _ = time.Parse(time.RFC3339, "2020-07-30T00:00:00Z")
		ShouldNotBeDateInAugust2020(t, date)
		date, _ = time.Parse(time.RFC3339, "2020-07-31T23:59:59Z")
		ShouldNotBeDateInAugust2020(t, date)

		date, _ = time.Parse(time.RFC3339, "2020-08-01T00:00:00Z")
		ShouldBeDateInAugust2020(t, date)

		date, _ = time.Parse(time.RFC3339, "2020-08-01T15:04:05Z")
		ShouldBeDateInAugust2020(t, date)

		date, _ = time.Parse(time.RFC3339, "2020-08-31T23:59:59Z")
		ShouldBeDateInAugust2020(t, date)

		date, _ = time.Parse(time.RFC3339, "2020-09-01T00:00:00Z")
		ShouldNotBeDateInAugust2020(t, date)

	})

}

func ShouldBeDateInAugust2020(t *testing.T, date time.Time) {
	if !dateInAugust2020(date) {
		t.Errorf("DateInAugust2020 should evaluate true for date:%v", date)
	}

}

func ShouldNotBeDateInAugust2020(t *testing.T, date time.Time) {
	if dateInAugust2020(date) {
		t.Errorf("DateInAugust2020 should NOT evaluate true for date:%v", date)
	}
}
