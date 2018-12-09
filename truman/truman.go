package truman

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/cyucelen/trumango/horspool"
	"github.com/cyucelen/trumango/nlp"
	"gopkg.in/jdkato/prose.v2"
)

type Truman struct {
	text      string
	sentences []string
	questions []*prose.Document
}

func New(questionsPath string, textPath string) *Truman {
	t := &Truman{}
	questionsFilePath, _ := filepath.Abs(questionsPath)
	textFilePath, _ := filepath.Abs(textPath)

	t.loadQuestions(questionsFilePath)
	t.loadText(textFilePath)

	textDoc, _ := prose.NewDocument(t.text, prose.WithExtraction(false), prose.WithTagging(false), prose.WithTokenization(false))
	sentences := textDoc.Sentences()

	for _, sentence := range sentences {
		t.sentences = append(t.sentences, sentence.Text)
	}

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
		questionDoc, _ := prose.NewDocument(scanner.Text(), prose.WithSegmentation(false))
		t.questions = append(t.questions, questionDoc)
	}
}

func (t *Truman) findAnswerSentence(index int) string {
	question := t.questions[index].Text
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
	var matchCount float64 = 0
	for j := 0; j < len(patterns); j++ {
		if horspool.Find(text, patterns[j]) != -1 {
			matchCount++
		}
	}

	return matchCount / float64(len(patterns)) * 100
}
