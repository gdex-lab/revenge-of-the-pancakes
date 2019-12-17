// **************************************************
// REVENGE OF THE PANCAKES
// Take home challenge for Weave
//
// Author: Grant Dexter - dexter.d.grant@gmail.com
// Applicant for Data Engineer Position
//
// Summary: flip each stack of pancakes to be completely
// happy-face side up in the least number of iterations,
// with clean, maintainable code.
// **************************************************
package main

import (
	"errors"
	"fmt"
)

// pancake represents a happy-face pancake with choclate chips only the on the "true" side
type pancake bool

// pancakeStack can be any length slice of the custom pancake type (bool). It is read from left (top) to right (bottom).
type pancakeStack []pancake

func main() {
	// take user input
	fmt.Println("Calculating ...")

}
func convertInputToPancakeStack(input string) (pancakeStack, error) {
	var stack pancakeStack
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

// calculateFlipsForSingleStack ranges over input pancake stack and determines the count of maneuvers required for all pancakes to be upright
func calculateFlipsForSingleStack(input pancakeStack, flipCount int) int {
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
		rightedPancakes := flip(input, indexToFlip+1)                        // +1 to index indicating count of pancakes to flip, rather that index position
		flipCount = calculateFlipsForSingleStack(rightedPancakes, flipCount) // recurse through this func
	}
	return flipCount
}

// flip is used in calculation to take the top N of a stack and flip them
func flip(input pancakeStack, numberToFlip int) pancakeStack {
	var flippedPancakes pancakeStack

	// for N to be flipped, append the inverse
	for _, p := range input[:numberToFlip] {
		flippedPancakes = append(flippedPancakes, !p)
	}

	// for remaining pancakes in stack, keep the same
	return append(flippedPancakes, input[numberToFlip:]...)
}
