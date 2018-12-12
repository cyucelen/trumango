package nlp

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Stem(t *testing.T) {
	Convey("Given a text consist of one word", t, func() {
		text := "surprised"
		Convey("When Stem called with text", func() {
			actualStemmedText := Stem(text)
			Convey("Then actualStemmedText should equal to expectedStemmedText", func() {
				expectedStemmedText := "surpris"
				So(actualStemmedText, ShouldEqual, expectedStemmedText)
			})
		})
	})

	Convey("Given a text consist of multiple words", t, func() {
		text := "Who is suddenly aware that the hundreds of other beachgoers have stopped their activities to stare at him?"
		Convey("When Stem called with text", func() {
			actualStemmedText := Stem(text)
			Convey("Then actualStemmedText should equal to expectedStemmedText", func() {
				expectedStemmedText := "who is sudden awar that the hundr of other beachgoer have stop their activ to stare at him?"
				So(actualStemmedText, ShouldEqual, expectedStemmedText)
			})
		})
	})
}

func Test_ClearStopWords(t *testing.T) {
	Convey("Given a text consist of multiple words", t, func() {
		text := "What time is read on a clock on the wall?"
		Convey("When ClearStopWords called with text", func() {
			actualCleanText := ClearStopWords(text)
			Convey("Then actualStemmedText should equal to expectedStemmedText", func() {
				expectedCleanText := "time read clock wall"
				So(actualCleanText, ShouldEqual, expectedCleanText)
			})
		})
	})

	Convey("Given a text consist of multiple words and numbers", t, func() {
		text := "What time is read on a clock on the wall? is it 5"
		Convey("When ClearStopWords called with text", func() {
			actualCleanText := ClearStopWords(text)
			Convey("Then actualStemmedText should equal to expectedStemmedText", func() {
				expectedCleanText := "time read clock wall 5"
				So(actualCleanText, ShouldEqual, expectedCleanText)
			})
		})
	})

}

func Test_ClearAndStem(t *testing.T) {
	Convey("Given a text consist of multiple words", t, func() {
		text := "Who is suddenly aware that the hundreds of other beachgoers have stopped their activities to stare at him?"
		Convey("When Stem called with text", func() {
			actualCleanText := ClearAndStem(text)
			Convey("Then actualCleanText should equal to expectedCleanText", func() {
				expectedCleanText := "sudden awar hundr beachgoer stop activ stare"
				So(actualCleanText, ShouldEqual, expectedCleanText)
			})
		})
	})
}

func Test_SplitSentences(t *testing.T) {
	Convey("Given a text consis of multiple sentences", t, func() {
		text := `THE TRUMAN SHOW A Screen Play By Andrew M. Niccol FADE IN A white title appears on a black screen. "One doesn't discover new lands without consenting to lose sight of the shore for a very long time." Andre Gide The title fades off, replaced by a second title.`
		Convey("When SplitSentences called with text", func() {
			actualSentences := SplitSentences(text)
			Convey("Then actualSentences should resemble expectedSentences", func() {
				expectedSentences := []string{
					"THE TRUMAN SHOW A Screen Play By Andrew M. Niccol FADE IN A white title appears on a black screen.",
					"\"One doesn't discover new lands without consenting to lose sight of the shore for a very long time.\"",
					"Andre Gide The title fades off, replaced by a second title."}
				So(actualSentences, ShouldResemble, expectedSentences)
			})
		})
	})
}
