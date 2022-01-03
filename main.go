package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var csvFile string
	flag.StringVar(&csvFile, "csv", "questions.csv", "This file contains list of questions in the format of 'question, answer'")
	flag.Parse()

	file, err := os.Open(csvFile)

	if err != nil {
		errMsg := fmt.Sprintf("Error occurred: %s", err)
		exit(errMsg)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to read file")
	}

	 parseLines(lines);
}

func parseLines(input [][]string) {
	questions := make([]Question, len(input))

	for i, line := range input {
		questions[i] = Question{
			question: line[0],
			answer: strings.TrimSpace(line[1]),
		}
	}

	var correct int

	for i, question := range questions {
		fmt.Printf("Question #%d: %s \n", i + 1, question.question)

		var answer string
		fmt.Scanf("%s\n", &answer)

		if answer == question.answer {
			correct++
			fmt.Println("Correct!")
		} else {
			fmt.Println("Ooops, incorrect!")
		}
	}

	fmt.Printf("You scored %d out of %d \n", correct, len(questions))
}

type Question struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
