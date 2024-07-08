// Package weather forecasts the current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition represents the current weather condition in the city.
var CurrentCondition string

// CurrentLocation represents the city of the weather forecast.
var CurrentLocation string

// Forecast returns a string stating the city and the weather condition of the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
