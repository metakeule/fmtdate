package fmtdate

import (

	// "fmt"
	"testing"
	"time"
)

// 2006-01-02T15:04:05Z07:00
var date, _ = time.Parse("2006-01-02 15:04:05", "2007-02-03 16:05:06")

func TestDays(t *testing.T) {
	res := Format("D", date)
	if res != "3" {
		t.Errorf("D should return 3, but returns %#v\n", res)
	}

	res = Format("DD", date)
	if res != "03" {
		t.Errorf("D should return 03, but returns %#v\n", res)
	}

	res = Format("DDD", date)
	if res != "Sat" {
		t.Errorf("DDD should return Sat, but returns %#v\n", res)
	}

	res = Format("DDDD", date)
	if res != "Saturday" {
		t.Errorf("DDDD should return Saturday, but returns %#v\n", res)
	}
}

func TestMonths(t *testing.T) {
	res := Format("M", date)
	if res != "2" {
		t.Errorf("M should return 2, but returns %#v\n", res)
	}

	res = Format("MM", date)
	if res != "02" {
		t.Errorf("MM should return 02, but returns %#v\n", res)
	}

	res = Format("MMM", date)
	if res != "Feb" {
		t.Errorf("MMM should return Feb, but returns %#v\n", res)
	}

	res = Format("MMMM", date)
	if res != "February" {
		t.Errorf("MMMM should return February, but returns %#v\n", res)
	}
}

func TestYears(t *testing.T) {
	res := Format("YY", date)
	if res != "07" {
		t.Errorf("YY should return 07, but returns %#v\n", res)
	}

	res = Format("YYYY", date)
	if res != "2007" {
		t.Errorf("YYYY should return 2007, but returns %#v\n", res)
	}
}

func TestTime(t *testing.T) {
	res := Format("hh", date)
	if res != "16" {
		t.Errorf("hh should return 16, but returns %#v\n", res)
	}

	res = Format("mm", date)
	if res != "05" {
		t.Errorf("mm should return 05, but returns %#v\n", res)
	}

	res = Format("ss", date)
	if res != "06" {
		t.Errorf("ss should return 06, but returns %#v\n", res)
	}
}

func TestDefaults(t *testing.T) {
	now := time.Now()

	tester := func(formatter func(time.Time) string, parser func(string) (time.Time, error)) {
		formatted := formatter(now)
		timed, err := parser(formatted)

		if err != nil {
			t.Errorf("error while parsing %#v: %s", formatted, err.Error())
		}

		if formatter(timed) != formatted {
			t.Errorf("time not correctly retransformatted, was: %#v, is: %#v", formatted, formatter(timed))
		}
	}

	tester(FormatDate, ParseDate)
	tester(FormatTime, ParseTime)
	tester(FormatDateTime, ParseDateTime)
}
