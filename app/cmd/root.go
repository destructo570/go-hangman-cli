package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type promptContent struct {
	errorMsg string
	label    string
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "A simple hangman game built in Go lang",
	Long:  `A simple hangman game built in Go lang`,

	Run: func(cmd *cobra.Command, args []string) {
		triggerInit()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func triggerInit() {
	selectedOption := promptGetInput()
	if selectedOption == "1) Start game" {
		startGame()
	} else {
		os.Exit(1)
	}

}

func promptGetInput() string {

	prompt := promptui.Select{
		Label: "Select One Option",
		Items: []string{"1) Start game", "2) Exit"},
	}
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return result
}

func promptGetGuess() string {
	validate := func(input string) error {
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Guess a character",
		Validate: validate,
	}
	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	// println("Returning")
	return result
}

func getRandomWord() string {
	words := []string{"shark", "dolphin", "elephant"}
	randomNum := rand.Intn(len(words))
	return words[randomNum]
}

func getGuessCount(word string) (int, map[string]bool) {
	mySet := make(map[string]bool)
	for _, char := range word {
		_, present := mySet[string(char)]
		if !present {
			mySet[string(char)] = true
		}
	}
	return len(mySet), mySet
}

func getGuessedWord(word string, currentGuess string, input string) string {
	str := []rune(currentGuess)

	for i, char := range word {
		if string(char) == input {
			str[i] = char
		}
	}
	return string(str)
}

func printHangmanStatus(remainingAttempts int, guessedWord string) {
	var hangmanStages = []string{
		`
   +---+
   |   |
       |
       |
       |
       |
=========`,
		`
   +---+
   |   |
   O   |
       |
       |
       |
=========`,
		`
   +---+
   |   |
   O   |
   |   |
       |
       |
=========`,
		`
   +---+
   |   |
   O   |
  /|   |
       |
       |
=========`,
		`
   +---+
   |   |
   O   |
  /|\  |
       |
       |
=========`,
		`
   +---+
   |   |
   O   |
  /|\  |
  /    |
       |
=========`,
		`
   +---+
   |   |
   O   |
  /|\  |
  / \  |
       |
=========`,
	}

	fmt.Println()
	fmt.Println("*****************")
	fmt.Printf("Remaining lives: %d\n", remainingAttempts)
	fmt.Println("*****************")
	fmt.Printf("Word: %s", guessedWord)
	fmt.Println()
	fmt.Println()
	fmt.Println(hangmanStages[(len(hangmanStages)-1)-remainingAttempts])
	fmt.Println()

}

func startGame() {
	word := getRandomWord()
	_, guessSet := getGuessCount(word)
	lives := 6
	guessWord := strings.Repeat("_", len(word))

	for {
		printHangmanStatus(lives, guessWord)
		if len(guessSet) == 0 {
			fmt.Println("You Won! :)")
			os.Exit(1)

		} else if lives == 0 {
			fmt.Println("You Lost! :(")
			os.Exit(1)

		} else {
			userInput := promptGetGuess()

			_, present := guessSet[userInput]
			if present {
				guessWord = getGuessedWord(word, guessWord, userInput)
				delete(guessSet, userInput)
			} else {
				lives--
			}
		}
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
