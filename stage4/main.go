package main

/*
[Memorization Tool - Stage 4/4: The Leitner system](https://hyperskill.org/projects/159/stages/829/implement)
-------------------------------------------------------------------------------
[Computer algorithms](https://hyperskill.org/learn/step/16547)
*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

	CorrectCount int

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
		fmt.Printf("press \"y\" to see the answer:\n")
		fmt.Printf("press \"n\" to skip:\n")
		fmt.Printf("press \"u\" to update:\n")

		var input string
		fmt.Scanln(&input)

		switch input {
		case "y":
			fmt.Printf("Answer: %s\n", flashcard.Answer)
			fmt.Printf("press \"y\" if your answer is correct:\n")
			fmt.Printf("press \"n\" if your answer is wrong:\n")

			fmt.Scanln(&input)

			switch input {
			case "y":
				mt.CorrectCount += 1
				if mt.CorrectCount == 3 {
					mt.DB.Delete(&flashcard)
					mt.CorrectCount = 0
				}
			case "n":
				continue
			default:
				fmt.Printf("%s is not an option\n", input)
			}

		case "n":
			continue

		case "u":
			fmt.Printf("press \"d\" to delete the flashcard:\n")
			fmt.Printf("press \"e\" to edit the flashcard:\n")

			fmt.Scanln(&input)

			switch input {
			case "d":
				mt.DB.Delete(&flashcard)

			case "e":
				fmt.Printf("current question: %s\n", flashcard.Question)
				fmt.Printf("please write a new question:\n")

				scanner := bufio.NewScanner(os.Stdin)
				var newQuestion string
				for len(newQuestion) == 1 || !isASCII(newQuestion) || newQuestion == "" {
					scanner.Scan()
					newQuestion = scanner.Text()
				}
				flashcard.Question = newQuestion
				mt.DB.Save(&flashcard)

				fmt.Printf("current answer: %s\n", flashcard.Answer)
				fmt.Printf("please write a new answer:\n")

				var newAnswer string
				for len(newAnswer) == 1 || !isASCII(newAnswer) || newAnswer == "" {
					scanner.Scan()
					newAnswer = scanner.Text()
				}
				flashcard.Answer = newAnswer
				mt.DB.Save(&flashcard)

			default:
				fmt.Printf("%s is not an option\n", input)
			}
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
