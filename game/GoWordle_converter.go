package game

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

// GetGuessesAndAnswers will call the GetText method on both the answer and guesses github pages, and then combined the lists together to make one big guess list. It will then return the big guess list and the answer list.
func GetGuessesAndAnswers() ([]string, []string) {

	// allGuesses will hold all the possible guesses in a string array
	var allGuesses []string

	//Calling get text to get the answer list, and storing any errors in the err variable.
	answers, err := GetText("https://gist.github.com/cfreshman/a7b776506c73284511034e63af1017ee/raw/wordle-nyt-answers-alphabetical.txt")
	//if their is an error in getting the text, end the program.
	if err != nil {
		log.Fatal(err)
	}
	//Calling get text to get the guess list, and storing any errors in the err variable.
	guesses, err := GetText("https://gist.github.com/cfreshman/40608e78e83eb4e1d60b285eb7e9732f/raw/wordle-nyt-allowed-guesses.txt")
	if err != nil {
		log.Fatal(err)
	}

	//appending the guesses to the answer list to form the allguesses list.
	allGuesses = append(answers, guesses...)
	//Returning all of the guesses and the answers.
	return allGuesses, answers
}

// MapIt takes a string array and turns it into a map where each key and its value are the same string from the string array.
func MapIt(listOfWords []string) (map[string]string, error) {

	// wordToMap initializes the map
	wordToMap := make(map[string]string)
	//The for loop gets each word in the list of words and makes the key and value of the wordToMap map that word.
	for _, word := range listOfWords {
		//If the word is too long, then return an error. This is only really meant to be used for testing.
		if !IsValidWord(word) {
			return nil, fmt.Errorf(" Word was not 5 letters long")
		}
		//Putting each word into the key and value spots on the map.
		wordToMap[word] = word
	}

	//Returns a map with all of the words in it as key value pairs.
	return wordToMap, nil
}

// Input gets the users input and returns it as a string. funcInput holds a string containing a message that indicates if we are
// getting a word input or a single character. It returns the input as a string that is properly formatted according to the rules of the word or "character" input types.
func Input(funcInput string) string {
	// word is a string that will hold the guess for the user initially.
	var word string
	// err will hold any errors from reading in the user input.
	var err error

	reader := bufio.NewReader(os.Stdin)
	word, err = reader.ReadString('\n')

	if err != nil {

		log.Fatal(err)
	}
	word = strings.ToLower(word)

	//If regular input has been put into the function, make a new reader and read in the users input as a string.
	if funcInput == "regular input" {
		//Checking that the word isn't too long or too short
		if !IsValidInput(word, 6) {
			funcInput = "Failure"

		} else {

			//Getting the first 5 letters off of the string (technically 6)
			funcInput = word[0:5]

		}
		//Getting the input for if the user enters y at the end of the game.
	} else if funcInput == "y input" {
		//If the word is too long or too short, return a failure string
		if !IsValidInput(word, 2) {
			funcInput = "Failure"
		} else {
			//Take only the first character of the string.
			funcInput = string(word[0])
		}
	}

	// Return the users input.
	return funcInput
}

// GetAnswer takes in a string array of all the possible answers and then it randomly chooses one.
func GetAnswer(answerList []string) string {
	// Checking that the list isn't too short aka there has to be at least 1 word in the list.
	if IsValidListSize(answerList) {
		// length holds the length of the inputted string array
		length := len(answerList)
		//Generating a new random seed
		rand.Seed(time.Now().UnixNano())
		//Generating a random number from 0 to the length of the array - 1
		//Using Intn means that you only have to give the upper bound and the lower bound will be set at 0.
		randomNum := rand.Intn(length)

		//returning the randomNum position of the answer list, A.K.A the answer the user will now be looking for.
		return answerList[randomNum]
	} else {
		return "failure"
	}
}

// GetText works when given a github page link as input, will get the file and return each line of the file
// as a string array. It has an error output as well for the purpose of testing.
func GetText(link string) ([]string, error) {

	/*Using http.get with the inputted link to get the file located at that link and storing it in resp.
	Any errors will be put into err.
	When error is nil, resp will always contain a non-nil resp.Body
	*/
	resp, err := http.Get(link)
	//If err gets no errors then it just has nil in it.
	if err != nil {
		fmt.Println("Error in the github page download, check your internet connection please.")
		panic(err)

	}
	//This closes the response body associated with an HTTP response just before
	//this function is finished running. All http response bodies
	//must be closed after you are finished with them.
	defer resp.Body.Close()
	//http.StatusOk is defined in http as 200
	//I think that resp is actually a struct with statuscode being one of its
	//fields.
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error with response code after downloading github page")
		//Returning nil for the string and the actual error so that way I can error check.
		return nil, fmt.Errorf("error with response code after downloading github page")
	}
	//Creating a new Scanner object with the body of the response.
	scanner := bufio.NewScanner(resp.Body)
	//Creating a string array that will hold each line of the file.
	var listOfWords []string

	//Scanning in the file.
	for scanner.Scan() {
		//Getting each token from the scanner and storing it into line as a string.
		word := scanner.Text()
		//putting all the singular lines into string array lines
		listOfWords = append(listOfWords, word)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Bad scanner")
		panic(err)
	}

	//returning the list of words as a string array and any errors although if it reached this point the err should be nil.
	return listOfWords, err
}
