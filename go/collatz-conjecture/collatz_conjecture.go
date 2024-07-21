package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("zero or negative value is not allowed")
	}
	return 0, nil
}
