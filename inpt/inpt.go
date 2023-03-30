// Package inpt is used to get a users input.
package inpt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Inpt struct {
	private privateMethodInterface
}

func New() *Inpt {
	inpt := &Inpt{}
	inpt.private = inpt
	return inpt
}

// GameContinueDecision gets the users input on if they would like to continue the game.
// The users input is checked against the continue game input of "y" and returns a bool based on this along with any errors.
func (i *Inpt) GameContinueDecision() (bool, error) {

	discontinueGame := false
	var decisionChar string
	var err error

	decisionChar, err = i.private.getInput()
	if err != nil {
		return discontinueGame, err
	}

	return decisionChar == "y", err
}

// GetValidWordInput gets the users input and validates and formats it and returns it along with any errors
func (i *Inpt) GetValidWordInput() (string, error) {
	input := ""
	var err error

	// Runs until the user's input is valid or an error happens.
	for {
		input, err = i.private.getInput()
		if err != nil {
			return "", err
		}

		if i.private.isValidInput(input, validWordLength) {
			break
		}

		fmt.Println("Invalid input, please try again")
	}

	input = input[0:5]
	return input, nil
}

// getInput gets the users input and returns it as a lowercase string
func (i *Inpt) getInput() (string, error) {
	// word is a string that will hold the guess for the user initially.
	input, err := i.private.retrieveCLIInput()

	if err != nil {
		return "", err

	}

	//Trimming new line and converting to lowercase
	input = strings.TrimSuffix(strings.ToLower(input), "\n")

	// Return the users input.
	return input, nil
}

// isValidInput checks the length of the input against the exact length the input can be
func (i *Inpt) isValidInput(input string, exactLength int) bool {

	return len(input) == exactLength
}

func (i *Inpt) retrieveCLIInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	return reader.ReadString('\n')
}
