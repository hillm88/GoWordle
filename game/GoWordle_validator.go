package game

// IsValidGuessOrAnswer This function checks the user input against the map list given.
func IsValidGuessOrAnswer(list map[string]string, input string) bool {

	//Using the string input by the user as the lookup key for the map, if it exists
	//this will return true and if not, it will return false.
	return input == list[input]
}

// IsValidInput checks the length of the input against the maximum length the input can be
func IsValidInput(input string, exactLength int) bool {

	return len(input) == exactLength
}

// IsValidListSize checks that the list size is greater than or equal to one
func IsValidListSize(aList []string) bool {
	return len(aList) >= 1
}

// IsValidWord checks that the length of the word is five which would make it valid.
func IsValidWord(word string) bool {
	return len(word) == 5
}
