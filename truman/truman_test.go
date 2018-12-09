package truman

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_New(t *testing.T) {
	Convey("Given a questions file and a text file", t, func() {
		questionsPath := "./test_questions.txt"
		textPath := "./test_text.txt"
		Convey("When New called with questionsFile and textFile", func() {
			t := New(questionsPath, textPath)
			Convey("Then t.questions.Text should resemble expectedQuestions", func() {
				expectedQuestions := []string{
					"Who is suddenly aware that the hundreds of other beachgoers have stopped their activities to stare at him?",
					"What color is the cardigan she is already removing from the drycleaning bag on the back seat as Truman pulls away from the curb?",
				}
				for i := range t.questions {
					So(t.questions[i].Text, ShouldResemble, expectedQuestions[i])
				}
			})

			Convey("then t.text should equal to expectedText", func() {
				expectedText := `THE TRUMAN SHOW A Screen Play By Andrew M. Niccol FADE IN A white title appears on a black screen. "One doesn't discover new lands without consenting to lose sight of the shore for a very long time." Andre Gide The title fades off, replaced by a second title.`
				So(t.text, ShouldResemble, expectedText)
			})
			Convey("then t.sentences should resemble expectedSenteces", func() {
				expectedSentences := []string{
					"THE TRUMAN SHOW A Screen Play By Andrew M. Niccol FADE IN A white title appears on a black screen.",
					"\"One doesn't discover new lands without consenting to lose sight of the shore for a very long time.\"",
					"Andre Gide The title fades off, replaced by a second title."}
				So(t.sentences, ShouldResemble, expectedSentences)
			})
		})
	})
}

func Test_findAnswerSentence(t *testing.T) {
	Convey("Given Truman instance", t, func() {
		questionsPath := "./test_questions.txt"
		textPath := "./test_answer_text.txt"
		t := New(questionsPath, textPath)
		Convey("When t.findAnswerSentence called with index 0", func() {
			actualAnswerSentence := t.findAnswerSentence(0)
			Convey("Then actualAnswerSentence should equal to expectedAnswerSentence", func() {
				expectedAnswerSentence := "Truman is suddenly aware that the hundreds of other beachgoers have stopped their activities to stare at him."
				So(actualAnswerSentence, ShouldEqual, expectedAnswerSentence)
			})
		})
		Convey("When t.findAnswerSentence called with index 1", func() {
			actualAnswerSentence := t.findAnswerSentence(1)
			Convey("Then actualAnswerSentence should equal to expectedAnswerSentence", func() {
				expectedAnswerSentence := "As Truman pulls away from the curb, she is already removing the lavender cardigan from the drycleaning bag on the back seat."
				So(actualAnswerSentence, ShouldEqual, expectedAnswerSentence)
			})
		})
	})
}
