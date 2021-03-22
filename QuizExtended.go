package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func askQuestion(id int, question string) {
	prefix := "Quiz question #" + fmt.Sprint(id)
	fmt.Print(prefix + ": " + question + "? ")
}

func readAndCheckAnswer(answer string) bool {
	var userAnswer string
	fmt.Scanf("%s", &userAnswer)
	return answer == userAnswer
}

func main() {
	quizFileFlag := flag.String("file", "problems.csv", "Fully qualified file path of the file containing the quizes")
	timeoutFlag := flag.Int("time", 60, "Time (in seconds) under which the quiz should be completed")
	flag.Parse()
	file, err := os.Open(*quizFileFlag)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The quiz is going to start with the time limit " + fmt.Sprint(*timeoutFlag) + ". Press enter to begin")
	fmt.Scanln()
	score, maxScore := 0, len(lines)
	printFinalScore := func() {
		fmt.Println("\nTotal user score: " + fmt.Sprint(score) + " out of " + fmt.Sprint(maxScore))
		os.Exit(0)
	}
	timeoutTimer := time.AfterFunc(time.Duration(*timeoutFlag)*time.Second, printFinalScore)
	for i, line := range lines {
		question, answer := line[0], line[1]
		askQuestion(i+1, question)
		if readAndCheckAnswer(answer) {
			score++
		}
	}
	timeoutTimer.Stop()
	printFinalScore()
}
