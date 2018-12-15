package truman

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/fatih/color"

	"github.com/cyucelen/trumango/horspool"
	"github.com/cyucelen/trumango/nlp"
	"github.com/cyucelen/trumango/util"
)

var yellow = color.New(color.FgYellow)
var red = color.New(color.FgRed)
var blue = color.New(color.FgBlue)
var white = color.New(color.FgWhite).SprintFunc()

// Truman is happy to answer your questions :)
type Truman struct {
	text                    string
	sentences               []string
	questionAnswerSentences map[string]string
	matchPercentage         float64
}

// New loads questions and text from given file paths, splits text into sentences and
// returns a Truman instance
func New(questionsPath string, textPath string, matchPercentage float64) *Truman {
	t := &Truman{}
	t.sentences = make([]string, 0, 10000)
	t.questionAnswerSentences = make(map[string]string)

	t.matchPercentage = matchPercentage

	questionsFilePath, _ := filepath.Abs(questionsPath)
	textFilePath, _ := filepath.Abs(textPath)

	t.loadQuestions(questionsFilePath)
	t.loadText(textFilePath)

	t.sentences = nlp.SplitSentences(t.text)

	return t
}

func (t *Truman) loadText(textFilePath string) {
	textByte, _ := ioutil.ReadFile(textFilePath)
	t.text = strings.Trim(string(textByte), "\n")
}

func (t *Truman) loadQuestions(questionsFilePath string) {
	questionsFile, _ := os.Open(questionsFilePath)
	defer questionsFile.Close()

	scanner := bufio.NewScanner(questionsFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		question := scanner.Text()
		t.questionAnswerSentences[question] = ""
	}
}

// Answer answers the given questions by printing
func (t *Truman) Answer() {
	questionAnswerSentences := t.findAllAnswerSentences()
	questionExactAnswer := t.findExactAnswers(questionAnswerSentences)

	t.printQuestionAnswerMap(questionExactAnswer)
}

func (t *Truman) findExactAnswers(map[string]string) map[string]string {
	questionExactAnswers := make(map[string]string)

	for question, answerSentence := range t.questionAnswerSentences {
		questionExactAnswers[question] = t.findExactAnswer(question, answerSentence)
	}

	return questionExactAnswers
}

func (t *Truman) findExactAnswer(question string, answerSentence string) string {
	stemmedToOriginalWordMap := make(map[string]string)

	cleanAnswerSentence := nlp.ClearStopWords(answerSentence)
	cleanAnswerWords := strings.Split(cleanAnswerSentence, " ")
	cleanStemmedAnswerWords := make([]string, 0)

	for _, word := range cleanAnswerWords {
		stemmedWord := nlp.Stem(word)
		cleanStemmedAnswerWords = append(cleanStemmedAnswerWords, stemmedWord)
		stemmedToOriginalWordMap[stemmedWord] = word
	}

	cleanStemmedQuestion := nlp.ClearAndStem(question)
	cleanStemmedQuestionWords := strings.Split(cleanStemmedQuestion, " ")

	answers := util.Difference(cleanStemmedAnswerWords, cleanStemmedQuestionWords)

	for i := range answers {
		answers[i] = stemmedToOriginalWordMap[answers[i]]
	}

	return strings.Join(answers, " ")
}

func (t *Truman) findAllAnswerSentences() map[string]string {
	var wg sync.WaitGroup
	wg.Add(len(t.questionAnswerSentences))

	for question := range t.questionAnswerSentences {
		go func(question string) {
			t.questionAnswerSentences[question] = t.findAnswerSentence(question)
			wg.Done()
		}(question)
	}

	wg.Wait()

	return t.questionAnswerSentences
}

func (t *Truman) findAnswerSentence(question string) string {
	cleanQuestion := nlp.ClearAndStem(question)
	cleanQuestionWords := strings.Split(cleanQuestion, " ")

	for i := 0; i < len(t.sentences); i++ {
		cleanSentence := nlp.ClearAndStem(t.sentences[i])
		if calculateMatchPercentage(cleanSentence, cleanQuestionWords) > t.matchPercentage {
			return t.sentences[i]
		}
	}

	return ""
}

func calculateMatchPercentage(text string, patterns []string) float64 {
	var matchCount float64
	for j := 0; j < len(patterns); j++ {
		if horspool.Find(text, patterns[j]) != -1 {
			matchCount++
		}
	}

	return matchCount / float64(len(patterns)) * 100
}

func (t *Truman) printQuestionAnswerMap(qa map[string]string) {
	for question, answer := range qa {
		red.Printf("Question : %s\n", white(question))
		blue.Printf("Answer   : %s\n", white(answer))
		fmt.Println("-------------------------------")
	}
}
