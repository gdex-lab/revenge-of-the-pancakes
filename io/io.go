package io

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/gdexlab/revenge-of-the-pancakes/types"
)

// CollectInput scans the cli for user input and returns the slice of pancake stacks for calculation
func CollectInput() ([]types.PancakeStack, error) {

	var input []types.PancakeStack
	var stack types.PancakeStack
	var err error
	numTestCases := 0
	scanCount := 0
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Input number of test cases: ")
	for scanner.Scan() {

		// if number of test cases is 0, we need to get test cases count first
		if numTestCases == 0 {
			numTestCases, err = strconv.Atoi(scanner.Text())
			if err != nil || numTestCases == 0 {
				return nil, fmt.Errorf("must enter a valid number greater than 0: %v", err)
			}
			continue
		}

		// add items to input
		scanCount++
		stack, err = convertInputToPancakeStack(scanner.Text())
		if err != nil {
			return nil, err
		}
		input = append(input, stack)

		// if we have reached the limit, stop receiving input
		if scanCount >= numTestCases {
			break
		}

	}
	return input, err
}

// convertInputToPancakeStack takes a string of +, - and converts it to a pancake stack of bools
func convertInputToPancakeStack(input string) (types.PancakeStack, error) {
	var stack types.PancakeStack
	for _, item := range input {
		if item == '+' {
			stack = append(stack, types.Pancake(true))
		} else if item == '-' {
			stack = append(stack, types.Pancake(false))
		} else {
			return nil, errors.New("Input stack must only include '+' or '-'")
		}
	}
	return stack, nil
}
