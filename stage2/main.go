package main

/*
[Memorization Tool - Stage 2/4: Store the flashcards](https://hyperskill.org/projects/159/stages/827/implement)
-------------------------------------------------------------------------------
[Intro to computational thinking](https://hyperskill.org/learn/step/8742)
[Components of computational thinking](https://hyperskill.org/learn/step/8745)
[Design principles](https://hyperskill.org/learn/step/8956)
[Single Responsibility Principle](https://hyperskill.org/learn/step/8963)
[Function decomposition](https://hyperskill.org/learn/topic/1893)
[Structs](https://hyperskill.org/learn/topic/1891)
[Methods](https://hyperskill.org/learn/topic/1928)
[Debugging Go code in GoLand](https://hyperskill.org/learn/step/23118)
[Introduction to GORM](https://hyperskill.org/learn/step/20695)
[Migrations](https://hyperskill.org/learn/step/22043)
[Declaring GORM Models] — TODO
[CRUD Operations — Create](https://hyperskill.org/learn/step/22859)
[CRUD Operations — Read](https://hyperskill.org/learn/step/24151)
*/

import (
	"bufio"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"unicode"
)

type Flashcard struct {
	ID       uint `gorm:"primaryKey"`
	Question string
	Answer   string
}

// ====== HELPER FUNCTION =====

func isASCII(s string) bool {
	for _, c := range s {
		if c > unicode.MaxASCII {
			return false
		}
	}
	return true
}

type MemorizationTool struct {
	DB         *gorm.DB
	Flashcards []Flashcard

	MainMenu      [3]string
	FlashcardMenu [2]string
}

func (mt *MemorizationTool) DisplayFlashcardMenu() {
	for _, elem := range mt.FlashcardMenu {
		fmt.Printf("%s\n", elem)
	}
}

func (mt *MemorizationTool) DisplayMainMenu() {
	for _, elem := range mt.MainMenu {
		fmt.Printf("%s\n", elem)
	}
}

func (mt *MemorizationTool) BuildFlashcard() {
	var question, answer string
	scanner := bufio.NewScanner(os.Stdin)

	for len(question) == 1 || !isASCII(question) || question == "" {
		fmt.Printf("Question:\n")
		scanner.Scan()
		question = scanner.Text()
	}

	for len(answer) == 1 || !isASCII(answer) || answer == "" {
		fmt.Printf("Answer:\n")
		scanner.Scan()
		answer = scanner.Text()
	}

	mt.DB.Create(&Flashcard{Question: question, Answer: answer})

	var mainMenuChoice string
	mt.DisplayFlashcardMenu()
	fmt.Scanln(&mainMenuChoice)
	mt.FlashcardMenuSelection(mainMenuChoice)
}

func (mt *MemorizationTool) MainMenuSelection(choice string) {
	var mainMenuChoice string
	switch choice {
	case "1":
		mt.DisplayFlashcardMenu()
		fmt.Scanln(&mainMenuChoice)
		mt.FlashcardMenuSelection(mainMenuChoice)
	case "2":
		mt.PracticeFlashcards()
	case "3":
		fmt.Printf("Bye!\n")
		os.Exit(0)
	default:
		fmt.Printf("%s is not an option\n", choice)
		mt.DisplayMainMenu()
		fmt.Scanln(&choice)
		mt.MainMenuSelection(choice)
	}
}

func (mt *MemorizationTool) FlashcardMenuSelection(choice string) {
	switch choice {
	case "1":
		mt.BuildFlashcard()
	case "2":
		mt.DisplayMainMenu()
		var mainMenuChoice string
		fmt.Scanln(&mainMenuChoice)
		mt.MainMenuSelection(mainMenuChoice)
	default:
		fmt.Printf("%s is not an option\n", choice)
		mt.DisplayFlashcardMenu()
		fmt.Scanln(&choice)
		mt.FlashcardMenuSelection(choice)
	}
}

func (mt *MemorizationTool) PracticeFlashcards() {
	mt.DB.Find(&mt.Flashcards)

	if len(mt.Flashcards) == 0 {
		fmt.Printf("There is no flashcard to practice!\n")
	}

	if len(mt.Flashcards) > 0 {
		mt.DisplayPracticeMenu()
	}

	mt.DisplayMainMenu()
	var choice string
	fmt.Scanln(&choice)
	mt.MainMenuSelection(choice)
}

func (mt *MemorizationTool) DisplayPracticeMenu() {
	for _, flashcard := range mt.Flashcards {
		fmt.Printf("Question: %s\n", flashcard.Question)
		fmt.Printf("Please press \"y\" to see the answer or press \"n\" to skip:\n")

		var input string
		fmt.Scanln(&input)

		switch input {
		case "y":
			fmt.Printf("Answer: %s\n", flashcard.Answer)

		case "n":
			continue

		default:
			fmt.Printf("%s is not an option\n", input)
		}
	}
}

func main() {
	db, err := gorm.Open(sqlite.Open("flashcard.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	var mt = MemorizationTool{
		MainMenu:      [3]string{"1. Add flashcards", "2. Practice flashcards", "3. Exit"},
		FlashcardMenu: [2]string{"1. Add a new flashcard", "2. Exit"},
		DB:            db,
	}

	if !mt.DB.Migrator().HasTable(&Flashcard{}) {
		err = mt.DB.Migrator().CreateTable(&Flashcard{})
		if err != nil {
			log.Fatal(err)
		}
	}

	mt.DisplayMainMenu()
	var choice string
	fmt.Scanln(&choice)
	mt.MainMenuSelection(choice)
}
