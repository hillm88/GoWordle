package view

import (
	"testing"

	"git.ops.betable.com/go-betable-lib/errors"

	. "github.com/onsi/ginkgo/v2"

	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"

	. "github.com/petergtz/pegomock/ginkgo_compatible"
)

func TestView(t *testing.T) {
	pegomock.RegisterMockFailHandler(Fail)
	RegisterFailHandler(Fail)
	RunSpecs(t, "view Package Test Suite")
}

var _ = Describe("View Controller", func() {
	var v View
	var lettersToColorStruct LettersToColor
	var mockValidator *MockvalidatorInterface

	BeforeEach(func() {
		mockValidator = NewMockvalidatorInterface()

	})

	JustBeforeEach(func() {

		v = View{
			validator: mockValidator,
		}

	})
	Describe("LetterColoring", func() {
		var err error

		JustBeforeEach(func() {
			err = v.LetterColoring(&lettersToColorStruct)
		})
		Context("when validation fails", func() {
			validationErr := errors.NewError("error")
			BeforeEach(func() {
				Whenever(mockValidator.validateLetterColoring(&lettersToColorStruct)).ThenReturn(validationErr)
			})
			It("returns an error", func() {
				Expect(errors.IsError(err, validationErr)).To(BeTrue())
			})

		})
		Context("when the letters are colored correctly", func() {
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
		})
		BeforeEach(func() {
			lettersToColorStruct = LettersToColor{
				GuessToColor:    "smile",
				LetterPositions: []PositionType{NotInWord, InWordWrongPos, RightPos, InWordWrongPos, NotInWord},
			}
			Whenever(mockValidator.validateLetterColoring(&lettersToColorStruct)).ThenReturn(nil)
		})

	})
})
