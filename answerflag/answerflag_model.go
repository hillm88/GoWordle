package answerflag

import "git.ops.betable.com/go-betable-lib/errors"

var (
	validListLength = 1

	answerListEmpty      = errors.NewError("AnswerList cannot be empty")
	validAnswerListEmpty = errors.NewError("ValidAnswers cannot be empty")
)

type AnswerInfoHolder struct {
	AnswerList     []string
	ValidAnswers   map[string]string
	ShowAnswerFlag bool
	SetAnswerFlag  bool
}
