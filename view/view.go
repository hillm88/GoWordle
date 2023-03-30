// Package view displays most of the terminal related output. This includes the initial rule set and the guess of the user colorized.
package view

import (
	"fmt"
)

type View struct {
	validator validatorInterface
}

func New() *View {
	View := &View{
		validator: validator{},
	}

	return View
}

// view prints out all the necessary rules and text that is present at the start of a game of wordle.
func (v *View) DisplayRules() {
	fmt.Println()
	fmt.Println("Welcome to Wordle by the best programmer in the world, Mason 'The best of all time' Hill")
	fmt.Println("Here are the rules:")
	fmt.Println("You get 6 guesses to get the word.")
	fmt.Println("You have to guess an actual 5 letter word, no guesses like 'aabbc'")
	fmt.Println("You cannot guess the same word twice.")
	fmt.Println("You will be able to see if a letter is in the right position or not using the following rubric:")
	fmt.Println("Red means that the letter is not in the word at all")
	fmt.Println("Yellow means that the letter is in the word but not in the right place")
	fmt.Println("Green means that the letter is in the right place")
	fmt.Println("Get it?")
	fmt.Println("Got it?")
	fmt.Println("Good.")
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++")

	fmt.Println("***********************************************")
	fmt.Println("The answer has been set and the game is now ready, so type in a 5 letter word on the command line and lets begin!")

}

// LetterColoring takes in an int array of the positions of the letters in the guess and whether they are in the word or not and compares
// them to the enum representing each of those scenarios. The guess letters are then printed with their corresponding background color.
func (v *View) LetterColoring(l *LettersToColor) error {

	if err := v.validator.validateLetterColoring(l); err != nil {
		fmt.Println("error", err)
		return err
	}
	for x := range l.LetterPositions {

		if l.LetterPositions[x] == NotInWord {
			fmt.Print(red + string(l.GuessToColor[x]) + reset)
		}
		if l.LetterPositions[x] == InWordWrongPos {
			fmt.Print(yellow + string(l.GuessToColor[x]) + reset)
		}
		if l.LetterPositions[x] == RightPos {
			fmt.Print(green + string(l.GuessToColor[x]) + reset)
		}
	}
	fmt.Println()

	return nil
}
