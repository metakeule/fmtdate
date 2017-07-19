/*
fmtdate provides a date formatter and parser using the syntax of Microsoft Excel (TM).

Additionally it offers default conversions for date time and datetime.

Why?

Microsoft Excel (TM) has a well known syntax for date formatting, that more
memorable than the syntax chosen in the time package in the go library.

Usage

	package main
	import (
		"gopkg.in/metakeule/fmtdate.v1"
		"fmt"
	)

	func main() {
		date := fmtdate.Format("DD.MM.YYYY", time.Now())
		fmt.Println(date)

		var err
		date, err = fmtdate.Parse("M/D/YY", "2/3/07")
		fmt.Println(date, err)
	}
*/
package fmtdate

import (
	"strings"
	"time"
)

/*
	Formats:

	M    - month (1)
	MM   - month (01)
	MMM  - month (Jan)
	MMMM - month (January)
	D    - day (2)
	DD   - day (02)
	DDD  - day (Mon)
	DDDD - day (Monday)
	YY   - year (06)
	YYYY - year (2006)
  hh   - hours (15)
	mm   - minutes (04)
	ss   - seconds (05)

	AM/PM hours: 'h' followed by optional 'mm' and 'ss' followed by 'pm', e.g.

  hpm        - hours (03PM)
  h:mmpm     - hours:minutes (03:04PM)
  h:mm:sspm  - hours:minutes:seconds (03:04:05PM)

  Time zones: a time format followed by 'ZZZZ', 'ZZZ' or 'ZZ', e.g.

  hh:mm:ss ZZZZ (16:05:06 +0100)
  hh:mm:ss ZZZ  (16:05:06 CET)
	hh:mm:ss ZZ   (16:05:06 +01:00)


*/

func replace(in string) (out string) {
	out = in
	for _, ph := range Placeholder {
		out = strings.Replace(out, ph.find, ph.subst, -1)
	}
	return
}

// Format formats a date based on Microsoft Excel (TM) conventions
func Format(format string, date time.Time) string {
	if format == "" {
		format = DefaultDateTimeFormat
	}
	return date.Format(replace(format))
}

// Parse parses a value to a date based on Microsoft Excel (TM) formats
func Parse(format string, value string) (time.Time, error) {
	if format == "" {
		format = DefaultDateTimeFormat
	}
	return time.Parse(replace(format), value)
}

type p struct{ find, subst string }

var Placeholder = []p{
	{"hh", "15"},
	{"h", "03"},
	{"mm", "04"},
	{"ss", "05"},
	{"MMMM", "January"},
	{"MMM", "Jan"},
	{"MM", "01"},
	{"M", "1"},
	{"pm", "PM"},
	{"ZZZZ", "-0700"},
	{"ZZZ", "MST"},
	{"ZZ", "Z07:00"},
	{"YYYY", "2006"},
	{"YY", "06"},
	{"DDDD", "Monday"},
	{"DDD", "Mon"},
	{"DD", "02"},
	{"D", "2"},
}

var (
	DefaultTimeFormat     = "hh:mm:ss"
	DefaultDateFormat     = "YYYY-MM-DD"
	DefaultDateTimeFormat = "YYYY-MM-DD hh:mm:ss"
)

// FormatDate formats the given date to the DefaultDateFormat
func FormatDate(date time.Time) string {
	return Format(DefaultDateFormat, date)
}

// FormatTime formats the given date to the DefaultTimeFormat
func FormatTime(date time.Time) string {
	return Format(DefaultTimeFormat, date)
}

// FormatTime formats the given date to the DefaultDateTimeFormat
func FormatDateTime(date time.Time) string {
	return Format(DefaultDateTimeFormat, date)
}

// Parse parses a date in DefaultDateFormat to a date
func ParseDate(value string) (time.Time, error) {
	return Parse(DefaultDateFormat, value)
}

// Parse parses a date in DefaultTimeFormat to a date
func ParseTime(value string) (time.Time, error) {
	return Parse(DefaultTimeFormat, value)
}

// Parse parses a date in DefaultDateTimeFormat to a date
func ParseDateTime(value string) (time.Time, error) {
	return Parse(DefaultDateTimeFormat, value)
}
