package view

type validatorInterface interface {
	validateLetterColoring(l *LettersToColor) error
}
