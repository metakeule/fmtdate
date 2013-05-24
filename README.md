fmtdate
=======

[![Build Status](https://secure.travis-ci.org/metakeule/fmtdate.png)](http://travis-ci.org/metakeule/fmtdate)

a simple go date formatter, based on conventions of Microsoft Excel (TM)

Usage
-----

```go

	package main
	import (
		"github.com/metakeule/fmtdate"
		"fmt"
	)

	func main() {
		date := Format("DD.MM.YYYY", time.Now())
		fmt.Println(date)
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