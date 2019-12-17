package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCalculateFlipsForSingleStack(t *testing.T) {
	Convey("Given example stacks as input", t, func() {
		stack1 := pancakeStack{false}
		stack2 := pancakeStack{false, true}
		stack3 := pancakeStack{true, false}
		stack4 := pancakeStack{true, true, true}
		stack5 := pancakeStack{false, false, true, false}

		Convey("When the calculate function is called", func() {
			res1 := calculateFlipsForSingleStack(stack1, 0)
			res2 := calculateFlipsForSingleStack(stack2, 0)
			res3 := calculateFlipsForSingleStack(stack3, 0)
			res4 := calculateFlipsForSingleStack(stack4, 0)
			res5 := calculateFlipsForSingleStack(stack5, 0)

			Convey("The output should match sample output", func() {
				So(res1, ShouldEqual, 1)
				So(res2, ShouldEqual, 1)
				So(res3, ShouldEqual, 2)
				So(res4, ShouldEqual, 0)
				So(res5, ShouldEqual, 3)

			})
		})
	})
}

func TestConvertInputToPancakeStack(t *testing.T) {
	Convey("Given input string", t, func() {
		validInput := "+--+"
		validOutput := pancakeStack{true, false, false, true}
		invalidInput := "+XX+"

		Convey("When the conversion function is called", func() {
			out, err := convertInputToPancakeStack(validInput)
			_, err2 := convertInputToPancakeStack(invalidInput)
			Convey("The output should match examples", func() {
				So(err, ShouldBeNil)
				So(out, ShouldResemble, validOutput)

				So(err2.Error(), ShouldResemble, "Input stack must only include '+' or '-'")

			})
		})
	})
}

func TestFlip(t *testing.T) {
	Convey("Given input stack", t, func() {
		input := pancakeStack{true, false, false, true}
		topTwoFlipped := pancakeStack{false, true, false, true}

		Convey("When the flip function is called for the first two items", func() {
			out := flip(input, 2)

			Convey("The output should be inverted up to N", func() {
				So(out, ShouldResemble, topTwoFlipped)

			})
		})
	})
}
