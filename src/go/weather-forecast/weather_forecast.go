// Package weather provides forecasts for different locations.
package weather

// CurrentCondition the current weather condition.
var CurrentCondition string

// CurrentLocation currently set location.
var CurrentLocation string

// Forecast returns teh forcast for a given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
