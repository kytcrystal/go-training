package lasagna

func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

func Quantities(layers []string)(int, float64) {
	noodles := 0
	sauce := 0.0

	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 	50
		}		
		if layer == "sauce" {
			sauce +=0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient() {}
func ScaleRecipe() {}
