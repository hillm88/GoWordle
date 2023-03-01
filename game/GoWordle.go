// Package game is used to run the wordle game.
package game

import (
	"fmt"
)

// RunTheGame is a method that will handle the main execution of the game, doing most of the function calling and it will call itself if the user inputs 'y' to start another game after
// the current game has finished. The fourthFlag variable is meant to see if the user is to be able to see a latter that is in the right place but their are 2 or more of the letter.
func RunTheGame(answerFlag string, validGuesses map[string]string, answerList []string, validAnswers map[string]string, fourthFlag bool) {
	// This for loop allows the game to be played round after round until the user eventually presses a key other than y at the end to end the game.
	for {

		// answer is set to the answer selected in the start round function. This will be the answer that the users are guessing at for the rest of the game.
		answer := FlagHandling(answerFlag, answerList, validAnswers)

		// winnerFlag holds the value of whether a player has won or not in a bool value.
		winnerFlag := false

		//This is the game loop, where the 6 rounds take place
		for i := 0; i < len(validGuesses) && i < 6; {

			//guess is set to the users guess which is recieved in the input function. The empty "" was put so that when testing, a user could put a string in to see if the input function worked
			//NOTE: ASK DREW ABOUT THIS ******************************************
			guess := Input("regular input")

			//IsValidGuessOrAnswer will see if the users guess exists in the list of possible valid guesses. If it does, the rest of the game logic can be carried out on the guess.
			if IsValidGuessOrAnswer(validGuesses, guess) {
				//winnerFlag is set to the boolean result of the LetterPositionSetter function where the guess is checked to see how right it is.
				winnerFlag, _ = LetterPositionSetter(answer, guess, fourthFlag)
				//if winner flag == true, then the user has won and the program will break out of the current games loop and prompt the user for if they would like to keep playing or not.
				if winnerFlag {
					fmt.Println("You won!")
					break
				}
				//incrementing game progress by counting this guess as a valid guess
				i++
			} else {
				//If an invalid guess is made, then the invalid guess message displays
				fmt.Println("Invalid guess, guess again")
			}

		}
		//If by the end of the game, the user hasn't guesses the right word, the answer will be displayed
		if !winnerFlag {
			fmt.Println("Welp, you didn't get it, the answer was: " + answer)
		}

		//Getting the user input to see if they would like to play another game.
		fmt.Println("If you would like to play another game, enter y on the command line, if not, then press any other button and get outta here ya rascal.")
		//input holds a singular letter of the string type, representing the users input
		input := Input("y input")
		//if the user input is anything besides y, then they don't want to play another game and the program will break out of the loop and conclude.
		//if the user input is y, then they do want to play another game and the for loop will run.
		if input != "y" {
			break
		}
	}
}

// LetterPositionSetter takes in an answer, and a guess, and it checks if the user guessed right and got the answer, and if they did it will return true.
// If they didn't, then it will show them the status of each letter in their guess through the int array and it will return false. It also takes in the flag to see if the number four is to be made possible on the letterPosition array.
func LetterPositionSetter(answer string, guess string, fourthFlag bool) (bool, []int) {
	// winnerFlag holds whether the user has won or not in the form of true or false.
	winnerFlag := false

	// letterPosition is an array of intergers that represent the guess and each letter and its position, and each letters relation to the answer.
	letterPosition := []int{1, 1, 1, 1, 1}

	//Going through each letter of the guess and checking it against the answer.
	for j := 0; j < len(answer); j++ {

		//If the guess and the answer have the same letter at the same position, the
		//corresponding letterPosition value is set to 3.
		if guess[j] == answer[j] {
			letterPosition[j] = 3
		}
	}

	//These for loops allow the program to check each position of the letter and the answer against each other to see if a letter exists but not in the position is appears in the answer.
	for m := 0; m < len(answer); m++ {
		for k := 0; k < len(answer); k++ {

			/* This is meant to say that when we are guess checking each letter against the letters in the answer,
			if the guess at the location of the letter in the answer is set to 3, then we check that the answer and guess aren't already matched up by checking
			k against m, then we check if the guess letter we are looking at is set to 3 and if it isn't, we know that
			the guess at the answer position is correct but their is another one of that letter in the word, hence we set it to 4.

			If their is a similar letter in both the answer and the guess, and the position of the letter in the guess is not set to 3 or 4,
			and the letter hasn't already been correctly guessed at the position it appears in the answer, then set the letter position in the guess to 2.
			*/

			if guess[m] == answer[k] && letterPosition[m] == 3 && k != m && letterPosition[k] != 3 && fourthFlag {

				letterPosition[m] = 4

			} else if guess[m] == answer[k] && letterPosition[m] != 3 && letterPosition[m] != 4 && letterPosition[k] != 3 {
				letterPosition[m] = 2
			}
		}
	}

	if guess == answer {
		fmt.Println("You got it!")
		winnerFlag = true
	}

	//Printing the positions of all the letters.
	fmt.Println(letterPosition)
	//returning whether the user has won or lost.
	return winnerFlag, letterPosition
}
