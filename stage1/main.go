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
[Public and private scopes](https://hyperskill.org/learn/topic/1894)
*/

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var Flashcards = make(map[string]string)

func DisplayFlashcardMenu() {
	var flashcardMenu = [2]string{"1. Add a new flashcard", "2. Exit"}
	for _, elem := range flashcardMenu {
		fmt.Println(elem)
	}
}

func DisplayMainMenu() {
	var mainMenu = [3]string{"1. Add flashcards", "2. Practice flashcards", "3. Exit"}
	for _, elem := range mainMenu {
		fmt.Println(elem)
	}
}

func isASCII(s string) bool {
	for _, c := range s {
		if c > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func BuildFlashcard() {
	var question, answer string
	scanner := bufio.NewScanner(os.Stdin)

	for len(question) == 1 || !isASCII(question) || question == "" {
		fmt.Println("Question:")
		scanner.Scan()
		question = scanner.Text()
	}

	for len(answer) == 1 || !isASCII(answer) || answer == "" {
		fmt.Println("Answer:")
		scanner.Scan()
		answer = scanner.Text()
	}

	Flashcards[question] = answer

	var mainMenuChoice string
	DisplayFlashcardMenu()
	fmt.Scanln(&mainMenuChoice)
	FlashcardMenuSelection(mainMenuChoice)
}

func MainMenuSelection(choice string) {
	var mainMenuChoice string
	switch choice {
	case "1":
		DisplayFlashcardMenu()
		fmt.Scanln(&mainMenuChoice)
		FlashcardMenuSelection(mainMenuChoice)
	case "2":
		PracticeFlashcards()
	case "3":
		fmt.Println("Bye!")
		os.Exit(0)
	default:
		fmt.Println(choice, "is not an option")
		DisplayMainMenu()
		fmt.Scanln(&choice)
		MainMenuSelection(choice)
	}
}

func FlashcardMenuSelection(choice string) {
	switch choice {
	case "1":
		BuildFlashcard()
	case "2":
		DisplayMainMenu()
		var mainMenuChoice string
		fmt.Scanln(&mainMenuChoice)
		MainMenuSelection(mainMenuChoice)
	default:
		fmt.Println(choice, "is not an option")
		DisplayFlashcardMenu()
		fmt.Scanln(&choice)
		FlashcardMenuSelection(choice)
	}
}

func PracticeFlashcards() {
	if len(Flashcards) == 0 {
		fmt.Println("There are no flashcards to practice!")
	}

	if len(Flashcards) > 0 {
		DisplayPracticeMenu()
	}

	DisplayMainMenu()
	var choice string
	fmt.Scanln(&choice)
	MainMenuSelection(choice)
}

func DisplayPracticeMenu() {
	for key := range Flashcards {
		fmt.Println("Question:", key)
		fmt.Println("Please press \"y\" to see the answer or press \"n\" to skip:")
		var input string
		fmt.Scanln(&input)

		switch input {
		case "y":
			fmt.Println("Answer:", Flashcards[key])
		case "n":
			continue
		default:
			fmt.Printf("%s is not an option\n", input)
		}
	}
}

func main() {
	DisplayMainMenu()
	var choice string
	fmt.Scanln(&choice)
	MainMenuSelection(choice)
}
