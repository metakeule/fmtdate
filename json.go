package fmtdate

import (
	"time"
)

func NewTimeDate(format, value string) (td TimeDate, err error) {
	td.Format = format
	var t time.Time
	t, err = Parse(format, value)
	if err != nil {
		return
	}
	td.Time = &t
	return
}

type TimeDate struct {
	Format string
	*time.Time
}

func (t TimeDate) MarshalJSON() ([]byte, error) {
	if t.Time == nil {
		return []byte("null"), nil
	}
	return []byte(Format(t.Format, *t.Time)), nil
}

func (t TimeDate) IsNil() bool {
	return t.Time == nil
}

func (t *TimeDate) UnmarshalJSON(data []byte) error {
	s := string(data)

	if s == "null" {
		t.Time = nil
		return nil
	}

	td, err := Parse(t.Format, s)

	if err == nil {
		t.Time = &td
	}

	return err
}
