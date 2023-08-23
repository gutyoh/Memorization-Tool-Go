package main

/*
[Memorization Tool - Stage 2/4: Store the flashcards](https://hyperskill.org/projects/159/stages/827/implement)
-------------------------------------------------------------------------------
[Intro to computational thinking](https://hyperskill.org/learn/step/8742)
[Components of computational thinking](https://hyperskill.org/learn/step/8745)
[Design principles](https://hyperskill.org/learn/step/8956)
[Single Responsibility Principle](https://hyperskill.org/learn/step/8963)
[Variadic functions](https://hyperskill.org/learn/step/21798)
[Structs](https://hyperskill.org/learn/topic/1891)
[Methods](https://hyperskill.org/learn/topic/1928)
[Debugging Go code in GoLand](https://hyperskill.org/learn/step/23118)
[Introduction to GORM](https://hyperskill.org/learn/step/20695)
[Declaring GORM Models](https://hyperskill.org/learn/step/28639)
[Migrations](https://hyperskill.org/learn/step/22043)
[CRUD Operations — Create](https://hyperskill.org/learn/step/22859)
[CRUD Operations — Read](https://hyperskill.org/learn/step/24151)
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

// SQLite Database name
const DatabaseName = "flashcard.db"

// General prompts and messages
const (
	InvalidOptionMsg = "%s is not an option\n"
	GoodbyeMsg       = "Bye!"
)

// Flashcard specific prompts and settings
const (
	AnswerPrompt    = "Answer:"
	QuestionPrompt  = "Question:"
	NoFlashcardsMsg = "There is no flashcard to practice!"

	PromptViewOrSkipAnswer = "Please press \"y\" to see the answer or press \"n\" to skip:"

	PressForAnswer = "press \"y\" to see the answer:"
	PressToSkip    = "press \"n\" to skip:"
	PressToUpdate  = "press \"u\" to update:"
)

// Main menu options
const (
	MainMenuAddFlashcards      = "1. Add flashcards"
	MainMenuPracticeFlashcards = "2. Practice flashcards"
	MainMenuExit               = "3. Exit"
)

// Flashcard menu options
const (
	FlashcardMenuAddNew       = "1. Add a new flashcard"
	FlashcardMenuReturnToMain = "2. Exit"
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

type Flashcard struct {
	gorm.Model
	Question string
	Answer   string
}

//// ⚠️ Tests will also pass with a non-gorm.Model struct! ⚠️
//type Flashcard struct {
//	ID       uint `gorm:"primaryKey"`
//	Question string
//	Answer   string
//}

type FlashcardStore struct {
	DB *gorm.DB
}

func (fs *FlashcardStore) CreateFlashcard(question, answer string) {
	fs.DB.Create(&Flashcard{Question: question, Answer: answer})
}

func (fs *FlashcardStore) RetrieveAllFlashcards() []Flashcard {
	var flashcards []Flashcard
	fs.DB.Find(&flashcards)
	return flashcards
}

type UserInterface struct {
	Scanner *bufio.Scanner
}

func (_ *UserInterface) DisplayMenu(items ...string) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func (ui *UserInterface) DisplayMainMenu() {
	ui.DisplayMenu(MainMenuAddFlashcards, MainMenuPracticeFlashcards, MainMenuExit)
}

func (ui *UserInterface) DisplayFlashcardMenu() {
	ui.DisplayMenu(FlashcardMenuAddNew, FlashcardMenuReturnToMain)
}

func (ui *UserInterface) DisplayFlashcardQuestion(flashcard *Flashcard) {
	ui.DisplayMenu(QuestionPrompt+flashcard.Question, PressForAnswer, PressToSkip, PressToUpdate)
}

type MemorizationTool struct {
	UI    UserInterface
	Store FlashcardStore
}

func (mt *MemorizationTool) BuildFlashcard() {
	scanner := bufio.NewScanner(os.Stdin)
	var question, answer string

	question = getValidInput(QuestionPrompt, scanner)

	answer = getValidInput(AnswerPrompt, scanner)

	mt.Store.CreateFlashcard(question, answer)
}

func (mt *MemorizationTool) PracticeFlashcards() {
	flashcards := mt.Store.RetrieveAllFlashcards()

	if len(flashcards) == 0 {
		fmt.Println(NoFlashcardsMsg)
		return
	}

	for _, flashcard := range flashcards {
		fmt.Println(QuestionPrompt, flashcard.Question)
		fmt.Println(PromptViewOrSkipAnswer)
		fmt.Println()

		var input string
		fmt.Scanln(&input)

		switch input {
		case "y":
			fmt.Println(AnswerPrompt, flashcard.Answer)
		case "n":
			continue
		default:
			fmt.Printf(InvalidOptionMsg, input)
		}
	}
}

func (mt *MemorizationTool) FlashcardMenuSelection() {
	for {
		mt.UI.DisplayFlashcardMenu()
		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			mt.BuildFlashcard()
		case "2":
			return
		default:
			fmt.Printf(InvalidOptionMsg, choice)
		}
	}
}

func (mt *MemorizationTool) MainMenuSelection() {
	for {
		mt.UI.DisplayMainMenu()
		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			mt.FlashcardMenuSelection()
		case "2":
			mt.PracticeFlashcards()
		case "3":
			fmt.Println(GoodbyeMsg)
			return
		default:
			fmt.Printf(InvalidOptionMsg, choice)
		}
	}
}

func (mt *MemorizationTool) Initialize() {
	db, err := gorm.Open(sqlite.Open(DatabaseName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if !db.Migrator().HasTable(&Flashcard{}) {
		err = db.Migrator().CreateTable(&Flashcard{})
		if err != nil {
			log.Fatal(err)
		}
	}

	mt.Store = FlashcardStore{DB: db}
	mt.UI = UserInterface{Scanner: bufio.NewScanner(os.Stdin)}
}

func main() {
	var mt MemorizationTool
	mt.Initialize()
	mt.MainMenuSelection()
}
