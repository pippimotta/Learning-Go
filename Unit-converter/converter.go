package main

type converter struct {
	name     string
	calc     func(v float64) float64
	fromUnit string
	toUnit   string
}

// celsius to fahrenheit
func celsiusToFahrenheit() *converter {
	return &converter{
		name: "Celcius[°C] -> Fahrenheit[°F]",
		calc: func(v float64) float64 {
			return v*1.8 + 32
		},
		fromUnit: "°C",
		toUnit:   "°F",
	}
}

// fahrenheit to celsius
func fahrenheitToCelsius() *converter {
	return &converter{
		name: "Fahrenheit[°F] -> Celcius[°C]",
		calc: func(v float64) float64 {
			return (v - 32) / 1.8
		},
		fromUnit: "°F",
		toUnit:   "°C",
	}
}

// calorie to joule
func calToJoule() *converter {
	return &converter{
		name: "Calorie[cal] -> Joule[J]",
		calc: func(v float64) float64 {
			return v * 4.18
		},
		fromUnit: "cal",
		toUnit:   "J",
	}
}

// joule to calorie
func jouleToCal() *converter {
	return &converter{
		name: "Joule[J] -> Calorie[cal]",
		calc: func(v float64) float64 {
			return v * 0.239
		},
		fromUnit: "J",
		toUnit:   "cal",
	}

}
