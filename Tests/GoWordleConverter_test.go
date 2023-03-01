package main_test

import (
	"fmt"
	"log"
	"masonwordle/game"
	"testing"
	// . "github.com/onsi/ginkgo/v2"
	// . "github.com/onsi/gomega"
)

func TestGetGuessesAndAnswers(t *testing.T) {

	guesses, answers := game.GetGuessesAndAnswers()
	for x := 0; x < len(guesses); x++ {
		if len(guesses[x]) != 5 {
			t.Fatalf("One of the guesses on the list of valid guesses is longer than 5 letters")
		}
	}

	for x := 0; x < len(answers); x++ {
		if len(answers[x]) != 5 {
			t.Fatalf("One of the answers on the list of valid answers is longer than 5 letters")
		}
	}
}

func TestMapIt(t *testing.T) {
	guesses, answers := game.GetGuessesAndAnswers()

	guessesMapped, _ := game.MapIt(guesses)
	answersMapped, _ := game.MapIt(answers)
	for i := 0; i < len(guesses); i++ {

		if guesses[i] != guessesMapped[guesses[i]] {
			fmt.Println("MapIt did not work on guesses")
			t.Fatalf("Mapit was fatal")
			break
		}
	}

	for j := 0; j < len(answers); j++ {
		if answers[j] != answersMapped[guesses[j]] {
			fmt.Println("MapIt did not work on answers")
			t.Fatalf("Mapit was fatal")
			break
		}
	}

	incorrectText := []string{"bob", "jean", "duggan", "weheh", "beans"}

	_, err := game.MapIt(incorrectText)

	if err == nil {
		t.Fatalf("Incorrect text worked.")
	}

}

func TestGetAnswer(t *testing.T) {
	var x []string
	v := game.GetAnswer(x)

	if v != "failure" {
		fmt.Println("GetAnswer did not work")
		t.Fatalf("GetAnswer was fatal")
	}

	b := []string{"hello", "swell", "sewer", "river"}
	n := game.GetAnswer(b)
	if n != "hello" && n != "swell" && n != "sewer" && n != "river" {
		fmt.Println("GetAnswer did not work")
		t.Fatalf("GetAnswer was fatal")
	}
}

func TestGetText(t *testing.T) {

	answers, err := game.GetText("https://gist.github.com/cfreshman/a7b776506c73284511034e63af1017ee/raw/wordle-nyt-answers-alphabetical.txt")
	if err != nil {
		log.Fatal(err)
	}

	guesses, err := game.GetText("https://gist.github.com/cfreshman/40608e78e83eb4e1d60b285eb7e9732f/raw/wordle-nyt-allowed-guesses.txt")
	if err != nil {
		log.Fatal(err)
	}

	incorrectAnswers, err := game.GetText("https://gist.github.com/cfreshman/a7b776506c73284511034e63af1017ee/raw/wordle-nyt-answers-alphatical.txt")
	if err == nil {
		log.Fatal(err)
	}

	incorrectGuesses, err := game.GetText("https://gist.github.com/cfreshman/40608e78e83eb4e1d60b285eb7e9732f/raw/wordle-nyt-allowed-gues.txt")
	if err == nil {
		log.Fatal(err)
	}

	_ = answers
	_ = guesses
	_ = incorrectAnswers
	_ = incorrectGuesses
}

// // Commented out because I dont really know how to test for input when input is already checked
// // in the function.
// func TestInput(t *testing.T) {
// 	testInput := game.Input("regular input")

// 	if len(testInput) > 5 {
// 		t.Fatalf("One of the test input is longer than 5 letters")
// 	}
// }
