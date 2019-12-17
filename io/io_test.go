package io

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestValidateInput(t *testing.T) {
	Convey("Given input string", t, func() {
		validInput := "+--+"
		invalidInput := "+XX+"

		Convey("When the conversion function is called", func() {
			err := validateInput(validInput)
			err2 := validateInput(invalidInput)
			Convey("The output should match examples", func() {
				So(err, ShouldBeNil)
				So(err2.Error(), ShouldResemble, "Input stack must only include '+' or '-'")

			})
		})
	})
}
