package fmtdate

import (

	// "fmt"
	"testing"
	"time"
)

// 2006-01-02T15:04:05Z07:00
var date, _ = time.Parse("2006-01-02 15:04:05 MST", "2007-02-03 16:05:06 UTC")

func TestFormat(t *testing.T) {

	tests := []struct {
		format   string
		expected string
	}{
		{"hh:mm:ss ZZZZ", "16:05:06 +0000"},
		{"hh:mm:ss ZZZ", "16:05:06 UTC"},
		{"hh:mm:ss ZZ", "16:05:06 Z"},
		{"D", "3"},
		{"DD", "03"},
		{"DDD", "Sat"},
		{"DDDD", "Saturday"},
		{"M", "2"},
		{"MM", "02"},
		{"MMM", "Feb"},
		{"MMMM", "February"},
		{"YY", "07"},
		{"YYYY", "2007"},
		{"hh", "16"},
		{"mm", "05"},
		{"ss", "06"},
		{"hpm", "04PM"},
		{"h:mm:sspm", "04:05:06PM"},
	}

	for _, test := range tests {

		if got, want := Format(test.format, date), test.expected; got != want {
			t.Errorf("Format(%#v,date) = %#v; want %#v", test.format, got, want)
		}
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

func TestTranslate(t *testing.T) {

	goformat := Translate("MM.YYYY.DD hh:mm")
	expected := "01.2006.02 15:04"

	if goformat != expected {
		t.Errorf("format not correctly translated, got: %#v, expected: %#v", goformat, expected)
	}
}
