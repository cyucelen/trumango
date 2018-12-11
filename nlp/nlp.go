package nlp

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bbalet/stopwords"
	"github.com/dchest/stemmer/porter2"
	prose "gopkg.in/jdkato/prose.v2"
)

const languageCode = "en"

func Stem(text string) string {
	words := strings.Fields(text)
	for i, word := range words {
		words[i] = porter2.Stemmer.Stem(word)
	}
	return strings.Join(words, " ")
}

func ClearStopWords(text string) string {
	stopwords.DontStripDigits()
	cleanText := stopwords.CleanString(text, languageCode, false)
	return strings.Trim(cleanText, " ")
}

func ClearAndStem(text string) string {
	clearText := ClearStopWords(text)
	return Stem(clearText)
}

func SplitSentences(text string) []string {
	textDoc, _ := prose.NewDocument(text, prose.WithExtraction(false), prose.WithTagging(false), prose.WithTokenization(false))
	sentences := textDoc.Sentences()
	sentencesText := make([]string, 0, len(sentences))

	for _, sentence := range sentences {
		sentencesText = append(sentencesText, sentence.Text)
	}

	return sentencesText
}

func Tokenize(text string) *prose.Document {
	textDoc, _ := prose.NewDocument(text, prose.WithSegmentation(false))

	return textDoc
}

func TokenizeMap(m map[string]string) map[*prose.Document]*prose.Document {
	tokenizedMap := make(map[*prose.Document]*prose.Document)
	var wg sync.WaitGroup
	wg.Add(len(m))

	for key, value := range m {
		go func(key string, value string) {
			tokenizedKey := Tokenize(key)
			tokenizedMap[tokenizedKey] = Tokenize(value)
			wg.Done()
		}(key, value)
	}

	wg.Wait()
	return tokenizedMap
}

func PrintTokens(textDoc *prose.Document) {
	for _, tok := range textDoc.Tokens() {
		if tok.Tag != "," && tok.Tag != "?" && tok.Tag != "." {
			fmt.Printf("%-12s", tok.Tag)
		}
	}
	fmt.Println()
}
