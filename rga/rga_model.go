package rga

import "git.ops.betable.com/go-betable-lib/errors"

const (
	answerlink = "https://gist.githubusercontent.com/cfreshman/a03ef2cba789d8cf00c08f767e0fad7b/raw/c915fa3264be6d35990d0edb8bf927df7a015602/wordle-answers-alphabetical.txt"
	guesslink  = "https://gist.github.com/cfreshman/40608e78e83eb4e1d60b285eb7e9732f/raw/wordle-nyt-allowed-guesses.txt"

	answersBackupPath = "rga/wordle-answers-alphabetical.txt"

	guessesBackupPath = "rga/wordle-allowed-guesses.txt"
)

var (
	backUpFilePathEmpty = errors.NewError("path cannot be empty")
	responseCodeErr     = errors.NewError("error with response code after downloading github page")
)
