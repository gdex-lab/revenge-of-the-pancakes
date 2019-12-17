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
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// pancake represents a happy-face pancake with choclate chips only the on the "true" side
type pancake bool

// pancakeStack can be any length slice of the custom pancake type (bool). It is read from left (top) to right (bottom).
type pancakeStack []pancake

func main() {
	var input []string
	var err error
	numTestCases := 0
	scanCount := 0
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Input number of test cases: ")
	for scanner.Scan() {
		// if number of test cases is 0, we need to get it first
		if numTestCases == 0 {
			numTestCases, err = strconv.Atoi(scanner.Text())
			if err != nil || numTestCases == 0 {
				fmt.Errorf("must enter a valid number greater than 0: %v", err)
				break
			}
			// jump to next iteration
			continue
		}

		// add items to input
		scanCount++
		input = append(input, scanner.Text())

		// if we have reachaed the limit, stop receiving input
		if scanCount >= numTestCases {
			break
		}
	}
	for i, item := range input {
		stack, err := convertInputToPancakeStack(item)
		if err != nil {
			fmt.Println(err)
		}

		countOfManeuvers := calculateFlipsForSingleStack(stack, 0) // pass zero, indicating first iteration
		fmt.Printf("Case #%d: %d\n", i, countOfManeuvers)
	}
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
// flipCount should always start at zero when called by user
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
	for _, pancake := range input[:numberToFlip] {
		flippedPancakes = append(flippedPancakes, !pancake)
	}

	// for remaining pancakes in stack, keep the same
	return append(flippedPancakes, input[numberToFlip:]...)
}
