package main

import "time"

func parseDate(

	text string,

) time.Time {

	t,

		_ := time.Parse(

		"2006-01-02",

		text,

	)

	return t

}
