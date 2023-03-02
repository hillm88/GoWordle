# GoWordle

# Cloning
To clone, use ```git clone https://ghe.ops.betable.com/mason/GoWordle.git```


# Running the Program
To run it, use ```go run .```

From there, it will start the game, these are the rules:\
You get 5 guesses to get the word.\
You cannot guess a duplicate\
You have to guess an actual 5 letter word, no guesses like 'aabbc'\
You will be able to see if a letter is in the right position or not using the following rubric: 

| Color      | Meaning |
| ----------- | ----------- |
| Red      | The letter is not in the word at all|
| Yellow   | The letter is in the word but not in the right place|
| Green    | The letter is in the right place|

Once the game starts, enter a 5 letter word (it is case insensistive, so HeLlO and hello and HELLO are all equivalent) and the game will show you how your guess compared to the answer in the form of your guess but with colors behind the letters.\
Using the above color rubric, you can determine how close you are to the answer.\
This is what a regular no flag run of the game looks like:
\
\
<img width="944" alt="Screen Shot 2023-03-02 at 10 39 01 AM" src="https://user-images.githubusercontent.com/85005952/222508558-ad27870d-6171-482d-9827-aec60912aa05.png">


All of the valid guesses can be found at: https://gist.github.com/cfreshman/a7b776506c73284511034e63af1017ee/raw/wordle-nyt-answers-alphabetical.txt and https://gist.github.com/cfreshman/40608e78e83eb4e1d60b285eb7e9732f/raw/wordle-nyt-allowed-guesses.txt

All of the valid answers can be found at https://gist.github.com/cfreshman/a7b776506c73284511034e63af1017ee/raw/wordle-nyt-answers-alphabetical.txt


# Options

If you need to see the answer, run it with ```go run . -a```. This is what a run of this option would look like: 

<img width="1019" alt="Screen Shot 2023-03-02 at 10 43 19 AM" src="https://user-images.githubusercontent.com/85005952/222509341-d6c3a06e-dd89-40f0-82aa-e64d4e971ac3.png">




If you want to set the answer, use ```go run . -s``` to then be prompted for your answer. You can only enter an answer that is from the list of valid answers found at:
https://gist.github.com/cfreshman/a7b776506c73284511034e63af1017ee/raw/wordle-nyt-answers-alphabetical.txt
Here is a sample run with this option:
\
\
<img width="957" alt="Screen Shot 2023-03-02 at 10 45 07 AM" src="https://user-images.githubusercontent.com/85005952/222509666-3de51626-dd20-4b10-bd2d-8b23cd05170e.png">


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

