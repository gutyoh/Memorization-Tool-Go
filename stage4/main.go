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
	RequiredCorrectCount  = 3

	PressForAnswer = "press \"y\" to see the answer:"
	PressToSkip    = "press \"n\" to skip:"
	PressToUpdate  = "press \"u\" to update:"
	PressToDelete  = "press \"d\" to delete the flashcard:"
	PressToEdit    = "press \"e\" to edit the flashcard:"
	PressCorrect   = "press \"y\" if your answer is correct:"
	PressWrong     = "press \"n\" if your answer is wrong:"
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
	for len(input) <= 1 || !isASCII(input) {
		fmt.Println(prompt)
		scanner.Scan()
		input = scanner.Text()
	}
	return input
}

type Flashcard struct {
	gorm.Model
	Question     string
	Answer       string
	CorrectCount uint
}

//// 🚨 Attention 🚨: Tests will also work if students want to use a non-gorm.Model struct! ✅
//type Flashcard struct {
//	ID           uint `gorm:"primaryKey"`
//	Question     string
//	Answer       string
//	CorrectCount uint
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

func (fs *FlashcardStore) UpdateFlashcard(fc *Flashcard) {
	fs.DB.Save(fc)
}

func (fs *FlashcardStore) DeleteFlashcard(fc *Flashcard) {
	fs.DB.Delete(fc)
}

type UserInterface struct {
	Scanner *bufio.Scanner
}

func (*UserInterface) DisplayMenu(items ...string) {
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

func (ui *UserInterface) DisplayFlashcardAnswer(flashcard *Flashcard) {
	ui.DisplayMenu(AnswerPrompt+flashcard.Answer, PressCorrect, PressWrong)
}

type MemorizationTool struct {
	UI    UserInterface
	Store FlashcardStore
}

func NewMemorizationTool(db *gorm.DB) (*MemorizationTool, error) {
	if !db.Migrator().HasTable(&Flashcard{}) {
		err := db.Migrator().CreateTable(&Flashcard{})
		if err != nil {
			return nil, fmt.Errorf("failed to create flashcards table: %w", err)
		}
	}
	scanner := bufio.NewScanner(os.Stdin)
	return &MemorizationTool{
		Store: FlashcardStore{DB: db},
		UI:    UserInterface{Scanner: scanner},
	}, nil
}

func (mt *MemorizationTool) BuildFlashcard() {
	question := getValidInput(QuestionPrompt, mt.UI.Scanner)
	answer := getValidInput(AnswerPrompt, mt.UI.Scanner)
	mt.Store.CreateFlashcard(question, answer)
}

func (mt *MemorizationTool) ProcessCorrectAnswer(flashcard *Flashcard) {
	flashcard.CorrectCount++
	if flashcard.CorrectCount == RequiredCorrectCount {
		mt.Store.DeleteFlashcard(flashcard)
		flashcard.CorrectCount = 0
	} else {
		mt.Store.UpdateFlashcard(flashcard)
	}
}

func (mt *MemorizationTool) ProcessIncorrectAnswer(flashcard *Flashcard) {
	flashcard.CorrectCount = 0
	mt.Store.UpdateFlashcard(flashcard)
}

func (mt *MemorizationTool) ProcessAnswerInput(flashcard *Flashcard) {
	mt.UI.DisplayFlashcardAnswer(flashcard)
	var input string
	fmt.Scanln(&input)

	switch input {
	case "y":
		mt.ProcessCorrectAnswer(flashcard)
	case "n":
		mt.ProcessIncorrectAnswer(flashcard)
	default:
		fmt.Printf(InvalidOptionMsg, input)
	}
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
		mt.Store.UpdateFlashcard(flashcard)
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
		mt.ProcessAnswerInput(flashcard)
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
	fmt.Println(CurrentQuestionPrompt, flashcard.Question)
	fmt.Println(NewQuestionPrompt)
	newQuestion := getValidInput(QuestionPrompt, mt.UI.Scanner)
	flashcard.Question = newQuestion

	fmt.Println(CurrentAnswerPrompt, flashcard.Answer)
	fmt.Println(NewAnswerPrompt)
	newAnswer := getValidInput(AnswerPrompt, mt.UI.Scanner)
	flashcard.Answer = newAnswer
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

func main() {
	db, err := gorm.Open(sqlite.Open(DatabaseName), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to open %s: %v", DatabaseName, err)
	}

	mt, err := NewMemorizationTool(db)
	if err != nil {
		log.Fatalf("failed to initialize the application: %v", err)
	}

	mt.MainMenuSelection()
}
