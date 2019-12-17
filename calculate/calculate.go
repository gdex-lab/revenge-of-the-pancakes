package calculate

import "github.com/gdexlab/revenge-of-the-pancakes/types"

// FlipsForSingleStack ranges over input pancake stack and determines the count of maneuvers required for all pancakes to be upright
// flipCount should always start at zero when called by user
func FlipsForSingleStack(input types.PancakeStack, flipCount int) int {
	// indexToFlip is used to track the max index that needs alteration
	var indexToFlip int

	// foundUpsideDownPancakes is a flag to check that an upside down pancake was found. Avoids ambiguity of indexToFlip since null value is 0, which is a valid index
	foundUpsideDownPancakes := false

	// range through input stack checking for the deepest upside down pancake
	for i, p := range input {
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
		rightedPancakes := flip(input, indexToFlip+1)               // +1 to index indicating count of pancakes to flip, rather that index position
		flipCount = FlipsForSingleStack(rightedPancakes, flipCount) // recurse through this func
	}
	return flipCount
}

// flip is used inside calculate to take the top N of a stack and flip them
func flip(input types.PancakeStack, numberToFlip int) types.PancakeStack {
	var flippedPancakes types.PancakeStack

	// for N to be flipped, append the inverse
	for _, pancake := range input[:numberToFlip] {
		flippedPancakes = append(flippedPancakes, !pancake)
	}

	// for remaining pancakes in stack, keep the same
	return append(flippedPancakes, input[numberToFlip:]...)
}
