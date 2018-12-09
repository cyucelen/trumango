package nlp

import (
	"github.com/bbalet/stopwords"
	"github.com/dchest/stemmer/porter2"
	"strings"
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
