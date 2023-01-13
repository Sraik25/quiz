package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/Sraik25/quiz/utils"
	"os"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		utils.Exit(fmt.Sprintf("Failed to open the CSV file: %s \n", *csvFilename))
		os.Exit(1)
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()

	if err != nil {
		utils.Exit("Failed to parse the provided CSV file.")
	}

	problems := utils.ParseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0

problemloop:

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.Q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerCh:
			if answer == p.A {
				correct++
			}
		}

	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}
