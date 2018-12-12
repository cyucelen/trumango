package util

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Difference(t *testing.T) {
	Convey("Given two string slices", t, func() {
		sliceOne := []string{"foo", "bar", "too", "far"}
		sliceTwo := []string{"bar", "far"}
		Convey("When Difference called with given slices", func() {
			actualDifference := Difference(sliceOne, sliceTwo)
			Convey("Then actualDifference should equal to expectedDifference", func() {
				expectedDifference := []string{"foo", "too"}
				So(actualDifference, ShouldResemble, expectedDifference)
			})
		})
	})
}
