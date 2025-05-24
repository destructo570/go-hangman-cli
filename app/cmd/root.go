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
	var words = []string{"shark", "dolphin", "elephant", "octopus"}
	return words[rand.Intn(len(words))]
}

func getWordSet(word string) map[string]bool {
	m := make(map[string]bool)
	for _, char := range word {
		m[string(char)] = true
	}
	return m
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
	wordSet := getWordSet(word)
	lives := 6
	guessWord := strings.Repeat("_", len(word))

	for {
		constants.PrintHangmanStatus(lives, guessWord)
		if len(wordSet) == 0 {
			fmt.Println("You Won! :)")
			os.Exit(1)
		} else if lives == 0 {
			fmt.Println("You Lost! :(")
			os.Exit(1)
		} else {
			userInput := promptGetGuess()
			_, present := wordSet[userInput]
			if present {
				guessWord = getGuessedWord(word, guessWord, userInput)
				delete(wordSet, userInput)
			} else {
				lives--
			}
		}
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
