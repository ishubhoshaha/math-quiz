package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer string
}

func parseLines(lines [][]string)[]problem {
	returnVal := make([]problem, len(lines))
	for i, line := range lines {
		returnVal[i] = problem{
			question: line[0],
			answer: strings.TrimSpace(line[1]),
		}

	}
	fmt.Println(returnVal)
	return returnVal
}

func errorOnExit(msg string)  {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "csv file with question,answer format")
	timeLimit := flag.Int("limit", 30,"hello")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		errorOnExit(fmt.Sprintf("Failed to open the CSV file"))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		errorOnExit("Failed to parse csv file")
	}
	problems := parseLines(lines)
	correct := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	problemLoop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		answerChannel := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n",&answer)
			answerChannel <- answer
		}()
		select {
		case <-timer.C:
			fmt.Println()
			break problemLoop
		case answer := <-answerChannel:
			if answer == p.answer {
				correct++
			}
		}


	}
	fmt.Printf("You scored %d out of %d in %d seconds.\n", correct, len(lines), *timeLimit)
}

