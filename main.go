package main

import (
	"time"

	"github.com/cyucelen/trumango/truman"
	"github.com/fatih/color"
)

var green = color.New(color.FgGreen)

func main() {
	total := time.Now()

	t := truman.New("./questions.txt", "./the_truman_show_script.txt")

	start := time.Now()

	t.PrintAnswersSentences()

	elapsed := time.Since(start)
	totalElapsed := time.Since(total)

	green.Printf("Answer time: %s \n", elapsed)
	green.Printf("Total time : %s \n", totalElapsed)

	color.Unset()
}
