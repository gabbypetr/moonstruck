package main

func DetectPhase(

	age float64,

) Phase {

	switch {

	case age < 1.8:

		return Phase{

			"New Moon",

			"🌑",

		}

	case age < 7.4:

		return Phase{

			"Waxing Crescent",

			"🌒",

		}

	case age < 9.5:

		return Phase{

			"First Quarter",

			"🌓",

		}

	case age < 14.8:

		return Phase{

			"Waxing Gibbous",

			"🌔",

		}

	case age < 16.8:

		return Phase{

			"Full Moon",

			"🌕",

		}

	case age < 22.1:

		return Phase{

			"Waning Gibbous",

			"🌖",

		}

	case age < 24.2:

		return Phase{

			"Last Quarter",

			"🌗",

		}

	default:

		return Phase{

			"Waning Crescent",

			"🌘",

		}

	}

}
