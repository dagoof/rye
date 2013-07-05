package rye

import (
	"testing"
	"time"
)


func TestDays(t *testing.T) {
	if MakeDuration(time.Hour * time.Duration(23)).Days() != 0 {
		t.Fatal("23 hours should be 0 days")
	}
	if MakeDuration(time.Hour * time.Duration(24)).Days() != 1 {
		t.Fatal("24 hours should be 1 days")
	}
	if MakeDuration(time.Hour * time.Duration(30)).Days() != 1 {
		t.Fatal("30 hours should be 1 days")
	}
	if MakeDuration(time.Hour * time.Duration(56)).Days() != 2 {
		t.Fatal("42 hours should be 2 days")
	}

}

func TestWeeks(t *testing.T) {
	if MakeDuration(Day * time.Duration(6)).Weeks() != 0 {
		t.Fatal("6 days should be 0 weeks")
	}
	if MakeDuration(Day * time.Duration(14)).Weeks() != 2 {
		t.Fatal("14 days should be 2 weeks")
	}
	if MakeDuration(Day * time.Duration(23)).Weeks() != 3 {
		t.Fatal("23 days should be 3 weeks")
	}

}

func TestMonths(t *testing.T) {
	if MakeDuration(Week * time.Duration(1)).Months() != 0 {
		t.Fatal("1 weeks should be 0 months")
	}
	if MakeDuration(Week * time.Duration(9)).Months() != 2 {
		t.Fatal("9 weeks should be 2 months")
	}
}

func TestYears(t *testing.T) {
	if MakeDuration(Month * time.Duration(5)).Years() != 0 {
		t.Fatal("5 months should be 0 years")
	}
	if MakeDuration(Month * time.Duration(12)).Years() != 1 {
		t.Fatal("12 months should be 1 years")
	}
	if MakeDuration(Month * time.Duration(15)).Years() != 1 {
		t.Fatal("15 months should be 1 years")
	}
	if MakeDuration(Month * time.Duration(25)).Years() != 2 {
		t.Fatal("25 months should be 2 years")
	}

}

func TestRemainingHours(t *testing.T) {
	if MakeDuration(time.Hour * time.Duration(36)).RemainingHours() != 12 {
		t.Fatal("36 Remaining hours should be 12")
	}
}

