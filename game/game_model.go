package game

import "git.ops.betable.com/go-betable-lib/errors"

type GameState struct {
	Answer       string
	ValidGuesses map[string]string
	AnswerList   []string
	ValidAnswers map[string]string
}

const (
	maxRounds       = 6
	validWordLength = 5
	validCharLength = 1
)

type PositionType int

const (
	NotInWord      PositionType = 1
	InWordWrongPos PositionType = 2
	RightPos       PositionType = 3
)

var (
	answerAlreadyGuessed = errors.NewError("cannot guess the same answer twice")
	invalidGuess         = errors.NewError("not a valid guess")
	answerEmpty          = errors.NewError("Answer cannot be empty")
	validGuessListEmpty  = errors.NewError("ValidGuesses cannot be empty")
	answerListEmpty      = errors.NewError("AnswerList cannot be empty")
	validAnswerListEmpty = errors.NewError("ValidAnswers cannot be empty")
	// isValidGuessNotWorking = errors.NewError("isValidGuess method is not working")
)
