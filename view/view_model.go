package view

import "git.ops.betable.com/go-betable-lib/errors"

const (
	red    = "\033[41m"
	reset  = "\033[0m"
	yellow = "\033[43m"
	green  = "\033[42m"
)

type PositionType int

const (
	NotInWord      PositionType = 1
	InWordWrongPos PositionType = 2
	RightPos       PositionType = 3
)

type LettersToColor struct {
	LetterPositions []PositionType
	GuessToColor    string
}

var (
	noGuess        = errors.NewError("the guess in LettersToColor was empty")
	noLetterArray  = errors.NewError("their is no array of numbers pass into LettersToColor")
	invalidNumber  = errors.NewError("the number provided in LetterPostions was not 1, 2, or 3")
	emptyPointer   = errors.NewError("the pointer to LettersToColor was empty")
	inEqualLength  = errors.NewError("the length of the GuessToColor and the LetterPositions was not equal")
	wrongSliceType = errors.NewError("the type of slice given for letterPositions was not of type PositionType")
)
