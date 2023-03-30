package game

import "masonwordle/view"

type inptInterface interface {
	GetValidWordInput() (string, error)
}

type viewInterface interface {
	DisplayRules()
	LetterColoring(l *view.LettersToColor) error
}

type converterInterface interface {
	toCreateLetterColoring(letterPosition []PositionType, guess string) *view.LettersToColor
}

type validatorInterface interface {
	validateGuess(input string) error
	validateGame(g *GameState) error
}

type privateMethodInterface interface {
	letterPositionWinCheckerAndSetter(g *GameState, guess string) (bool, []PositionType, error)
	isValidGuess(g *GameState, input string, listOfPreviousGuesses map[string]string) error
	playTheGameForSixRounds(g *GameState) (bool, error)
	addToMap(currentMap map[string]string, word string) map[string]string
}
