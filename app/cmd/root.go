package cmd

import (
	"fmt"
	"hangman/app/constants"
	"math/rand"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "A simple hangman game built in Go lang",
	Long:  `A simple hangman game built in Go lang`,

	Run: func(cmd *cobra.Command, args []string) {
		startGame()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
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
	return result
}

func getRandomWord() string {
	words := constants.WordsList
	randomNum := rand.Intn(len(words))
	return words[randomNum]
}

func getGuessCount(word string) map[string]bool {
	mySet := make(map[string]bool)
	for _, char := range word {
		_, present := mySet[string(char)]
		if !present {
			mySet[string(char)] = true
		}
	}
	return mySet
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

func startGame() {
	word := getRandomWord()
	guessSet := getGuessCount(word)
	lives := 6
	guessWord := strings.Repeat("_", len(word))

	for {
		constants.PrintHangmanStatus(lives, guessWord)
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
