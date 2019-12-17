package calculate

import (
	"errors"
)
// pancake represents a happy-face pancake with choclate chips only the on the "true" side
type pancake bool

// pancakeStack can be any length slice of the custom pancake type (bool). It is read from left (top) to right (bottom).
type pancakeStack []pancake

// ManeuversToGetHappySideUp ranges over input pancake stack and determines the minimal count of maneuvers required for all pancakes to be upright
func ManeuversToGetHappySideUp(input string) (int, error) {
	stack, err := convertStringToPancakeStack(input)
	if err != nil {
		return 0, err
	}
	return flipsForSingleStack(stack, 0), nil
}

// flipsForSingleStack ranges over input pancake stack and determines the count of maneuvers required for all pancakes to be upright
// flipCount should always start at zero when called from outside recursion
func flipsForSingleStack(stack pancakeStack, flipCount int) int {
	// indexToFlip is used to track the max index that needs alteration
	var indexToFlip int

	// foundUpsideDownPancakes is a flag to check that an upside down pancake was found. Avoids ambiguity of indexToFlip since null value is 0, which is a valid index
	foundUpsideDownPancakes := false

	// range through input stack checking for the deepest upside down pancake
	for i, p := range stack {
		if !p { // if a pancake is upside down (false) in the stack, track it's index and keep looking deeper
			foundUpsideDownPancakes = true
			indexToFlip = i
		} else {
			continue
		}
	}

	// if we found pancakes upside down, take the deepest index and flip everything at or above it
	if foundUpsideDownPancakes {
		flipCount++
		rightedPancakes := flip(stack, indexToFlip+1)               // +1 to index indicating count of pancakes to flip, rather that index position
		flipCount = flipsForSingleStack(rightedPancakes, flipCount) // recurse through this func
	}
	return flipCount
}

// flip is used inside calculate to take the top N of a stack and flip them
func flip(input pancakeStack, numberToFlip int) pancakeStack {
	var flippedPancakes pancakeStack

	// for N to be flipped, append the inverse
	for _, pancake := range input[:numberToFlip] {
		flippedPancakes = append(flippedPancakes, !pancake)
	}

	// for remaining pancakes in stack, keep the same
	return append(flippedPancakes, input[numberToFlip:]...)
}


// convertStringToPancakeStack takes the string-form of a stack of pancakes, e.g. "+--+" and converts it to the bool type pancake stack "true, false, false, true"
func convertStringToPancakeStack(input string) (pancakeStack, error) {
	stack := pancakeStack{}
	for _, item := range input {
		if item == '+' {
			stack = append(stack, pancake(true))
		} else if item == '-' {
			stack = append(stack, pancake(false))
		} else {
			return nil, errors.New("Input stack must only include '+' or '-'")
		}
	}
	return stack, nil
}