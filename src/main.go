package main

import (
	"fmt"
	"satriyoaji/prospace-code-challenge-advance/src/app"
)

func main() {
	converter := app.NewIntergalacticConverter()

	// Example input lines
	inputLines := []string{
		"glob is I",
		"prok is V",
		"pish is X",
		"tegj is L",
		"glob glob Silver is 34 Credits",
		"glob prok Gold is 57800 Credits",
		"pish pish Iron is 3910 Credits",
		"how much is pish tegj glob glob ?",
		"how many Credits is glob prok Silver ?",
		"how many Credits is glob prok Gold ?",
		"how many Credits is glob prok Iron ?",
		"how much wood could a woodchuck chuck if a woodchuck could chuck wood ?",
	}

	for _, line := range inputLines {
		response := converter.ProcessInput(line)
		if response != nil {
			fmt.Println(response)
		}
	}
}
