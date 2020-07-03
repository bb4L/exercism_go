package space

// Planet name of the planet
type Planet string

// YearInSeconds is the amount of seconds that equal a year
const YearInSeconds = 31557600

// ConversionConstants constants to convert to different planets
var ConversionConstants = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age computes the age on the specified planet
func Age(seconds float64, planet Planet) float64 {
	return seconds / (YearInSeconds * ConversionConstants[planet])
}
