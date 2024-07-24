package thefarm

import "errors"

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, nc int) (float64, error) {
	amount, err := fc.FodderAmount(nc)
	if err != nil {
		return 0, err
	}
	factor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}
	totalFood := amount * factor
	return totalFood/float64(nc), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, nc int) (float64, error) {
	if nc <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, nc)
}

// TODO: define the 'ValidateNumberOfCows' function
var errNumberOfCows = errors.New("cows are invalid")

func ValidateNumberOfCows(nc int) error {
	if nc <= 0 {
		return errNumberOfCows
	}
	return nil
}
