# GoWordle

# Cloning
To clone go to the terminal and use ```git clone https://ghe.ops.betable.com/mason/GoWordle.git```

# If the program doesn't work:
The links used in the program to get the answers and guesses may get updated, deleted, or changed, so if the program is not working, please send me an email at mason.m.hill1@gmail.com and I will update the links.

# Running the Program 
To run it open a terminal and navigate into the ```GoWordle/``` folder (using ```cd GoWordle``` if your terminal allows that command) then use ```go run .```

From there, it will start the game, these are the rules:
You get 6 guesses to get the word.
You have to guess an actual 5 letter word, no guesses like 'aabbc'
You will be able to see if a letter is in the right position or not using the following rubric:

| Color | Description |
| ----------- | ----------- |
| Red | The letter is not in the word |
| Yellow | The letter is in the word but not in the right place |
|Green|The letter is in the right place|

Once the game starts, enter a 5 letter word (it is case insensitive, so HeLlO and hello and HELLO are all equivalent) and the game will show you how your guess compared to the answer in the form of your guess with the background colors of each letter corresponding to an equivalent. 

All of the valid guesses can be found at: https://gist.github.com/cfreshman/a7b776506c73284511034e63af1017ee/raw/wordle-nyt-answers-alphabetical.txt and https://gist.github.com/cfreshman/40608e78e83eb4e1d60b285eb7e9732f/raw/wordle-nyt-allowed-guesses.txt

All of the valid answers can be found at https://gist.github.com/cfreshman/a7b776506c73284511034e63af1017ee/raw/wordle-nyt-answers-alphabetical.txt


# Options

## View Answer

If you need to see the answer, run it with ```go run . -a```.


## Set Answer

If you want to set the answer, use ```go run . -s``` to then be prompted for your answer. You can only enter an answer that is from the list of valid answers found at:
https://gist.github.com/cfreshman/a7b776506c73284511034e63af1017ee/raw/wordle-nyt-answers-alphabetical.txt

# How to test

To test the program, navigate to the folder you would like to test, in this case it would be one of the ```game```, ```view```,```rga```,```answerflag```, or ```inpt``` folders, then run ```go test```

Please note that some of the functions are designed to output error messages so if you see something like the ```Error with response code after downloading github page```, that appears because in the testing suite I also test for if it would reject a bad github page link, and when it does, that is the message that is displayed. If it doesn't look something like this
```
PASS
ok  	masonwordle/Tests	1.816s
``` 
at the bottom then that is concerning and shows that the tests did not actually pass.





