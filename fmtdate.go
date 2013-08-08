package fmtdate

import (
	"strings"
	"time"
)

/*
	Formats a date based on Microsoft Excel (TM) conventions:
	hh   - hours
	mm   - minutes
	ss   - seconds
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
*/
func replace(in string) (out string) {
	out = in
	for _, ph := range Placeholder {
		out = strings.Replace(out, ph.find, ph.subst, -1)
	}
	return
}

func Format(format string, date time.Time) string {
	return date.Format(replace(format))
}

func Parse(format string, value string) (time.Time, error) {
	return time.Parse(replace(format), value)
}

type p struct{ find, subst string }

var Placeholder = []p{
	p{"hh", "15"},
	p{"mm", "04"},
	p{"ss", "05"},
	p{"MMMM", "January"},
	p{"MMM", "Jan"},
	p{"MM", "01"},
	p{"M", "1"},
	p{"YYYY", "2006"},
	p{"YY", "06"},
	p{"DDDD", "Monday"},
	p{"DDD", "Mon"},
	p{"DD", "02"},
	p{"D", "2"},
}
