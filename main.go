package main

import (
	"flag"
	"time"

	"github.com/cyucelen/trumango/truman"
	"github.com/fatih/color"
)

var green = color.New(color.FgGreen)

func main() {

	var questionFile string
	var textFile string
	var matchPercentage float64

	flag.StringVar(&questionFile, "q", "questions.txt", "questions file path")
	flag.StringVar(&textFile, "t", "the_truman_show_script.txt", "text file path")
	flag.Float64Var(&matchPercentage, "m", 70, "desired match percentage for accepting a sentence as answer sentence")
	flag.Parse()

	start := time.Now()

	t := truman.New(questionFile, textFile, matchPercentage)
	t.Answer()

	elapsed := time.Since(start)
	green.Printf("Answer time: %s \n", elapsed)
}
