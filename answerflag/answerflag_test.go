package answerflag

import (
	"testing"

	"git.ops.betable.com/go-betable-lib/errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"
	. "github.com/petergtz/pegomock/ginkgo_compatible"
)

func TestAnswerFlagHandler(t *testing.T) {

	pegomock.RegisterMockFailHandler(Fail)
	RegisterFailHandler(Fail)
	RunSpecs(t, "answerflag Package Test Suite")
}

var _ = Describe("AnswerFlag Controller", func() {
	var f Flag

	var mockInpt *MockinptInterface

	var mockValidator *MockvalidatorInterface
	var mockPrivate *MockprivateMethodInterface

	BeforeEach(func() {
		mockInpt = NewMockinptInterface()

		mockValidator = NewMockvalidatorInterface()
		mockPrivate = NewMockprivateMethodInterface()

	})

	JustBeforeEach(func() {
		f = Flag{
			validator: mockValidator,
			private:   mockPrivate,
			inpt:      mockInpt,
		}

	})

	Describe("AnswerFlagHandler", func() {
		var err error
		var answer string
		var answerInfoHolderStruct AnswerInfoHolder
		var noFlagsAnswerInfoHolderStruct, setAnswerFlagEnabledAnswerInfoHolderStruct, showAnswerFlagEnabledAnswerInfoHolderStruct AnswerInfoHolder
		JustBeforeEach(func() {
			answer, err = f.AnswerFlagHandler(&answerInfoHolderStruct)
		})
		Context("when validation fails", func() {

			validationErr := errors.NewError("error")
			BeforeEach(func() {
				answerInfoHolderStruct = noFlagsAnswerInfoHolderStruct
				Whenever(mockValidator.validateAnswerLists(&answerInfoHolderStruct)).ThenReturn(validationErr)
			})
			It("returns the validation error", func() {
				Expect(errors.IsError(err, validationErr)).To(BeTrue())
			})
		})

		Context("when setAnswer returns an error", func() {

			setAnswerErr := errors.NewError("error")
			BeforeEach(func() {
				answerInfoHolderStruct = setAnswerFlagEnabledAnswerInfoHolderStruct
				Whenever(mockPrivate.setAnswer(answerInfoHolderStruct.ValidAnswers)).ThenReturn("", setAnswerErr)
			})
			It("returns an error", func() {
				Expect(errors.IsError(err, setAnswerErr)).To(BeTrue())
			})
		})
		Context("when setAnswer runs without error", func() {

			BeforeEach(func() {
				answerInfoHolderStruct = setAnswerFlagEnabledAnswerInfoHolderStruct
				Whenever(mockPrivate.setAnswer(answerInfoHolderStruct.ValidAnswers)).ThenReturn("hello", nil)
			})
			It("returns the answer string", func() {
				Expect(answer).To(Equal("hello"))
			})
		})

		Context("when getAnswer throws an error", func() {

			getAnswerErr := errors.NewError("error")
			BeforeEach(func() {
				answerInfoHolderStruct = showAnswerFlagEnabledAnswerInfoHolderStruct

				Whenever(mockPrivate.getAnswer(answerInfoHolderStruct.AnswerList)).ThenReturn("", getAnswerErr)
			})
			It("returns an error", func() {
				Expect(errors.IsError(err, getAnswerErr)).To(BeTrue())
			})
		})
		Context("when the showAnswerFlag is set to true", func() {

			BeforeEach(func() {
				answerInfoHolderStruct = showAnswerFlagEnabledAnswerInfoHolderStruct

				Whenever(mockPrivate.getAnswer(answerInfoHolderStruct.AnswerList)).ThenReturn("hello", nil)
			})
			It("returns an answer in the form of a string", func() {
				Expect(answer).To(BeEquivalentTo("hello"))
			})
		})
		Context("when nothing is wrong", func() {

			It("returns the random answer", func() {
				Expect(answer).To(BeAssignableToTypeOf(""))
			})
			It("returns a nil error", func() {
				Expect(err).To(BeNil())
			})
		})
		BeforeEach(func() {
			noFlagsAnswerInfoHolderStruct = AnswerInfoHolder{
				AnswerList:     []string{"hello", "smell"},
				ValidAnswers:   map[string]string{"hello": "hello", "smell": "smell"},
				SetAnswerFlag:  false,
				ShowAnswerFlag: false,
			}
			setAnswerFlagEnabledAnswerInfoHolderStruct = AnswerInfoHolder{
				AnswerList:     []string{"hello", "smell"},
				ValidAnswers:   map[string]string{"hello": "hello", "smell": "smell"},
				SetAnswerFlag:  true,
				ShowAnswerFlag: false,
			}
			showAnswerFlagEnabledAnswerInfoHolderStruct = AnswerInfoHolder{
				AnswerList:     []string{"hello", "smell"},
				ValidAnswers:   map[string]string{"hello": "hello", "smell": "smell"},
				SetAnswerFlag:  false,
				ShowAnswerFlag: true,
			}
			answerInfoHolderStruct = noFlagsAnswerInfoHolderStruct
			Whenever(mockValidator.validateAnswerLists(&answerInfoHolderStruct)).ThenReturn(nil)

		})

	})

	Describe("getAnswer", func() {
		var err error
		var answer string
		var singleAnswerList [1]string

		JustBeforeEach(func() {

			answer, err = f.getAnswer(singleAnswerList[0:1])
		})
		Context("when answerList size is 0", func() {

			BeforeEach(func() {

				Whenever(mockPrivate.isValidListSize(singleAnswerList[0:1])).ThenReturn(false)
			})

			It("returns an error", func() {
				Expect(err).To(Equal(answerListEmpty))
			})
		})
		Context("when answerList is a valid size", func() {

			It("returns a valid answer", func() {
				Expect(answer).To(Equal("hello"))
			})

		})
		BeforeEach(func() {
			singleAnswerList[0] = "hello"

			Whenever(mockPrivate.isValidListSize(singleAnswerList[0:1])).ThenReturn(true)
		})
	})

	Describe("setAnswer", func() {
		var err error
		var answer string

		JustBeforeEach(func() {

			answer, err = f.setAnswer(map[string]string{"hello": "hello"})
		})
		Context("when user inputs an invalid word", func() {
			inptError := errors.NewError("error")

			BeforeEach(func() {

				Whenever(mockInpt.GetValidWordInput()).ThenReturn("", inptError)
			})
			It("returns the inptError", func() {
				Expect(errors.IsError(err, inptError)).To(BeTrue())
			})
		})

		Context("when the answer doesn't equal a valid answer and the retry input throws an error", func() {
			inptError := errors.NewError("error")
			BeforeEach(func() {
				Whenever(mockInpt.GetValidWordInput()).ThenReturn("asasd", nil).ThenReturn("", inptError)

			})
			It("returns an error", func() {
				Expect(errors.IsError(err, inptError)).To(BeTrue())
			})

		})
		Context("when the answer doesn't equal a valid answer and the retry input is valid", func() {

			BeforeEach(func() {

				Whenever(mockInpt.GetValidWordInput()).ThenReturn("asasd", nil).ThenReturn("hello", nil)

			})
			It("returns the answer", func() {
				Expect(answer).To(Equal("hello"))
			})

		})
		Context("when the answer is a valid answer", func() {

			It("returns the answer", func() {
				Expect(answer).To(Equal("hello"))
			})

		})
		BeforeEach(func() {

			Whenever(mockInpt.GetValidWordInput()).ThenReturn("hello", nil)

		})
	})
	Describe("isValidListSize", func() {
		var isValidSize bool
		var answerList [1]string
		var answerListSlice []string
		JustBeforeEach(func() {
			isValidSize = f.isValidListSize(answerListSlice)
		})
		Context("when list size is 0", func() {
			BeforeEach(func() {
				answerList[0] = ""
				answerListSlice = answerList[0:0]

			})
			It("returns false", func() {
				Expect(isValidSize).To(BeFalse())
			})
		})
		Context("when list size is 1", func() {

			It("returns true", func() {
				Expect(isValidSize).To(BeTrue())
			})
		})
		BeforeEach(func() {
			answerList[0] = "hello"
			answerListSlice = answerList[0:1]

		})

	})

})
