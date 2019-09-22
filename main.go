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

	flag.StringVar(&questionFile, "q", "questions.txt", "questions file path")
	flag.StringVar(&textFile, "t", "the_truman_show_script.txt", "text file path")
	flag.Parse()

	start := time.Now()

	t := truman.New(questionFile, textFile)
	t.Answer()

	elapsed := time.Since(start)
	green.Printf("Answer time: %s \n", elapsed)
}
