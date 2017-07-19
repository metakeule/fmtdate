fmtdate
=======

[![Build Status](https://secure.travis-ci.org/metakeule/fmtdate.png)](http://travis-ci.org/metakeule/fmtdate) [![Total views](https://sourcegraph.com/api/repos/github.com/metakeule/fmtdate/counters/views.png)](https://sourcegraph.com/github.com/metakeule/fmtdate)

100% test coverage (that was easy :-))

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

For json

```go

    package main

    import (
        "github.com/metakeule/fmtdate"
        "fmt"
        "encoding/json"
    )

    type Person struct {
        Name string
        BirthDay fmtdate.TimeDate
    }

    func main() {
        bday, err := fmtdate.NewTimeDate("YYYY-MM-DD", "2000-12-04")
        // do error handling
        paul := &Person{"Paul", bday}

        data, err := json.Marshal(paul)
        // do error handling
    }
```

Placeholders
------------

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
  

Documentation
-------------

see http://godoc.org/github.com/metakeule/fmtdate