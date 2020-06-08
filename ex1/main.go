package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open file : %s", *csvFilename))
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to parse file : %s", *csvFilename))
	}

	problems := parseLines(lines)
	// fmt.Println(problems)

	score := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			score++
		}
	}
	fmt.Printf("you scored %d out of %d\n", score, len(problems))

}

func parseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))
	for i, line := range lines {
		ret[i] = Problem{
			q: line[0],
			a: strings.Trim(line[1], " "),
		}
	}
	return ret
}

type Problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Print(msg)
	os.Exit(1)
}
