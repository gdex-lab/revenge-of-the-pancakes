package io

import (
	"github.com/gdexlab/revenge-of-the-pancakes/types"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestConvertInputToPancakeStack(t *testing.T) {
	Convey("Given input string", t, func() {
		validInput := "+--+"
		validOutput := types.PancakeStack{true, false, false, true}
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
