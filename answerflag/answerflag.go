// Package answerflag is meant to handle any of the incoming flags for the gowordle game that have to do with the answer.
package answerflag

import (
	"fmt"
	"math/rand"
	"time"

	"git.ops.betable.com/go-betable-lib/errors"
)

type Flag struct {
	validator validatorInterface
	inpt      inptInterface
	private   privateMethodInterface
}

func New(input inptInterface) *Flag {

	flag := &Flag{
		validator: validator{},
		inpt:      input,
	}
	flag.private = flag
	return flag
}

// AnswerFlagHandler takes in a AnswerInfoHolder struct pointer, and from it, it will either allow the user to set the answer or it will pick a random answer from the ValidAnswers list inside of the AnswerInfoHolder struct. It will also show the answer if the ShowAnswerFlag has been set to true. It will return the answer and an error.
func (c *Flag) AnswerFlagHandler(a *AnswerInfoHolder) (string, error) {
	//I chose not to validate the answer flags because they can only be true or false
	if validationErr := c.validator.validateAnswerLists(a); validationErr != nil {
		return "", errors.WrapError(errors.ValidationError, validationErr)
	}

	var answer string
	var err error
	if a.SetAnswerFlag {

		answer, err = c.private.setAnswer(a.ValidAnswers)
		if err != nil {
			return "", err
		}
	} else {
		answer, err = c.private.getAnswer(a.AnswerList)
		if err != nil {
			return "", err
		}
	}

	if a.ShowAnswerFlag {
		fmt.Println()
		fmt.Println("The answer is: ", answer)
		fmt.Println()

	}
	return answer, err

}

// getAnswer takes in a string array of all the possible answers and then it randomly chooses one.
func (c *Flag) getAnswer(answerList []string) (string, error) {
	// Checking that the list isn't too short aka there has to be at least 1 word in the list.
	if c.private.isValidListSize(answerList) {
		// length holds the length of the inputted string array
		length := len(answerList)
		//Generating a new random seed
		rand.Seed(time.Now().UnixNano())
		//Generating a random number from 0 to the length of the array - 1
		//Using Intn means that you only have to give the upper bound and the lower bound will be set at 0.
		randomNum := rand.Intn(length)

		//returning the randomNum position of the answer list, A.K.A the answer the user will now be looking for.
		return answerList[randomNum], nil
	} else {
		return "", answerListEmpty
	}
}

// setAnswer takes in a list of valid answers and allows the user to input one of them.
func (c *Flag) setAnswer(validAnswers map[string]string) (string, error) {
	fmt.Println("Set the answer: ")

	//Setting the answer equal to user input
	answer, err := c.inpt.GetValidWordInput()
	if err != nil {
		return "", err
	}
	//If the answer given by the user is not a valid answer, then prompt the user for another answer
	for !(answer == validAnswers[answer]) {
		fmt.Println("Answer doesn't match the list of accepted answers, please enter a valid answer:")
		answer, err = c.inpt.GetValidWordInput()
		if err != nil {
			return "", err
		}
	}

	//Returning the answer.
	return answer, err
}

// isValidListSize checks that the list size is greater than or equal to one
func (c *Flag) isValidListSize(aList []string) bool {

	return len(aList) >= validListLength
}
