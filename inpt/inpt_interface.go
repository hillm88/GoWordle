package inpt

type privateMethodInterface interface {
	getInput() (string, error)
	isValidInput(input string, exactLength int) bool
	retrieveCLIInput() (string, error)
}
