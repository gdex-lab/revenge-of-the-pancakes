package main

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

import (
	"fmt"
	"os"

	"github.com/gdexlab/revenge-of-the-pancakes/calculate"
	"github.com/gdexlab/revenge-of-the-pancakes/io"
)

func main() {
	// get user input
	input, err := io.CollectInput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// calculate maneuvers and display results
	fmt.Println("Results:")
	for i, stack := range input {
		countOfManeuvers := calculate.FlipsForSingleStack(stack, 0) // pass zero, indicating first iteration
		fmt.Printf("Case #%d: %d\n", i+1, countOfManeuvers)
	}
}
