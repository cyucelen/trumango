package main

import (
	"time"

	"github.com/cyucelen/trumango/truman"
	"github.com/fatih/color"
)

var green = color.New(color.FgGreen)

func main() {
	start := time.Now()

	t := truman.New("./questions.txt", "./the_truman_show_script.txt")

	t.Answer()

	elapsed := time.Since(start)

	green.Printf("Answer time: %s \n", elapsed)
}
