package truman

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_New(t *testing.T) {
	Convey("Given a questions file and a text file", t, func() {
		questionsPath := "./testdata/test_questions.txt"
		textPath := "./testdata/test_sentence_text.txt"
		Convey("When New called with questionsFile and textFile", func() {
			t := New(questionsPath, textPath, 70)
			Convey("Then expectedQuestions should exist in t.questionAnswerSentences", func() {
				expectedQuestions := []string{
					"Who is suddenly aware that the hundreds of other beachgoers have stopped their activities to stare at him?",
					"What color is the cardigan she is already removing from the drycleaning bag on the back seat as Truman pulls away from the curb?",
				}
				for _, expectedQuestion := range expectedQuestions {
					_, exists := t.questionAnswerSentences[expectedQuestion]
					So(exists, ShouldBeTrue)
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
		questionsPath := "./testdata/test_questions.txt"
		textPath := "./testdata/test_answer_text.txt"
		t := New(questionsPath, textPath, 70)
		Convey("When t.findAnswerSentence called with question", func() {
			question := "Who is suddenly aware that the hundreds of other beachgoers have stopped their activities to stare at him?"
			actualAnswerSentence := t.findAnswerSentence(question)
			Convey("Then actualAnswerSentence should equal to expectedAnswerSentence", func() {
				expectedAnswerSentence := "Truman is suddenly aware that the hundreds of other beachgoers have stopped their activities to stare at him."
				So(actualAnswerSentence, ShouldEqual, expectedAnswerSentence)
			})
		})
		Convey("When t.findAnswerSentence called with another question", func() {
			question := "What color is the cardigan she is already removing from the drycleaning bag on the back seat as Truman pulls away from the curb?"
			actualAnswerSentence := t.findAnswerSentence(question)
			Convey("Then actualAnswerSentence should equal to expectedAnswerSentence", func() {
				expectedAnswerSentence := "As Truman pulls away from the curb, she is already removing the lavender cardigan from the drycleaning bag on the back seat."
				So(actualAnswerSentence, ShouldEqual, expectedAnswerSentence)
			})
		})
	})
}

func Test_findAllAnswerSentences(t *testing.T) {
	Convey("Given Truman instance", t, func() {
		questionsPath := "../questions.txt"
		textPath := "../the_truman_show_script.txt"
		t := New(questionsPath, textPath, 70)
		Convey("When t.findAllAnswerSentecences called", func() {
			actualAnswers := t.findAllAnswerSentences()
			Convey("Then actualAnswers should resemble expected answers", func() {
				expectedAnswers := map[string]string{
					"Who is suddenly aware that the hundreds of other beachgoers have stopped their activities to stare at him?":                       "Truman is suddenly aware that the hundreds of other beachgoers have stopped their activities to stare at him.",
					"With an apathetic shrug, what does Truman replace?":                                                                               "With an apathetic shrug, Truman replaces the receiver.",
					"He picks up the framed picture of his wife from where?":                                                                           "He picks up the framed picture of his wife from his desk.",
					"The sound of the children triggers what in his head?":                                                                             "The sound of the children triggers a memory in his head.",
					"What does Truman exit to investigate?":                                                                                            "Truman exits the Oldsmobile to investigate.",
					"How many girls look up, surprised by the interruption?":                                                                           "The 3 girls look up, surprised by the interruption.",
					"What time is read on a clock on the wall?":                                                                                        "A clock on the wall reads 4:12pm.",
					"What color is the cardigan she is already removing from the drycleaning bag on the back seat as Truman pulls away from the curb?": "As Truman pulls away from the curb, she is already removing the lavender cardigan from the drycleaning bag on the back seat.",
				}
				So(actualAnswers, ShouldResemble, expectedAnswers)
			})
		})
	})
}

func Test_findExactAnswers(t *testing.T) {
	Convey("Given Truman instance and question-answerSentence map", t, func() {
		questionsPath := "../questions.txt"
		textPath := "../the_truman_show_script.txt"
		t := New(questionsPath, textPath, 70)
		questionAnswerSentences := t.findAllAnswerSentences()
		Convey("When t.findExactAnswers called with question-answerSentence map", func() {
			actualExactAnswers := t.findExactAnswers(questionAnswerSentences)
			Convey("Then actualExactAnswer[question0] should equal to expectedExactAnswer", func() {
				question := "Who is suddenly aware that the hundreds of other beachgoers have stopped their activities to stare at him?"
				expectedExactAnswer := "truman"
				So(actualExactAnswers[question], ShouldEqual, expectedExactAnswer)
			})
			Convey("Then actualExactAnswer[question1] should equal to expectedExactAnswer", func() {
				question := "How many girls look up, surprised by the interruption?"
				expectedExactAnswer := "3"
				So(actualExactAnswers[question], ShouldEqual, expectedExactAnswer)
			})
			Convey("Then actualExactAnswer[question2] should equal to expectedExactAnswer", func() {
				question := "With an apathetic shrug, what does Truman replace?"
				expectedExactAnswer := "receiver"
				So(actualExactAnswers[question], ShouldEqual, expectedExactAnswer)
			})
			Convey("Then actualExactAnswer[question3] should equal to expectedExactAnswer", func() {
				question := "The sound of the children triggers what in his head?"
				expectedExactAnswer := "memory"
				So(actualExactAnswers[question], ShouldEqual, expectedExactAnswer)
			})
			Convey("Then actualExactAnswer[question4] should equal to expectedExactAnswer", func() {
				question := "He picks up the framed picture of his wife from where?"
				expectedExactAnswer := "desk"
				So(actualExactAnswers[question], ShouldEqual, expectedExactAnswer)
			})
			Convey("Then actualExactAnswer[question5] should equal to expectedExactAnswer", func() {
				question := "What does Truman exit to investigate?"
				expectedExactAnswer := "oldsmobile"
				So(actualExactAnswers[question], ShouldEqual, expectedExactAnswer)
			})
			Convey("Then actualExactAnswer[question6] should equal to expectedExactAnswer", func() {
				question := "What color is the cardigan she is already removing from the drycleaning bag on the back seat as Truman pulls away from the curb?"
				expectedExactAnswer := "lavender"
				So(actualExactAnswers[question], ShouldEqual, expectedExactAnswer)
			})
			Convey("Then actualExactAnswer[question7] should equal to expectedExactAnswer", func() {
				question := "What time is read on a clock on the wall?"
				expectedExactAnswer := "4 12pm"
				So(actualExactAnswers[question], ShouldEqual, expectedExactAnswer)
			})
		})
	})
}
