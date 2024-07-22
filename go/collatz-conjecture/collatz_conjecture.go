package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("zero or negative value is not allowed")
	}
	steps := 0
	for n > 1 {
		if n % 2 == 0 {
			n /= 2
			steps += 1
		} else {
			return steps, nil
		}
	}
	
	return steps, nil
}
