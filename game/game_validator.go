package game

type validator struct {
}

// validateGuess This function checks if the users input string is empty.
func (validator) validateGuess(input string) error {

	//Using the string input by the user as the lookup key for the map, if it exists
	//this will return true and if not, it will return false.

	if len(input) != 5 {
		return invalidGuess
	}
	return nil
}

// validateGame Takes in a gamestate and checks if any of its fields are empty.
func (validator) validateGame(g *GameState) error {
	if g.Answer == "" {
		return answerEmpty
	}

	if len(g.ValidGuesses) == 0 {
		return validGuessListEmpty
	}

	if len(g.AnswerList) == 0 {
		return answerListEmpty
	}

	if len(g.ValidAnswers) == 0 {
		return validAnswerListEmpty
	}

	return nil

}
