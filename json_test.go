package fmtdate

import (

	// "fmt"
	"testing"
	// "time"
)

func mustTimeDate(format, date string) TimeDate {
	td, err := NewTimeDate(format, date)
	if err != nil {
		panic(err.Error())
	}

	return td
}

func TestMashalError(t *testing.T) {

	td, err := NewTimeDate("hh:mm:ss ZZZZ", "16:05,06 +0070")

	if err == nil {
		t.Errorf("expected error while parsing the date, but got: %s", td.String())
	}
}

func TestMashal(t *testing.T) {

	tests := []struct {
		format   string
		date     string
		expected string
	}{
		{"hh:mm:ss ZZZZ", "16:05:06 +0000", "16:05:06 +0000"},
		{"hh:mm:ss ZZZZ", "", "null"},
		{"", "2006-01-02 15:04:05", "2006-01-02 15:04:05"},
	}

	for _, test := range tests {
		var td TimeDate
		var err error
		if test.date != "" {
			td, err = NewTimeDate(test.format, test.date)

			if err != nil {
				t.Fatalf("NewTimeDate(%#v, %#v); returned error: %v", test.format, test.date, err)
			}

		}

		b, err := td.MarshalJSON()

		if err != nil {
			t.Fatalf("(%#v,%#v).MarshalJSON(); returned error: %v", test.format, test.date, err)
		}

		if got, want := string(b), test.expected; got != want {
			t.Errorf("(%#v,%#v).MarshalJSON() = %#v; want %#v", test.format, test.date, got, want)
		}
	}

}

func TestUnMashal(t *testing.T) {

	tests := []struct {
		format   string
		date     string
		expected string
	}{
		{"hh:mm:ss ZZZZ", "16:05:06 +0000", "16:05:06 +0000"},
		{"hh:mm:ss ZZZZ", "null", "null"},
		{"", "2006-01-02 15:04:05", "2006-01-02 15:04:05"},
	}

	for _, test := range tests {

		var td TimeDate
		td.Format = test.format

		err := td.UnmarshalJSON([]byte(test.date))

		if err != nil {
			t.Fatalf("(%#v,%#v).UnmarshalJSON(); returned error: %v", test.format, test.date, err)
		}

		if td.IsNil() {
			if got, want := "null", test.expected; got != want {
				t.Errorf("(%#v,%#v).UnmarshalJSON() = %#v; want %#v", test.format, test.date, got, want)
			}

			// t.Fatalf("(%#v,%#v).UnmarshalJSON(); returned nil", test.format, test.date)
		} else {

			if got, want := Format(test.format, *td.Time), test.expected; got != want {
				t.Errorf("(%#v,%#v).UnmarshalJSON() = %#v; want %#v", test.format, test.date, got, want)
			}
		}
	}

}
