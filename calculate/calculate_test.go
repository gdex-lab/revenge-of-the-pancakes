package calculate

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCalculateFlipsForSingleStack(t *testing.T) {
	Convey("Given example stacks as input", t, func() {
		stack1 := pancakeStack{false}
		stack2 := pancakeStack{false, true}
		stack3 := pancakeStack{true, false}
		stack4 := pancakeStack{true, true, true}
		stack5 := pancakeStack{false, false, true, false}

		Convey("When the calculate function is called", func() {
			res1 := flipsForSingleStack(stack1, 0)
			res2 := flipsForSingleStack(stack2, 0)
			res3 := flipsForSingleStack(stack3, 0)
			res4 := flipsForSingleStack(stack4, 0)
			res5 := flipsForSingleStack(stack5, 0)

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
