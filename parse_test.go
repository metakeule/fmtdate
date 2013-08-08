package fmtdate

import (
	"testing"
)

var tests = map[string]string{
	"D":    "3",
	"DD":   "03",
	"DDD":  "Sat",
	"DDDD": "Saturday",
	"M":    "2",
	"MM":   "02",
	"MMM":  "Feb",
	"MMMM": "February",
	"YY":   "07",
	"YYYY": "2007",
	"hh":   "16",
	"mm":   "05",
	"ss":   "06",
}

func TestParse(t *testing.T) {
	for format, value := range tests {
		date, err := Parse(format, value)
		if err != nil {
			t.Fatalf("error on parsing %#v: %s", format, err.Error())
		}
		// we rely on the working tests for Format here ;-)
		if Format(format, date) != value {
			t.Errorf("parsing %#v should return %#v, but returns %#v\n", format, value, Format(format, date))
		}
	}
}

func TestInvalid(t *testing.T) {
	_, err := Parse("x", "Sat")
	if err == nil {
		t.Errorf(`parsing "x" should return error, but does not`)
	}
}
