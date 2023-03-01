package main

import (
	"flag"
	"masonwordle/game"
)

// main is the starting point of the program. It deals with flag handling and setting up the conditions for the RunTheGame function
func main() {
	// Outputting all the intial text needed.
	game.View()
	//answerFlag is used to parse the -a flag that the program can be run with and is loaded with a bool value of whether the flag has been used or not.
	answerFlag := flag.Bool("a", false, "The answer is not being shown")
	//setAnswerFlag is used to parse the -s flag that the program can be run with.
	setAnswerFlag := flag.Bool("s", false, "The answer is not set")
	//answerFlag is used to parse the -s flag that the program can be run with.
	addFourFlag := flag.Bool("f", false, "The 4 flag is not set")
	//Parsing the flag from the cli
	flag.Parse()

	// option will hold the text value of whatever flag it is set to.
	var option string
	var fourthFlag bool
	if *setAnswerFlag {
		option = "set answer"
	} else if *answerFlag {
		option = "give answer"
	}
	if *addFourFlag {
		fourthFlag = true
	}

	// GetGuessesAndAnswers returns 2 string arrays, the guesses and the answers
	var guesses, answers = game.GetGuessesAndAnswers()

	//mapOfGuesses contains all the guesses but they are now all in a map[string]string with both strings being the same guess.
	//EX: guess is hello so map will contain map[hello]hello
	mapOfGuesses, _ := game.MapIt(guesses)

	mapOfAnswers, _ := game.MapIt(answers)

	//Starting the game.
	game.RunTheGame(option, mapOfGuesses, answers, mapOfAnswers, fourthFlag)
}
