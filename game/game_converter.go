package game

import "masonwordle/view"

type converter struct{}

func (converter) toCreateLetterColoring(letterPosition []PositionType, guess string) *view.LettersToColor {
	var viewLetterPosition []view.PositionType
	for _, v := range letterPosition {
		viewLetterPosition = append(viewLetterPosition, view.PositionType(v))
	}

	return &view.LettersToColor{
		LetterPositions: viewLetterPosition,
		GuessToColor:    guess,
	}
}
