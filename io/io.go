package io

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CollectInput scans the cli for user input and returns the slice of pancake stacks for calculation
func CollectInput() ([]string, error) {

	var input []string
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
		stack := scanner.Text()
		err = validateInput(stack)
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
func validateInput(input string) error {
	if strings.ReplaceAll(strings.ReplaceAll(input, "+", ""), "-", "") != "" {
		return errors.New("Input stack must only include '+' or '-'")
	}
	return nil
}
