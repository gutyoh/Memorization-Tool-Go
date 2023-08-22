package main

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
	AnswerPrompt          = "Answer:"
	QuestionPrompt        = "Question:"
	CurrentQuestionPrompt = "current question:"
	NewQuestionPrompt     = "please write a new question:"
	CurrentAnswerPrompt   = "current answer:"
	NewAnswerPrompt       = "please write a new answer:"
	NoFlashcardsMsg       = "There is no flashcard to practice!"

	PressForAnswer = "press \"y\" to see the answer:"
	PressToSkip    = "press \"n\" to skip:"
	PressToUpdate  = "press \"u\" to update:"
	PressToDelete  = "press \"d\" to delete the flashcard:"
	PressToEdit    = "press \"e\" to edit the flashcard:"
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

func (fs *FlashcardStore) UpdateFlashcard(fc *Flashcard, newQuestion, newAnswer string) {
	fc.Question = newQuestion
	fc.Answer = newAnswer
	fs.DB.Save(fc)
}

func (fs *FlashcardStore) DeleteFlashcard(fc *Flashcard) {
	fs.DB.Delete(fc)
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
	var question, answer string

	question = getValidInput(QuestionPrompt, mt.UI.Scanner)
	answer = getValidInput(AnswerPrompt, mt.UI.Scanner)

	mt.Store.CreateFlashcard(question, answer)
}

func (mt *MemorizationTool) UpdateOrDeleteFlashcard(flashcard *Flashcard) {
	fmt.Println(PressToDelete)
	fmt.Println(PressToEdit)
	var input string
	fmt.Scanln(&input)

	switch input {
	case "d":
		mt.Store.DeleteFlashcard(flashcard)
	case "e":
		mt.EditFlashcard(flashcard)
	default:
		fmt.Printf(InvalidOptionMsg, input)
	}
}

func (mt *MemorizationTool) ProcessFlashcard(flashcard *Flashcard) {
	mt.UI.DisplayFlashcardQuestion(flashcard)
	var input string
	fmt.Scanln(&input)

	switch input {
	case "y":
		fmt.Println(AnswerPrompt, flashcard.Answer)
	case "n":
		return
	case "u":
		mt.UpdateOrDeleteFlashcard(flashcard)
	default:
		fmt.Printf(InvalidOptionMsg, input)
	}
}

func (mt *MemorizationTool) PracticeFlashcards() {
	flashcards := mt.Store.RetrieveAllFlashcards()

	if len(flashcards) == 0 {
		fmt.Println(NoFlashcardsMsg)
		return
	}

	for _, flashcard := range flashcards {
		mt.ProcessFlashcard(&flashcard)
	}
}

func (mt *MemorizationTool) EditFlashcard(flashcard *Flashcard) {
	var newQuestion, newAnswer string

	fmt.Println(CurrentQuestionPrompt, flashcard.Question)
	fmt.Println(NewQuestionPrompt)
	newQuestion = getValidInput(QuestionPrompt, mt.UI.Scanner)

	fmt.Println(CurrentAnswerPrompt, flashcard.Answer)
	fmt.Println(NewAnswerPrompt)
	newAnswer = getValidInput(AnswerPrompt, mt.UI.Scanner)

	mt.Store.UpdateFlashcard(flashcard, newQuestion, newAnswer)
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
