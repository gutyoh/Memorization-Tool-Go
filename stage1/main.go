package main

/*
[Memorization Tool - Stage 1/4: Make your own flashcards](https://hyperskill.org/projects/159/stages/826/implement)
-------------------------------------------------------------------------------
[Primitive types](https://hyperskill.org/learn/topic/1807)
[Input/Output](https://hyperskill.org/learn/topic/1506)
[Slices](https://hyperskill.org/learn/topic/1672)
[Control statements](https://hyperskill.org/learn/topic/1728)
[Loops](https://hyperskill.org/learn/topic/1531)
[Advanced Input](https://hyperskill.org/learn/topic/2027)
[Errors](https://hyperskill.org/learn/topic/1795)
[Debugging Go code](https://hyperskill.org/learn/step/23076)
[Maps](https://hyperskill.org/learn/step/16999)
[Unicode package](https://hyperskill.org/learn/step/23087)
[Functions](https://hyperskill.org/learn/topic/1750)
[Single Responsibility Principle](https://hyperskill.org/learn/step/8963)
[Functional decomposition](https://hyperskill.org/learn/topic/1893)
*/

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func isASCII(str string) bool {
	for _, char := range str {
		if char > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func getValidInput(prompt string, scanner *bufio.Scanner) string {
	fmt.Println(prompt)
	scanner.Scan()
	input := scanner.Text()
	for len(input) == 1 || !isASCII(input) || input == "" {
		fmt.Println(prompt)
		scanner.Scan()
		input = scanner.Text()
	}
	return input
}

func buildFlashcard(flashcards map[string]string) {
	var question, answer string
	scanner := bufio.NewScanner(os.Stdin)

	question = getValidInput("Question:", scanner)
	answer = getValidInput("Answer:", scanner)

	flashcards[question] = answer
}

func practiceFlashcards(flashcards map[string]string) {
	if len(flashcards) == 0 {
		fmt.Println("There are no flashcards to practice!")
		return
	}

	for key := range flashcards {
		fmt.Println("Question:", key)
		fmt.Println("Please press \"y\" to see the answer or press \"n\" to skip:")
		var input string
		fmt.Scanln(&input)

		switch input {
		case "y":
			fmt.Println("Answer:", flashcards[key])
		case "n":
			continue
		default:
			fmt.Printf("%s is not an option\n", input)
		}
	}
}

func addFlashcardsMenu(flashcards map[string]string) {
	for {
		fmt.Println("1. Add a new flashcard")
		fmt.Println("2. Exit")

		var flashcardChoice string
		fmt.Scanln(&flashcardChoice)

		switch flashcardChoice {
		case "1":
			buildFlashcard(flashcards)
		case "2":
			return
		default:
			fmt.Println(flashcardChoice, "is not an option")
		}
	}
}

func mainMenuSelection(flashcards map[string]string) {
	for {
		fmt.Println("1. Add flashcards")
		fmt.Println("2. Practice flashcards")
		fmt.Println("3. Exit")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			addFlashcardsMenu(flashcards)
		case "2":
			practiceFlashcards(flashcards)
		case "3":
			fmt.Println("Bye!")
			return
		default:
			fmt.Println(choice, "is not an option")
		}
	}
}

func main() {
	flashcards := make(map[string]string)
	mainMenuSelection(flashcards)
}
