package game

import (
	"fmt"
)

// View prints out all the necceasry rules and text that is present at the start of a game of wordle.
func View() {
	fmt.Println("Welcome to Wordle by the best programmer in the world, Mason 'The best of all time' Hill")
	fmt.Println("Here are the rules:")
	fmt.Println("You get 6 guesses to get the word.")
	fmt.Println("You have to guess an actual 5 letter word, no guesses like 'aabbc'")
	fmt.Println("You will be able to see if a letter is in the right position or not using the following rubric:")
	fmt.Println("1 Means that the letter is not in the word at all")
	fmt.Println("2 Means that the letter is in the word but not in the right place")
	fmt.Println("3 Means that the letter is in the right place")
	fmt.Println("Get it?")
	fmt.Println("Got it?")
	fmt.Println("Good.")
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++")
}

// FlagHandling is meant to deal with any flags given by the user through the answerFlag variable, and starts the answer generation process. It will return an answer value that holds the actual answer.
func FlagHandling(answerFlag string, answerList []string, validAnswers map[string]string) string {
	// answer initialzing
	var answer string

	//If the answerFlag had set answer passed through, then get the users input for the answer.
	if answerFlag == "set answer" {
		fmt.Println("Set the answer: ")
		//Setting the answer equal to user input
		answer = Input("regular input")

		//If the answer given by the user is not a valid answer, then prompt the user for another answer
		for !IsValidGuessOrAnswer(validAnswers, answer) {
			fmt.Println("Answer doesn't match the list of accepted answers, please enter a valid answer:")
			answer = Input("regular input")
		}
	} else {
		//If the set answer flag wasn't used, then get a random answer from the list
		answer = GetAnswer(answerList)
	}
	//If the flag for giving the answer is set, then print the answer specified above.
	if answerFlag == "give answer" {
		fmt.Println("Heres the answer dude: ")
		fmt.Println(answer)
	}
	fmt.Println("***********************************************")
	fmt.Println("The answer has been generated and the game is now ready, so type in a 5 letter word on the command line and lets begin!")
	//Returning the answer.
	return answer
}
