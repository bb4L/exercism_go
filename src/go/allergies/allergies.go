package allergies

import (
	"math"
)

var allergenes = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

var iterationOrder = []string{
	"cats", "pollen", "chocolate", "tomatoes", "strawberries", "shellfish", "peanuts", "eggs",
}

func Allergies(allergies uint) []string {
	result := []string{}

	if allergies > 255 {
		k := uint(math.Pow(2, math.Floor(math.Log2(float64((allergies))))))

		for ; k > 128 && allergies >= k; k /= 2 {
			allergies -= k
		}
	}

	for _, allergen := range iterationOrder {
		allergenValue := allergenes[allergen]

		if allergies >= allergenValue {
			allergies -= allergenValue
			result = append(result, allergen)
		}
	}
	return result
}

func AllergicTo(allergies uint, allergen string) bool {
	for _, val := range Allergies(allergies) {
		if val == allergen {
			return true
		}
	}
	return false
}
