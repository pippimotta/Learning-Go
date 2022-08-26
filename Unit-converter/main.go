package main

func main() {
	tool := &tool{
		converters: []*converter{
			celsiusToFahrenheit(), // celsius -> fahrenheit
			fahrenheitToCelsius(), // fahrenheit -> celsius
			calToJoule(),          // calories -> joule
			jouleToCal(),          // joule -> calories
		},
	}

	tool.start()
}
