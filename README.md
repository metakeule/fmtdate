fmtdate
=======

[![Build Status](https://secure.travis-ci.org/metakeule/fmtdate.png)](http://travis-ci.org/metakeule/fmtdate)

`fmtdate` provides a date formatter and parser using the syntax of Microsoft Excel (TM). 

Additionally it offers default conversions for date time and datetime.

Why?
----

Microsoft Excel (TM) has a well known syntax for date formatting, that more 
memorable than the syntax chosen in the time package in the go library.

Usage
-----

```go

	package main
	import (
		"github.com/metakeule/fmtdate"
		"fmt"
	)

	func main() {
		date := fmtdate.Format("DD.MM.YYYY", time.Now())
		fmt.Println(date)

		var err
		date, err = fmtdate.Parse("M/D/YY", "2/3/07")
		fmt.Println(date, err)
	}

```

Placeholders
------------

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

Documentation
-------------

see http://godoc.org/github.com/metakeule/fmtdate