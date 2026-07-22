package main

import "fmt"

func Print(

	date string,

	age float64,

	phase Phase,

) {

	fmt.Println(

		"Moon Phase Visualizer\n",

	)

	fmt.Println(

		"Date\n",

	)

	fmt.Println(

		date,

	)

	fmt.Println()

	fmt.Println(

		"Moon Age\n",

	)

	fmt.Printf(

		"%.1f days\n\n",

		age,

	)

	fmt.Println(

		"Phase\n",

	)

	fmt.Println(

		phase.Name,

	)

	fmt.Println()

	fmt.Println(

		"Visualization\n",

	)

	fmt.Println(

		phase.Icon,

	)

}
