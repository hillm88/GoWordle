package answerflag

type validator struct {
}

// validateAnswerLists takes in a gamestate and checks that the ValidAnswers and the AnswerList have at least one item in them.
func (validator) validateAnswerLists(a *AnswerInfoHolder) error {
	if len(a.ValidAnswers) == 0 {
		return validAnswerListEmpty
	}
	if len(a.AnswerList) == 0 {
		return answerListEmpty
	}
	return nil
}
