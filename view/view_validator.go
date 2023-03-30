package view

import "reflect"

type validator struct {
}

func (validator) validateLetterColoring(l *LettersToColor) error {
	if l == nil {
		return emptyPointer
	}
	var x []PositionType
	if reflect.TypeOf(l.LetterPositions) != reflect.TypeOf(x) {
		return wrongSliceType
	}
	if len(l.GuessToColor) == 0 {
		return noGuess
	}
	if len(l.LetterPositions) == 0 {
		return noLetterArray
	}
	if len(l.GuessToColor) != len(l.LetterPositions) {
		return inEqualLength
	}
	for i := 0; i < len(l.LetterPositions); i++ {
		if l.LetterPositions[i] == 1 || l.LetterPositions[i] == 2 || l.LetterPositions[i] == 3 {
			return nil
		}
	}
	return invalidNumber
}
