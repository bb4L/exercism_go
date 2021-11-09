package lasagna

func PreparationTime(layers []string, avgTime int) int {
	if avgTime == 0 {
		avgTime = 2
	}
	return len(layers) * avgTime
}

const noodlesPerLayer = 50
const saucePerLayer = 0.2

func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		if layer == "sauce" {
			sauce += saucePerLayer
		}
		if layer == "noodles" {
			noodles += noodlesPerLayer
		}
	}
	return
}

func AddSecretIngredient(friendsList []string, myList []string) []string {
	return append(myList, friendsList[len(friendsList)-1])
}

func ScaleRecipe(quantities []float64, portiones int) (result []float64) {
	factor := float64(portiones) / 2
	for _, quantity := range quantities {
		result = append(result, quantity*factor)
	}
	return
}
