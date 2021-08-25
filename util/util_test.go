package util

import (
	"testing"
	"time"
)

func TestRunFilters(t *testing.T) {

	t.Run("positive", func(t *testing.T) {

		date, _ := time.Parse(time.RFC822, "01 Aug 2020 10:00 UTC")
		ShouldBeTrue(t, date)

		date, _ := time.Parse(time.RFC822, "01 Aug 2020 10:00 UTC")
		ShouldBeTrue(t, date)

	})

	t.Run("01 July 2020 is not in Aug 2020", func(t *testing.T) {

		date, _ := time.Parse(time.RFC822, "01 Jul 2020 10:00 UTC")
		ShouldBeFalse(t, date)

	})

}

func ShouldBeTrue(t *testing.T, date time.Time) {
	if !DateInAugust2020(date) {
		t.Errorf("DateInAugust2020 should evaluate true for date:%v", date)
	}

}

func ShouldBeFalse(t *testing.T, date time.Time) {
	if DateInAugust2020(date) {
		t.Errorf("DateInAugust2020 should NOT evaluate true for date:%v", date)
	}
}
