package inpt

import (
	"testing"

	"git.ops.betable.com/go-betable-lib/errors"

	. "github.com/onsi/ginkgo/v2"

	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"

	. "github.com/petergtz/pegomock/ginkgo_compatible"
)

func TestInpt(t *testing.T) {
	pegomock.RegisterMockFailHandler(Fail)
	RegisterFailHandler(Fail)
	RunSpecs(t, "Game Package Test Suite")
}

var _ = Describe("Inpt Controller", func() {
	var i Inpt
	var mockPrivate *MockprivateMethodInterface

	BeforeEach(func() {
		mockPrivate = NewMockprivateMethodInterface()

	})

	JustBeforeEach(func() {
		i = Inpt{
			private: mockPrivate,
		}
	})

	Describe("GameContinueDecision", func() {
		var err error
		var continueGame bool
		JustBeforeEach(func() {
			continueGame, err = i.GameContinueDecision()
		})
		Context("when getInput returns an error", func() {
			inputErr := errors.NewError("error")
			BeforeEach(func() {
				Whenever(mockPrivate.getInput()).ThenReturn("", inputErr)
			})
			It("returns an error", func() {
				Expect(errors.IsError(err, inputErr)).To(BeTrue())

			})
		})

		Context("when the input is to end the game", func() {

			BeforeEach(func() {
				Whenever(mockPrivate.getInput()).ThenReturn("n", nil)
			})
			It("returns false", func() {

				Expect(continueGame).To(BeFalse())
			})
		})

		Context("when the input is to continue the game", func() {
			It("returns true", func() {

				Expect(continueGame).To(BeTrue())
			})
		})
		BeforeEach(func() {
			Whenever(mockPrivate.getInput()).ThenReturn("y", nil)

		})
	})

	Describe("GetValidWordInput", func() {
		var err error
		var word string
		JustBeforeEach(func() {
			word, err = i.GetValidWordInput()
		})
		Context("when getInput returns an error", func() {
			inputErr := errors.NewError("error")
			BeforeEach(func() {
				Whenever(mockPrivate.getInput()).ThenReturn("", inputErr)
			})
			It("returns an error", func() {
				Expect(errors.IsError(err, inputErr)).To(BeTrue())
			})
		})

		Context("when the first word is invalid but the second one is not", func() {

			BeforeEach(func() {

				Whenever(mockPrivate.isValidInput("hello", 5)).ThenReturn(false).ThenReturn(true)

			})
			It("returns the word", func() {
				Expect(word).To(Equal("hello"))
			})
		})
		Context("when the first word is valid", func() {
			It("returns the word", func() {
				Expect(word).To(Equal("hello"))
			})
		})
		BeforeEach(func() {
			Whenever(mockPrivate.getInput()).ThenReturn("hello", nil)
			Whenever(mockPrivate.isValidInput("hello", 5)).ThenReturn(true)

		})

	})

	Describe("getInput", func() {
		var input string
		var err error
		JustBeforeEach(func() {
			input, err = i.getInput()
		})
		Context("when retrieving the cli input throws an error", func() {
			cliError := errors.NewError("error")
			BeforeEach(func() {
				Whenever(mockPrivate.retrieveCLIInput()).ThenReturn("", cliError)
			})
			It("returns an error", func() {
				Expect(errors.IsError(err, cliError)).To(BeTrue())
			})
		})
		Context("when there are no errors", func() {

			It("returns an string", func() {
				Expect(input).To(BeEquivalentTo("hello"))
			})
		})

		BeforeEach(func() {
			Whenever(mockPrivate.retrieveCLIInput()).ThenReturn("hello", nil)
		})

	})
	Describe("isValidInput", func() {
		var isValid bool
		var expectedLength int
		var wrongSize int
		JustBeforeEach(func() {
			isValid = i.isValidInput("smell", expectedLength)
		})
		Context("when size of word isn't the same as expected length", func() {
			BeforeEach(func() {
				expectedLength = wrongSize
			})
			It("returns false", func() {
				Expect(isValid).To(BeFalse())
			})
		})
		Context("when size of word is same as expected length", func() {

			It("returns true", func() {
				Expect(isValid).To(BeTrue())
			})
		})
		BeforeEach(func() {
			expectedLength = 5
			wrongSize = 4
		})
	})

})
