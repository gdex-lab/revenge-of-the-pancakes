package calculate

import (
	"testing"

	"github.com/gdexlab/revenge-of-the-pancakes/types"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCalculateFlipsForSingleStack(t *testing.T) {
	Convey("Given example stacks as input", t, func() {
		stack1 := types.PancakeStack{false}
		stack2 := types.PancakeStack{false, true}
		stack3 := types.PancakeStack{true, false}
		stack4 := types.PancakeStack{true, true, true}
		stack5 := types.PancakeStack{false, false, true, false}

		Convey("When the calculate function is called", func() {
			res1 := FlipsForSingleStack(stack1, 0)
			res2 := FlipsForSingleStack(stack2, 0)
			res3 := FlipsForSingleStack(stack3, 0)
			res4 := FlipsForSingleStack(stack4, 0)
			res5 := FlipsForSingleStack(stack5, 0)

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
		input := types.PancakeStack{true, false, false, true}
		topTwoFlipped := types.PancakeStack{false, true, false, true}

		Convey("When the flip function is called for the first two items", func() {
			out := flip(input, 2)

			Convey("The output should be inverted up to N", func() {
				So(out, ShouldResemble, topTwoFlipped)

			})
		})
	})
}
