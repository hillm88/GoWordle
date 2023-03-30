package game

import (
	"testing"

	"git.ops.betable.com/go-betable-lib/errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"
	. "github.com/petergtz/pegomock/ginkgo_compatible"
)

func TestGame(t *testing.T) {
	pegomock.RegisterMockFailHandler(Fail)
	RegisterFailHandler(Fail)
	RunSpecs(t, "Game Package Test Suite")
}

var _ = Describe("Game Controller", func() {
	var c Game
	var gameStruct GameState
	var mockInpt *MockinptInterface
	var mockView *MockviewInterface
	var mockConverter *MockconverterInterface
	var mockValidator *MockvalidatorInterface
	var mockPrivate *MockprivateMethodInterface

	BeforeEach(func() {
		mockInpt = NewMockinptInterface()
		mockView = NewMockviewInterface()
		mockConverter = NewMockconverterInterface()
		mockValidator = NewMockvalidatorInterface()
		mockPrivate = NewMockprivateMethodInterface()

	})

	JustBeforeEach(func() {
		c = Game{
			inpt:      mockInpt,
			view:      mockView,
			converter: mockConverter,
			validator: mockValidator,
			private:   mockPrivate,
		}

	})
	Describe("RunTheGame", func() {
		var err error
		var won bool
		JustBeforeEach(func() {
			err = c.RunTheGame(&gameStruct)
		})
		Context("when validation fails", func() {
			validationErr := errors.NewError("error")

			BeforeEach(func() {
				Whenever(mockValidator.validateGame(&gameStruct)).ThenReturn(validationErr)
			})
			It("returns the validation error", func() {
				Expect(errors.IsError(err, validationErr)).To(BeTrue())
			})
		})
		Context("when playTheGameForSixRounds returns an error", func() {
			var logicErr error
			BeforeEach(func() {
				logicErr = errors.NewError("playTheGameForSixRounds error")

				Whenever(mockPrivate.playTheGameForSixRounds(&gameStruct)).ThenReturn(false, logicErr)
			})
			It("returns the logicErr", func() {
				Expect(errors.IsError(err, logicErr)).To(BeTrue())
			})
		})
		Context("when the game is able to run successfully", func() {
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
		})
		BeforeEach(func() {
			won = true
			gameStruct = GameState{
				Answer:       "smell",
				ValidGuesses: map[string]string{"hello": "hello", "smell": "smell"},
				AnswerList:   []string{"hello", "smell"},
				ValidAnswers: map[string]string{"hello": "hello", "smell": "smell"},
			}
			Whenever(mockValidator.validateGame(&gameStruct)).ThenReturn(nil)
			Whenever(mockPrivate.playTheGameForSixRounds(&gameStruct)).ThenReturn(won, nil)

		})
	})
	Describe("letterPositionWinCheckerAndSetter", func() {
		var err error
		var win bool
		var arrayOfPositions []PositionType
		var guess string

		JustBeforeEach(func() {
			win, arrayOfPositions, err = c.letterPositionWinCheckerAndSetter(&gameStruct, guess)
		})

		Context("when guess is not the answer", func() {

			BeforeEach(func() {
				guess = "their"

			})
			It("returns false", func() {
				Expect(win).To(BeFalse())

			})
			It("returns an array of positions with how close each is to the actual answer", func() {
				Expect(arrayOfPositions).To(ContainElements(RightPos, RightPos, RightPos, NotInWord, InWordWrongPos))
			})
		})
		Context("guess is the answer", func() {

			It("returns a nil error", func() {

				Expect(err).To(BeNil())

			})
			It("returns a true bool", func() {
				Expect(win).To(BeTrue())

			})
			It("returns an array of position type with each position equal to RightPos", func() {
				Expect(arrayOfPositions).To(ContainElements(RightPos, RightPos, RightPos, RightPos, RightPos))
			})
		})
		BeforeEach(func() {
			gameStruct = GameState{
				Answer:       "there",
				ValidGuesses: map[string]string{"hello": "hello", "smell": "smell", "there": "there"},
				AnswerList:   []string{"hello", "smell", "there"},
				ValidAnswers: map[string]string{"hello": "hello", "smell": "smell", "there": "there"},
			}
			guess = "there"
		})
	})

	Describe("isValidGuess", func() {
		mapOfPreviousGuesses := make(map[string]string)
		var err error
		var guess string
		JustBeforeEach(func() {
			err = c.isValidGuess(&gameStruct, guess, mapOfPreviousGuesses)
		})
		Context("when guess has already been guessed", func() {
			BeforeEach(func() {
				mapOfPreviousGuesses["hello"] = "hello"
				guess = "hello"
			})
			It("returns answerAlreadyGuessed error", func() {
				Expect(err).To(Equal(answerAlreadyGuessed))
			})

		})
		Context("when guess is not a valid guess", func() {
			BeforeEach(func() {
				guess = "invalidGuessHere"
			})
			It("returns invalidGuess error", func() {
				Expect(err).To(Equal(invalidGuess))
			})

		})

		Context("when guess is a valid guess", func() {

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

		})

		BeforeEach(func() {

			gameStruct = GameState{
				Answer:       "smell",
				ValidGuesses: map[string]string{"hello": "hello", "smile": "smile", "smell": "smell"},
				AnswerList:   []string{"hello", "smell"},
				ValidAnswers: map[string]string{"hello": "hello", "smell": "smell"},
			}
			guess = "smile"

		})
	})

	Describe("playTheGameForSixRounds", func() {
		var err error
		var win bool

		JustBeforeEach(func() {
			win, err = c.playTheGameForSixRounds(&gameStruct)
		})
		Context("when there is invalid input", func() {
			inputErr := errors.NewError("error")
			BeforeEach(func() {
				Whenever(mockInpt.GetValidWordInput()).ThenReturn("", inputErr)
			})
			It("returns input error", func() {

				Expect(errors.IsError(err, inputErr)).To(BeTrue())

			})
		})
		Context("when the letterPositionWinCheckerAndSetter throws an error", func() {
			letterPositionWinCheckerAndSetterErr := errors.NewError("error")
			BeforeEach(func() {

				Whenever(mockPrivate.letterPositionWinCheckerAndSetter(&gameStruct, "hello")).ThenReturn(false, nil, letterPositionWinCheckerAndSetterErr)
			})
			It("returns an error", func() {

				Expect(errors.IsError(err, letterPositionWinCheckerAndSetterErr)).To(BeTrue())

			})

		})
		Context("when the LetterColoring throws an error", func() {
			letterColoringErr := errors.NewError("error")
			BeforeEach(func() {

				Whenever(mockPrivate.letterPositionWinCheckerAndSetter(&gameStruct, "hello")).ThenReturn(false, nil, nil)
				Whenever(mockView.LetterColoring(nil)).ThenReturn(letterColoringErr)
			})
			It("returns an error", func() {

				Expect(errors.IsError(err, letterColoringErr)).To(BeTrue())

			})

		})
		Context("when the maximum number of rounds have been played", func() {

			BeforeEach(func() {

				Whenever(mockPrivate.letterPositionWinCheckerAndSetter(&gameStruct, "hello")).ThenReturn(false, nil, nil)

			})
			It("returns no error", func() {

				Expect(err).To(BeNil())

			})
			It("returns false", func() {
				Expect(win).To(BeFalse())
			})
		})
		Context("when the winnerFlag is true", func() {
			It("returns no error", func() {

				Expect(err).To(BeNil())

			})
			It("returns true", func() {
				Expect(win).To(BeTrue())
			})
		})
		BeforeEach(func() {
			gameStruct = GameState{
				Answer:       "smell",
				ValidGuesses: map[string]string{"hello": "hello", "smell": "smell"},
				AnswerList:   []string{"hello", "smell"},
				ValidAnswers: map[string]string{"hello": "hello", "smell": "smell"},
			}
			Whenever(mockInpt.GetValidWordInput()).ThenReturn("hello", nil)
			Whenever(mockPrivate.isValidGuess(&gameStruct, "hello", nil)).ThenReturn(nil)
			Whenever(mockPrivate.letterPositionWinCheckerAndSetter(&gameStruct, "hello")).ThenReturn(true, nil, nil)
			Whenever(mockView.LetterColoring(nil)).ThenReturn(nil)
		})

	})

	Describe("addToMap", func() {
		var maps map[string]string
		currentMap := make(map[string]string)
		JustBeforeEach(func() {
			maps = c.addToMap(currentMap, "hello")
		})
		Context("when added to map", func() {
			It("returns map with string added", func() {
				Expect(maps["hello"]).To(Equal("hello"))
			})
		})

		BeforeEach(func() {
			currentMap["smile"] = "smile"
			currentMap["lunch"] = "lunch"
		})
	})
})
