package thefarm

import (
	"errors"
	"fmt"
)

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
	return totalFood / float64(nc), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, nc int) (float64, error) {
	if nc <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, nc)
}

// TODO: define the 'ValidateNumberOfCows' function
type NumberOfCowsError struct {
	number string
	details string
}

func (e *NumberOfCowsError) Error() string {
	return fmt.Sprintf("%v cows are invalid: %s", e.number, e.details)
}

func ValidateNumberOfCows(nc int) error {
	if nc < 0 {
		return &NumberOfCowsError{
			number: fmt.Sprintf("%v",nc),
			details: "there are no negative cows",
		}	
	}
	if nc == 0 {
		return &NumberOfCowsError{
			number: fmt.Sprintf("%v",nc),
			details: "no cows don't need food",
		}	
	}
	return nil
}
