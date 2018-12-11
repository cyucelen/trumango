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
	prose "gopkg.in/jdkato/prose.v2"

	"github.com/cyucelen/trumango/horspool"
	"github.com/cyucelen/trumango/nlp"
)

var yellow = color.New(color.FgYellow)
var red = color.New(color.FgRed)
var blue = color.New(color.FgBlue)
var white = color.New(color.FgWhite).SprintFunc()

type Truman struct {
	text                    string
	sentences               []string
	questionAnswerSentences map[string]string
}

func New(questionsPath string, textPath string) *Truman {
	t := &Truman{}
	t.sentences = make([]string, 0, 10000)
	t.questionAnswerSentences = make(map[string]string)

	questionsFilePath, _ := filepath.Abs(questionsPath)
	textFilePath, _ := filepath.Abs(textPath)

	t.loadQuestions(questionsFilePath)
	t.loadText(textFilePath)

	t.sentences = nlp.SplitSentences(t.text)

	return t
}

func FindNumber(doc *prose.Document) string {
	for _, token := range doc.Tokens() {
		if token.Tag == "CD" {
			return token.Text
		}
	}
	return ""
}

func (t *Truman) findExactAnswers() map[string]string {
	questionExactAnswers := make(map[string]string)
	tokenizedQuestionAnswers := nlp.TokenizeMap(t.questionAnswerSentences)

	for questionDoc, answerSentenceDoc := range tokenizedQuestionAnswers {
		fmt.Println(questionDoc.Text)
		questionExactAnswers[questionDoc.Text] = ""
		for _, token := range questionDoc.Tokens() {
			if token.Tag == "WP" {
				tokenLowerCase := strings.ToLower(token.Text)
				if tokenLowerCase == "who" {
					questionExactAnswers[questionDoc.Text] = answerSentenceDoc.Entities()[0].Text
				}
			} else if token.Tag == "WRB" {
				tokenLowerCase := strings.ToLower(token.Text)
				if tokenLowerCase == "how" {
					questionExactAnswers[questionDoc.Text] = FindNumber(answerSentenceDoc)
				}
			}
		}
	}

	return questionExactAnswers
}

func printWordsWithTab(text string) {
	words := strings.Split(text, " ")
	for _, word := range words {
		fmt.Printf("%-12s", word)
	}
	fmt.Println()
}

func (t *Truman) PrintAnswersSentences() {
	t.findAllAnswerSentences()

	tokenizedQuestionAnswers := nlp.TokenizeMap(t.questionAnswerSentences)

	for questionDoc, answerSentenceDoc := range tokenizedQuestionAnswers {
		red.Printf("Question : ")
		printWordsWithTab(questionDoc.Text)
		yellow.Printf("Tokens   : ")
		nlp.PrintTokens(questionDoc)

		fmt.Println()

		blue.Printf("Answer   : ")
		printWordsWithTab(answerSentenceDoc.Text)
		yellow.Printf("Tokens   : ")
		nlp.PrintTokens(answerSentenceDoc)

		// fmt.Println()

		// clearAnswer := nlp.ClearAndStem(answerSentenceDoc.Text)
		// clearAnswerTokenized := nlp.Tokenize(clearAnswer)
		// blue.Printf("ClrAnswer:")
		// printWordsWithTab(clearAnswer)
		// yellow.Printf("Tokens   : ")
		// nlp.PrintTokens(clearAnswerTokenized)

		fmt.Println("-------------------------------")
	}
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
		if calculateMatchPercentage(cleanSentence, cleanQuestionWords) > 70 {
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
