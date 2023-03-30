package main

import (
	"flag"
	"log"
	"masonwordle/answerflag"
	"masonwordle/game"
	"masonwordle/inpt"
	"masonwordle/rga"
	"masonwordle/view"
)

// main is the starting point of the program. It deals with flag handling and setting up the conditions for the RunTheGame function
func main() {

	//All 3 flags are handled here.
	showAnswerFlag := flag.Bool("a", false, "The answer is not being shown")
	setAnswerFlag := flag.Bool("s", false, "The answer is not set")
	//Parsing the command line flags from os.Args[1:]. Must be called after all flags are defined but before flags are accessed by program.
	flag.Parse()

	inputController := inpt.New()
	viewController := view.New()
	gameController := game.New(inputController, viewController)
	flagController := answerflag.New(inputController)
	guessAndAnswersController := rga.New()

	// GetGuessesAndAnswers returns 2 string arrays, the guesses and the answers plus any errors.
	answers, mapOfGuesses, mapOfAnswers, errA := guessAndAnswersController.GetGuessesAndAnswers()
	if errA != nil {
		log.Fatalf(errA.Error())
	}

	converter := converter{}

	var answer string

	for {

		//Creating a new instance of the game struct.
		gameStruct := &game.GameState{
			Answer:       answer,
			ValidGuesses: mapOfGuesses,
			AnswerList:   answers,
			ValidAnswers: mapOfAnswers,
		}
		//Getting an answer
		answer, err := flagController.AnswerFlagHandler(converter.convertToAnswerFlag(gameStruct, *setAnswerFlag, *showAnswerFlag))
		if err != nil {
			panic(err)
		}

		gameStruct.Answer = answer

		//Calling the main run the game method and storing the error it might have
		errA := gameController.RunTheGame(gameStruct)

		if errA != nil {
			panic(errA)
		}

		if tf, err := inputController.GameContinueDecision(); !tf || err != nil {
			break
		}
	}
}
