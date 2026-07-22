package main

func main() {

	repository :=

		Repository{}

	dates :=

		repository.Dates()

	calculator :=

		Calculator{}

	date :=

		dates[2]

	age :=

		calculator.Age(

			date,

		)

	phase :=

		DetectPhase(

			age,

		)

	Print(

		date,

		age,

		phase,

	)

	PrintStats(

		len(dates),

	)

}
