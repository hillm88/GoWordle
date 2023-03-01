# GoWordle

# Cloning
To clone, use ```git clone https://ghe.ops.betable.com/mason/GoWordle.git```


# Running the Program
To run it, use ```go run .```

From there, it will start the game, these are the rules:
You get 5 guesses to get the word.
You have to guess an actual 5 letter word, no guesses like 'aabbc'
You will be able to see if a letter is in the right position or not using the following rubric: \
1 means that the letter is not in the word at all\
2 means that the letter is in the word but not in the right place\
3 means that the letter is in the right place

Once the game starts, enter a 5 letter word (it is case insensistive, so HeLlO and hello and HELLO are all equivalent) and the game will show you how your guess compared to the answer in the form of 5 numbers like so: ```[3 2 1 3 1]```.
The position of each number corresponds to the position of each letter in the guess, so what this example would tell you is that the 1st and 4th letters are in the right place, while the 2nd letter is in the word but not in the right place. The letters in positions 3 and 5 are not in the word at all. This is what a regular no flag run of the game looks like: 
```Welcome to Wordle by the best programmer in the world, Mason 'The best of all time' Hill
Here are the rules:
You get 6 guesses to get the word.
You have to guess an actual 5 letter word, no guesses like 'aabbc'
You will be able to see if a letter is in the right position or not using the following rubric:
1 Means that the letter is not in the word at all
2 Means that the letter is in the word but not in the right place
3 Means that the letter is in the right place
Get it?
Got it?
Good.
+++++++++++++++++++++++++++++++++++++++++++++++
***********************************************
The answer has been generated and the game is now ready, so type in a 5 letter word on the command line and lets begin!
Hello
[1 2 2 2 1]
planet
Invalid guess, guess again
heavy
[1 2 1 1 1]
swell
[3 1 3 2 2]
sleep
[3 3 3 3 1]
sleet
[3 3 3 3 1]
sleek
You got it!
[3 3 3 3 3]
You won!
If you would like to play another game, enter y on the command line, if not, then press any other button and get outta here ya rascal.
```

All of the valid guesses can be found at: https://gist.github.com/cfreshman/a7b776506c73284511034e63af1017ee/raw/wordle-nyt-answers-alphabetical.txt and https://gist.github.com/cfreshman/40608e78e83eb4e1d60b285eb7e9732f/raw/wordle-nyt-allowed-guesses.txt

All of the valid answers can be found at https://gist.github.com/cfreshman/a7b776506c73284511034e63af1017ee/raw/wordle-nyt-answers-alphabetical.txt


# Options

If you need to see the answer, run it with ```go run . -a```. This is what a run of this option would look like: 
```
Get it?
Got it?
Good.
+++++++++++++++++++++++++++++++++++++++++++++++
Heres the answer dude: 
kitty
***********************************************
The answer has been generated and the game is now ready, so type in a 5 letter word on the command line and lets begin!
city
Invalid guess, guess again
hello
[1 1 1 1 1]
KiTty
You got it!
[3 3 3 3 3]
You won!
If you would like to play another game, enter y on the command line, if not, then press any other button and get outta here ya rascal.
```

If you want to set the answer, use ```go run . -s``` to then be prompted for your answer. You can only enter an answer that is from the list of valid answers found at:
https://gist.github.com/cfreshman/a7b776506c73284511034e63af1017ee/raw/wordle-nyt-answers-alphabetical.txt
Here is a sample run with this option:
```
Get it?
Got it?
Good.
+++++++++++++++++++++++++++++++++++++++++++++++
Set the answer: 
hello
***********************************************
The answer has been generated and the game is now ready, so type in a 5 letter word on the command line and lets begin!
swell
[1 1 2 3 2]
hello
You got it!
[3 3 3 3 3]
You won!
If you would like to play another game, enter y on the command line, if not, then press any other button and get outta here ya rascal.
```


When you guess an answer that has 2 of the same letter, but your guess only has one but your one is in the right place, normally the program will put a 3 there. If you would wish for their to be an indication that their may actually be another of that letter in the word when this instance arises, use ```go run . -f``` to see the number 4.
4 Means that the letter is in the right place but their are 2 or more of that letter in the word.

This is an example run with the -f flag, note that this run is actually done using ```go run . -s -f``` so that I can gaurantee one of the rare examples where a number 4 would be used.
```
Get it?
Got it?
Good.
+++++++++++++++++++++++++++++++++++++++++++++++
Set the answer: 
sewer
***********************************************
The answer has been generated and the game is now ready, so type in a 5 letter word on the command line and lets begin!
river
[1 1 1 4 3]
sewer
You got it!
[3 3 3 3 3]
You won!
If you would like to play another game, enter y on the command line, if not, then press any other button and get outta here ya rascal.
```
# How to test

To test the program, navigate to the ```Tests``` folder, then run ```go test```

Please note that some of the functions are designed to output error messages so if you see something like the ```Error with response code after downloading github page```, that appears because in the testing suite I also test for if it would reject a bad github page link, and when it does, that is the message that is displayed. If it doesn't look something like this
```
PASS
ok  	masonwordle/Tests	1.816s
``` 
at the bottom then that is concerning and shows that the tests did not actually pass.


# Things that aren't done
The testing suite for any of the input related things is not done due to manually testing inputs and not knowing how to put command line inputs in a testing script. There are also no tests for main.go since I don't actually know how to test that since it doesn't take any arguments and doesn't return any values. This includes the view testing suite as well. Also, while I tried to check every comment, some comments may be slightly outdated.

