// Package game is used to run the wordle game.
package game

import (
	"fmt"

	"git.ops.betable.com/go-betable-lib/errors"
)

// Giving the game a validator and inpt package dependency, I dont exactly know how this works but it does?
type Game struct {
	validator validatorInterface
	converter converterInterface
	inpt      inptInterface
	private   privateMethodInterface
	view      viewInterface
}

func New(input inptInterface, view viewInterface) *Game {
	game := &Game{
		validator: validator{},
		converter: converter{},
		inpt:      input,
		view:      view,
	}
	game.private = game

	return game
}

// RunTheGame receives a GameState and an input controller and carries out the necessary actions to play the game
func (c *Game) RunTheGame(g *GameState) error {

	//validating the public method inputs received.
	if validationErr := c.validator.validateGame(g); validationErr != nil {
		return errors.WrapError(errors.ValidationError, validationErr)
	}

	c.view.DisplayRules()

	win, err := c.private.playTheGameForSixRounds(g)
	if err != nil {

		return err
	}

	//If user lost
	if !win {
		fmt.Println("Well, you didn't get it, the answer was: " + g.Answer)
	}

	fmt.Println("If you would like to play another game, enter y on the command line, if not, then press any other button and get outta here ya rascal.")
	return err
}

// letterPositionWinCheckerAndSetter takes in a GameState and a guess as a string and returns whether the user has won or lost and how close this guess was to the answer.
func (c *Game) letterPositionWinCheckerAndSetter(g *GameState, guess string) (bool, []PositionType, error) {

	// letterPosition is an array of integers that represent the guess and each letter and its position, and each letters relation to the answer.
	letterPosition := []PositionType{NotInWord, NotInWord, NotInWord, NotInWord, NotInWord}

	if guess == g.Answer {
		fmt.Println("You got it!")
		win := true
		return win, []PositionType{RightPos, RightPos, RightPos, RightPos, RightPos}, nil
	}

	//Going through each letter of the guess and checking it against the answer.
	for j := 0; j < len(g.Answer); j++ {

		//If the guess and the answer have the same letter at the same position, the
		//corresponding letterPosition value is set to 3.
		if guess[j] == g.Answer[j] {
			letterPosition[j] = RightPos
		}
	}

	//These for loops allow the program to check each position of the letter and the answer against each other to see if a letter exists but not in the position is appears in the answer.
	for m := 0; m < len(g.Answer); m++ {
		for k := 0; k < len(g.Answer); k++ {

			/*
				If their is a similar letter in both the answer and the guess, and the position of the letter in the guess is not set to 3,
				and the letter hasn't already been correctly guessed at the position it appears in the answer, then set the letter position in the guess to 2.
			*/

			if guess[m] == g.Answer[k] && letterPosition[m] != RightPos && letterPosition[k] != RightPos {
				letterPosition[m] = InWordWrongPos
			}
		}
	}
	lose := false
	//returning whether the user has won or lost.
	return lose, letterPosition, nil
}

// isValidGuess checks the user input against the list of ValidGuesses held by the GameState and the list of previous guesses and returns an error if the guess is invalid.
func (c *Game) isValidGuess(g *GameState, guess string, listOfPreviousGuesses map[string]string) error {

	//Using the string input by the user as the lookup key for the map, if it exists
	//this will return true and if not, it will return false.
	if guess == listOfPreviousGuesses[guess] {
		fmt.Println("This has already been previously guessed, please guess again.")
		return answerAlreadyGuessed
	}
	if guess == g.ValidGuesses[guess] && guess != listOfPreviousGuesses[guess] {
		fmt.Println("This is a valid guess")
		return nil
	} else {
		fmt.Println("This is not a valid guess, please guess again")
		return invalidGuess
	}

}

// playTheGameForSixRounds runs the game for a full game (6 rounds.) It takes in a GameState. It will return a bool representing if the user has won or lost along with any error present.
func (c *Game) playTheGameForSixRounds(g *GameState) (bool, error) {
	win := true
	lose := false
	listOfPreviousGuesses := make(map[string]string)

	var err error
	//This is the game loop, where the 6 rounds take place
	for i := 0; i < len(g.ValidGuesses) && i < maxRounds; {

		//guess is set to the users guess which is checked in the GetValidWordInput function
		guess, err := c.inpt.GetValidWordInput()
		if err != nil {
			return lose, err
		}

		//isValidGuess will see if the users guess exists in the list of possible valid guesses while not having already been guessed.
		if err := c.private.isValidGuess(g, guess, listOfPreviousGuesses); err == nil {

			listOfPreviousGuesses = c.private.addToMap(listOfPreviousGuesses, guess)

			//winnerFlag is set to the boolean result of the LetterPositionSetter function where the guess is checked to see how right it is.
			winFlag, letterPositions, err := c.private.letterPositionWinCheckerAndSetter(g, guess)
			if err != nil {
				return lose, err
			}

			err = c.view.LetterColoring(c.converter.toCreateLetterColoring(letterPositions, guess))
			if err != nil {
				return lose, err
			}
			//if winnerFlag == true, then the user has won and the program will break out of the current games loop and prompt the user for if they would like to keep playing or not.
			if winFlag {
				fmt.Println("You won!")

				return win, err
			}

			//incrementing game progress by counting this guess as a valid guess
			i++
			fmt.Println("Next guess please:")
		}

	}

	return lose, err
}

// addToMap adds a string to an existing map and returns the new map
func (c *Game) addToMap(currentMap map[string]string, word string) map[string]string {
	currentMap[word] = word
	return currentMap
}
