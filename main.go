package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/Sraik25/quiz/utils"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
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
	correct := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.Q)
		var answer string
		fmt.Scanf("%s \n", &answer)

		if answer == p.A {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}
