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
			res1, err1 := calculateFlipsForSingleStack(stack1)
			res2, err2 := calculateFlipsForSingleStack(stack2)
			res3, err3 := calculateFlipsForSingleStack(stack3)
			res4, err4 := calculateFlipsForSingleStack(stack4)
			res5, err5 := calculateFlipsForSingleStack(stack5)

			Convey("The output should match sample output", func() {
				So(err1, ShouldBeNil)
				So(err2, ShouldBeNil)
				So(err3, ShouldBeNil)
				So(err4, ShouldBeNil)
				So(err5, ShouldBeNil)

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
