// Package rga is used to Retrieve the list of possible Guesses and Answers.
package rga

import (
	"bufio"
	"fmt"
	"os"

	"net/http"
)

type RGA struct {
	validator validatorInterface
	private   privateMethodInterface
}

func New() *RGA {
	rga := &RGA{
		validator: validator{},
	}
	rga.private = rga

	return rga
}

// mapIt takes a string array and turns it into a map where each key and its value are the same string from the string array.
func (r *RGA) mapIt(listOfWords []string) map[string]string {

	// wordToMap initializes the map
	wordToMap := make(map[string]string)
	//The for loop gets each word in the list of words and makes the key and value of the wordToMap map that word.
	for _, word := range listOfWords {
		//Putting each word into the key and value spots on the map.
		wordToMap[word] = word
	}

	//Returns a map with all of the words in it as key value pairs.
	return wordToMap
}

// GetGuessesAndAnswers will call the GetText method on both the answer and guesses github pages, and then combined the lists together to make one big guess list. It will then return the big guess list and the answer list.
func (r *RGA) GetGuessesAndAnswers() ([]string, map[string]string, map[string]string, error) {

	// allGuesses will hold all the possible guesses in a string array
	var allGuesses []string

	//Calling getText to get the answer list, and storing any errors in the err variable.
	answers, err := r.private.fileRetriever(answerlink, answersBackupPath)
	if err != nil {
		return nil, nil, nil, err
	}
	//Calling getText to get the guess list, and storing any errors in the err variable.
	guesses, err := r.private.fileRetriever(guesslink, guessesBackupPath)
	if err != nil {
		return nil, nil, nil, err

	}

	//appending the guesses to the answer list to form the allGuesses list.
	allGuesses = append(answers, guesses...)

	mapOfGuesses := r.private.mapIt(allGuesses)

	mapOfAnswers := r.private.mapIt(answers)

	//Returning all of the guesses and the answers.
	return answers, mapOfGuesses, mapOfAnswers, err
}

func (r *RGA) fileRetriever(fileLink string, backUpLocalPath string) ([]string, error) {
	file, err := r.private.getText(fileLink)
	if err != nil {
		fmt.Println("Trying to read from backup")
		file, err = r.private.backUpFiles(backUpLocalPath)
		if err != nil {
			fmt.Println("Can't read from backup")
			return nil, err

		}

	}
	return file, err
}

// getText takes a link in the form of a string and returns its contents as a string array. It has an error output as well for the purpose of testing.
func (r *RGA) getText(link string) ([]string, error) {

	/*Using http.get with the inputted link to get the file located at that link and storing it in resp.
	Any errors will be put into err.
	When error is nil, resp will always contain a non-nil resp.Body
	*/
	resp, err := http.Get(link)
	//If err gets no errors then it just has nil in it.
	if err != nil {
		fmt.Println("Error in the github page download, please check your internet connection.")
		return nil, err

	}
	//This closes the response body associated with an HTTP response just before
	//this function is finished running. All http response bodies
	//must be closed after you are finished with them.
	defer resp.Body.Close()
	//http.StatusOk is defined in http as 200
	//I think that resp is actually a struct with Statuscode being one of its
	//fields.
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error with response code after downloading github page")
		//Returning nil for the string and the actual error so that way I can error check.
		return nil, responseCodeErr
	}
	//Creating a new Scanner object with the body of the text file held in resp.
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

	err = scanner.Err()

	//returning the list of words as a string array and any errors although if it reached this point the err should be nil.
	return listOfWords, err
}

// backUpFiles is only to be called if the list of words from github cannot be retrieved. When the repo is cloned, a copy of the answer and guess lists is included and this
// will read from it and do all the functions that getText should do but on the backup files.
func (r *RGA) backUpFiles(backUpFile string) ([]string, error) {

	file, err := os.Open(backUpFile)

	if err != nil {

		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

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
		return nil, err
	}

	return listOfWords, err
}
