package answerflag

type privateMethodInterface interface {
	getAnswer(answerList []string) (string, error)
	setAnswer(validAnswers map[string]string) (string, error)
	isValidListSize(aList []string) bool
}

type validatorInterface interface {
	validateAnswerLists(a *AnswerInfoHolder) error
}

type inptInterface interface {
	GetValidWordInput() (string, error)
}
