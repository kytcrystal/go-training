package raindrops

import (
	"fmt"
	"sort"
)

func Convert(number int) string {
	result := ""
	sounds, keys := prepareSoundMap()

	for _, raindrop := range keys {
		if number%raindrop == 0 {
			result += sounds[raindrop]
		}
	}
	if result != "" {
		return result
	}

	return fmt.Sprint(number)
}

func prepareSoundMap() (map[int]string, []int) {
	sounds := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}

	keys := make([]int, 0, len(sounds))
	for k := range sounds {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	return sounds, keys
}
