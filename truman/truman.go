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
	textSentences []string
	questions     []string
}

// New loads questions and text from given file paths, splits text into sentences and
// returns a Truman instance
func New(questionsPath string, textPath string) *Truman {
	t := &Truman{}

	questionsFilePath, _ := filepath.Abs(questionsPath)
	textFilePath, _ := filepath.Abs(textPath)

	t.loadQuestions(questionsFilePath)
	t.loadText(textFilePath)

	return t
}

func (t *Truman) loadText(textFilePath string) {
	textByte, _ := ioutil.ReadFile(textFilePath)
	text := strings.Trim(string(textByte), "\n")
	t.textSentences = nlp.SplitSentences(text)
}

func (t *Truman) loadQuestions(questionsFilePath string) {
	questionsFile, _ := os.Open(questionsFilePath)
	defer questionsFile.Close()

	scanner := bufio.NewScanner(questionsFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		t.questions = append(t.questions, scanner.Text())
	}
}

func (t *Truman) Answer() {
	answerSentences := t.findAnswerSentenceToEachQuestion()
	exactAnswers := findExactAnswerFromEachSentence(answerSentences)
	t.printResults(exactAnswers)
}

func (t *Truman) findAnswerSentenceToEachQuestion() map[string]string {
	answerSentences := make(map[string]string)

	var wg sync.WaitGroup
	wg.Add(len(t.questions))
	for _, question := range t.questions {
		go func(question string) {
			answerSentences[question] = t.findAnswerSentenceToQuestion(question)
			wg.Done()
		}(question)
	}
	wg.Wait()

	return answerSentences
}

func (t *Truman) findAnswerSentenceToQuestion(question string) string {
	cleanQuestion := nlp.ClearAndStem(question)
	cleanQuestionWords := strings.Split(cleanQuestion, " ")

	mostMatchSentence := ""
	var mostMatchPercentage float64

	for i := 0; i < len(t.textSentences); i++ {
		matchPercentage := calculateWordMatchPercentage(t.textSentences[i], cleanQuestionWords)
		if matchPercentage > mostMatchPercentage {
			mostMatchPercentage = matchPercentage
			mostMatchSentence = t.textSentences[i]
		}
	}

	return mostMatchSentence
}

func calculateWordMatchPercentage(text string, patterns []string) float64 {
	var matchCount float64
	for j := 0; j < len(patterns); j++ {
		if horspool.Find(strings.ToLower(text), patterns[j]) != -1 {
			matchCount++
		}
	}
	return matchCount / float64(len(patterns)) * 100
}

func findExactAnswerFromEachSentence(sentences map[string]string) map[string]string {
	exactAnswers := make(map[string]string)
	for question, answerSentence := range sentences {
		exactAnswers[question] = findExactAnswerFromSentence(question, answerSentence)
	}
	return exactAnswers
}

func findExactAnswerFromSentence(question string, sentence string) string {
	stemmedToOriginalWordMap := make(map[string]string)

	cleanAnswerSentence := nlp.ClearStopWords(sentence)
	cleanAnswerWords := strings.Split(cleanAnswerSentence, " ")
	cleanStemmedAnswerWords := []string{}

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

func (t *Truman) printResults(qa map[string]string) {
	for question, answer := range qa {
		red.Printf("Question : %s\n", white(question))
		blue.Printf("Answer   : %s\n", white(answer))
		fmt.Println("-------------------------------")
	}
}
