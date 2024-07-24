package thefarm

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
	return 0.0, nil
}
// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(nc int) error {
	return nil
}
