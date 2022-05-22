//  * implement the command 'hint' - show to the user a random unguessed letter

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var inputReader = bufio.NewReader(os.Stdin)
var dictionary = []string{
	"Clifton",
	"Fresnaye",
	"Plattekloof",
	"Athlone",
	"Saltriver",
	"Rondebosch",
	"Claremont",
	"Durbanville",
	"Oranjezicht",
	"Vredehoek",
	"Plumstead",
	"Ottery",
	"Pinelands",
	"Tableview",
	"Kenilworth",
	"De Waterkant",
}

func main() {

	rand.Seed(time.Now().UnixNano())

	targetWord := randomWord()
	guessLetters := declareGuessedWords(targetWord)
	hangmanIndx := 0
	previousGuesses := make([]string, len(targetWord))

	fmt.Println("Hello! Welcome to Alex's Hangman Game!")
	fmt.Println(" ")
	fmt.Println("As this is a guessing game, the directions are simple.")
	fmt.Println("Just input the (lowercase) letter you believe comes next")
	fmt.Println("in the target word!")
	fmt.Println(" ")
	fmt.Println("Once the word is complete, you win! However, if you fail")
	fmt.Println("to input the correct letter 10x, game over...")
	fmt.Println("The words this game hands out are all various place-names")
	fmt.Println("from around the city of Cape Town (my city of residence).")
	fmt.Println("feel free to Google - should you need to - otherwise enjoy!")
	fmt.Println(" ")
	fmt.Println(" | ")
	fmt.Println(" | ")
	fmt.Println(" v ")
	fmt.Println("  ")



	for !isGameOver(targetWord, guessLetters, hangmanIndx) {
		renderGame(targetWord, guessLetters, hangmanIndx)
		input := readInput()
		if doesSliceContain(previousGuesses, input) {
			fmt.Println("Letter already used :\\")
			continue
		}
		previousGuesses = append(previousGuesses, input)

		if len(input) != 1 {
			fmt.Println("Please Enter Valid Input... (Single Letters Only)")
			continue
		}

		letter := rune(input[0])
		if correctGuessLetter(targetWord, letter) {
			guessLetters[letter] = true
		} else {
			fmt.Println("Wrong! Try again!")
			hangmanIndx++
		}

	}

	renderGame(targetWord, guessLetters, hangmanIndx)
	fmt.Print("Game Over... ")
	if isWordGuessed(targetWord, guessLetters) {
		fmt.Println("Winner!!! (^_^) ")
	} else if hangmanProgress(hangmanIndx) {
		fmt.Println("You lose, better luck next time (>_<)")
	} else {
		panic("Invalid Game State. Panic Executed. Game Over.")
	}

}

func randomWord() string {
	targetWord := dictionary[rand.Intn(len(dictionary))]
	return targetWord
}

func declareGuessedWords(targetWord string) map[rune]bool {
	guessLetters := map[rune]bool{}
	guessLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true

	return guessLetters
}

func isGameOver(
	targetWord string,
	guessLetters map[rune]bool,
	hangmanIndx int,
) bool {
	return isWordGuessed(targetWord, guessLetters) ||
		hangmanProgress(hangmanIndx)
}

func isWordGuessed(targetWord string, guessLetters map[rune]bool) bool {
	for _, ch := range targetWord {
		if !guessLetters[unicode.ToLower(ch)] {
			return false
		}
	}

	return true
}

func hangmanProgress(hangmanIndx int) bool {
	return hangmanIndx >= 10
}

func renderGame(
	targetWord string,
	guessLetters map[rune]bool,
	hangmanIndx int,
) {
	fmt.Println(getWordGuessingProgress(targetWord, guessLetters))
	fmt.Println()
	fmt.Println(getHangman(hangmanIndx))
}

func getWordGuessingProgress(
	targetWord string,
	guessLetters map[rune]bool,
) string {
	result := ""
	for _, ch := range targetWord {
		if ch == ' ' {
			result += " "
		} else if guessLetters[unicode.ToLower(ch)] {
			result += fmt.Sprintf("%c", ch)
		} else {
			result += "_"
		}

		result += " "
	}

	return result
}

func getHangman(hangmanIndx int) string {
	data, err := ioutil.ReadFile(
		fmt.Sprintf("states/hangman%d", hangmanIndx))
	if err != nil {
		panic(err)
	}

	return string(data)
}

func readInput() string {
	fmt.Print("> ")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(input)
}

func correctGuessLetter(targetWord string, letter rune) bool {
	return strings.ContainsRune(targetWord, letter)
}

func doesSliceContain(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
