package main

import "math"

type Calculator struct{}

func (Calculator) Age(

	date string,

) float64 {

	base := parseDate(

		"2000-01-06",

	)

	target := parseDate(

		date,

	)

	days :=

		target.Sub(base).Hours() /

			24

	age := math.Mod(

		days,

		SynodicMonth,

	)

	if age < 0 {

		age += SynodicMonth

	}

	return age

}
