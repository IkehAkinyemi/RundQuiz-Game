package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var csvFile string
	flag.StringVar(&csvFile, "csv", "questions.csv", "This file contains list of questions in the format of 'question, answer'")
	timeLimit := flag.Int("limit", 50, "Set time limet for the quiz")
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

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	questions := parseLines(lines)

	var correct int

	for i, question := range questions {
		fmt.Printf("Question #%d: %s = ", i+1, question.question)

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d \n", correct, len(questions))
			return
		case answer := <-answerCh:
			if answer == question.answer {
				correct++
				fmt.Println("Correct!")
			} else {
				fmt.Println("Ooops, incorrect!")
			}
		}

	}

	fmt.Printf("You scored %d out of %d \n", correct, len(questions))
}

func parseLines(input [][]string) []Question {
	questions := make([]Question, len(input))

	for i, line := range input {
		questions[i] = Question{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return questions
}

type Question struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
