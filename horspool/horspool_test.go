package horspool

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_createShiftTable(t *testing.T) {
	Convey("Given a pattern", t, func() {
		pattern := "foobar"
		Convey("When createShiftTable called with pattern", func() {
			actualShiftTable := createShiftTable(pattern)
			Convey("Then actualShiftTable should resemble expectedShiftTable", func() {
				expectedShiftTable := make(map[byte]int)
				expectedShiftTable[pattern[0]] = 5
				expectedShiftTable[pattern[2]] = 3
				expectedShiftTable[pattern[3]] = 2
				expectedShiftTable[pattern[4]] = 1
				So(actualShiftTable, ShouldResemble, expectedShiftTable)
			})
		})
	})

	Convey("Given a pattern", t, func() {
		pattern := "color"
		Convey("When createShiftTable called with pattern", func() {
			actualShiftTable := createShiftTable(pattern)
			Convey("Then actualShiftTable should resemble expectedShiftTable", func() {
				expectedShiftTable := make(map[byte]int)
				expectedShiftTable[pattern[0]] = 4
				expectedShiftTable[pattern[2]] = 2
				expectedShiftTable[pattern[3]] = 1
				So(actualShiftTable, ShouldResemble, expectedShiftTable)
			})
		})
	})
}

func Test_createReverseShiftTable(t *testing.T) {
	Convey("Given a pattern", t, func() {
		pattern := "foobar"
		Convey("When createReverseShiftTable called with pattern", func() {
			actualShiftTable := createReverseShiftTable(pattern)
			Convey("Then actualShiftTable should resemble expectedShiftTable", func() {
				expectedShiftTable := make(map[byte]int)
				expectedShiftTable[pattern[2]] = 1
				expectedShiftTable[pattern[3]] = 3
				expectedShiftTable[pattern[4]] = 4
				expectedShiftTable[pattern[5]] = 5
				So(actualShiftTable, ShouldResemble, expectedShiftTable)
			})
		})
	})
}

func Test_Find(t *testing.T) {
	Convey("Given a text and a pattern", t, func() {
		text := "Make everything as simple as possible, but..."
		pattern := "simple"
		Convey("When Find called with text and pattern", func() {
			index := Find(text, pattern)
			Convey("Then it should return 19", func() {
				So(index, ShouldEqual, 19)
			})
		})
	})

	Convey("Given a text that does not contains pattern", t, func() {
		text := "Lorem ipsum dolor, sit amet."
		pattern := "skydome"
		Convey("When Find called with text and pattern", func() {
			index := Find(text, pattern)
			Convey("Then it should return -1", func() {
				So(index, ShouldEqual, -1)
			})
		})
	})

	Convey("Given a text that contains the pattern multiple times", t, func() {
		text := "Make everything as simple as possible, but not simpler."
		pattern := "simple"
		Convey("When Find called with text and pattern", func() {
			index := Find(text, pattern)
			Convey("Then it should return 19", func() {
				So(index, ShouldEqual, 19)
			})
		})
	})

	Convey("Given a text that contains pattern at the end", t, func() {
		text := "truman sudden awar hundr beachgoer stop activ stare"
		pattern := "stare"
		Convey("When Find called with text and pattern", func() {
			index := Find(text, pattern)
			Convey("Then it should return 46", func() {
				So(index, ShouldEqual, 46)
			})
		})
	})

	Convey("Given a text that contains nearly matching version of pattern multiple times", t, func() {
		text := "isnt.. isnt.. isnt... isnt it the way she say to the way and isnt that the way that the saint say"
		pattern := "doesnt"
		Convey("When Find called with text and pattern", func() {
			index := Find(text, pattern)
			Convey("Then it should return -1", func() {
				So(index, ShouldEqual, -1)
			})
		})
	})

	// index out of range
	Convey("Given a text that stemmed and stop words cleaned", t, func() {
		text := "paus intersect quiet working-class suburban street spheric glass object sudden fail sky land deafen crash roadway yard idl car"
		pattern := "awar"
		Convey("When Find called with text and pattern", func() {
			index := Find(text, pattern)
			Convey("Then it should return -1", func() {
				So(index, ShouldEqual, -1)
			})
		})
	})

	Convey("Given a text that stemmed and stop words cleaned", t, func() {
		text := "bus sudden screech halt struggl group door open"
		pattern := "color"
		Convey("When Find called with text and pattern", func() {
			index := Find(text, pattern)
			Convey("Then it should return -1", func() {
				So(index, ShouldEqual, -1)
			})
		})
	})
}

func Test_FindLast(t *testing.T) {
	Convey("Given a text and a pattern", t, func() {
		text := "Make everything as simple as possible but..."
		pattern := "simple"
		Convey("When FindLast called with text and pattern", func() {
			index := FindLast(text, pattern)
			Convey("Then it should return 19", func() {
				So(index, ShouldEqual, 19)
			})
		})
	})

	Convey("Given a text that does not contains pattern", t, func() {
		text := "Lorem ipsum dolor, sit amet."
		pattern := "skydome"
		Convey("When FindLast called with text and pattern", func() {
			index := FindLast(text, pattern)
			Convey("Then it should return -1", func() {
				So(index, ShouldEqual, -1)
			})
		})
	})

	Convey("Given a text that contains the pattern multiple times", t, func() {
		text := "Make everything as simple as simple possible"
		pattern := "simple"
		Convey("When FindLast called with text and pattern", func() {
			index := FindLast(text, pattern)
			Convey("Then it should return 29", func() {
				So(index, ShouldEqual, 29)
			})
		})
	})

	Convey("Given a text that contains nearly matching version of pattern multiple times", t, func() {
		text := "isnt.. isnt.. isnt... isnt it the way she say to the way and isnt that the way that the saint say"
		pattern := "doesnt"
		Convey("When FindLast called with text and pattern", func() {
			index := FindLast(text, pattern)
			Convey("Then it should return -1", func() {
				So(index, ShouldEqual, -1)
			})
		})
	})
}
