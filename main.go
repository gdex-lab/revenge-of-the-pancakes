// **************************************************
// REVENGE OF THE PANCAKES
// Take home challenge for Weave
//
// Author: Grant Dexter - dexter.d.grant@gmail.com
// Applicant for Data Engineer Position
// **************************************************
package main

import (
	"errors"
	"fmt"
)

type pancake bool

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
func calculateFlipsForSingleStack(input pancakeStack) (int, error) {
	for _, p := range input {
		if p {
			continue
		} else {
			continue
		}
	}
	return 1, nil
}
