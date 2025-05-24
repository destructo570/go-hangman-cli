package constants

import "fmt"

func PrintHangmanStatus(remainingAttempts int, guessedWord string) {
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
