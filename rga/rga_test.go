package rga

import (
	"testing"

	"git.ops.betable.com/go-betable-lib/errors"
	. "github.com/onsi/ginkgo/v2"

	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"
	. "github.com/petergtz/pegomock/ginkgo_compatible"
)

func TestRGA(t *testing.T) {
	pegomock.RegisterMockFailHandler(Fail)
	RegisterFailHandler(Fail)
	RunSpecs(t, "Game Package Test Suite")
}

var _ = Describe("rga Controller", func() {
	var r RGA
	var mockPrivate *MockprivateMethodInterface
	var mockValidator *MockvalidatorInterface

	BeforeEach(func() {
		mockPrivate = NewMockprivateMethodInterface()
		mockValidator = NewMockvalidatorInterface()
	})

	JustBeforeEach(func() {
		r = RGA{
			private:   mockPrivate,
			validator: mockValidator,
		}
	})

	Describe("mapIt", func() {
		var answerList [2]string

		var answerListAsAMap map[string]string

		JustBeforeEach(func() {

			answerListAsAMap = r.mapIt(answerList[0:1])
		})
		Context("when the map is correct", func() {

			It("will return a map with the answerList words in it", func() {
				Expect(answerListAsAMap["hello"] == "hello").To(BeTrue())

			})
		})
		BeforeEach(func() {
			answerList[0] = "hello"
			answerList[1] = "world"
		})

	})
	Describe("GetGuessesAndAnswers", func() {
		var err error
		var listOfAnswers []string
		var mapOfGuesses map[string]string
		var guessesStringSlice []string
		var mapOfAnswers map[string]string
		JustBeforeEach(func() {
			listOfAnswers, mapOfGuesses, mapOfAnswers, err = r.GetGuessesAndAnswers()
		})

		//APPROACH THIS LATER, MAYBE REFACTOR.
		Context("when the link for the answers file is unable to be retrieved", func() {
			retrievingErr := errors.NewError("error")
			BeforeEach(func() {
				Whenever(mockPrivate.fileRetriever(answerlink, answersBackupPath)).ThenReturn(nil, retrievingErr)

			})
			It("will return an error", func() {
				Expect(errors.IsError(err, retrievingErr)).To(BeTrue())
			})
		})

		Context("when the link for the answers file is correct and the guessesStringSlice file is unable to be retrieved", func() {
			retrievingErr := errors.NewError("error")
			BeforeEach(func() {

				Whenever(mockPrivate.fileRetriever(guesslink, guessesBackupPath)).ThenReturn(nil, retrievingErr)

			})
			It("will return an error", func() {
				Expect(errors.IsError(err, retrievingErr)).To(BeTrue())
			})
		})

		Context("when the text is able to be retrieved successfully from both answers and guessesStringSlice", func() {

			It("will return an a string list", func() {
				Expect(listOfAnswers).To(BeEquivalentTo([]string{"hello", "smell"}))
			})
			It("will return a map of the answers", func() {
				Expect(mapOfAnswers).To(BeEquivalentTo(map[string]string{"hello": "hello", "smell": "smell"}))
			})
			It("will return a map of the guessesStringSlice", func() {
				Expect(mapOfGuesses).To(BeEquivalentTo(map[string]string{"hello": "hello", "smell": "smell"}))
			})
		})
		BeforeEach(func() {
			guessesStringSlice = append(listOfAnswers, listOfAnswers...)
			Whenever(mockPrivate.fileRetriever(answerlink, answersBackupPath)).ThenReturn([]string{"hello", "smell"}, nil)
			Whenever(mockPrivate.fileRetriever(guesslink, guessesBackupPath)).ThenReturn([]string{"hello", "smell"}, nil)
			Whenever(mockPrivate.mapIt(guessesStringSlice)).ThenReturn(map[string]string{"hello": "hello", "smell": "smell"})
			Whenever(mockPrivate.mapIt(listOfAnswers)).ThenReturn(map[string]string{"hello": "hello", "smell": "smell"})
		})

	})

	Describe("fileRetriever", func() {
		var fileWordsArray []string
		var err error
		JustBeforeEach(func() {
			fileWordsArray, err = r.fileRetriever("", "")
		})
		Context("when it cannot read the file link but can read the backup", func() {
			fileLinkRetrievalErr := errors.NewError("error")
			BeforeEach(func() {
				Whenever(mockPrivate.getText("")).ThenReturn(nil, fileLinkRetrievalErr)
				Whenever(mockPrivate.backUpFiles("")).ThenReturn([]string{"hello", "smile"}, nil)
			})
			It("will read from the backup path and return an slice of 2 strings", func() {
				Expect(fileWordsArray).To(BeEquivalentTo([]string{"hello", "smile"}))
			})

		})
		Context("when it cannot read the file link and the backup is also unreadable", func() {
			fileLinkRetrievalErr := errors.NewError("error")
			BeforeEach(func() {
				Whenever(mockPrivate.getText("")).ThenReturn(nil, fileLinkRetrievalErr)
				Whenever(mockPrivate.backUpFiles("")).ThenReturn(nil, fileLinkRetrievalErr)
			})
			It("will return an error", func() {
				Expect(errors.IsError(err, fileLinkRetrievalErr)).To(BeTrue())
			})

		})
		Context("when it can read the file link", func() {
			It("will read from the file and return an slice with 2 strings in it", func() {
				Expect(fileWordsArray).To(BeEquivalentTo([]string{"hello", "smile"}))
			})

		})
		BeforeEach(func() {
			Whenever(mockPrivate.getText("")).ThenReturn([]string{"hello", "smile"}, nil)
		})

	})

	Describe("getText", func() {
		var stringArrayOfText []string
		var err error
		var link string
		JustBeforeEach(func() {
			stringArrayOfText, err = r.getText(link)
		})
		Context("when the input link is invalid", func() {

			BeforeEach(func() {

				link = "https://gist.github.com/cfraddadwfsf"
			})
			It("will return an error", func() {
				//This test needs to be better
				Expect(err).ToNot(BeNil())
			})
		})
		Context("when the input link is valid and nothing fails", func() {
			It("will return a string slice of all the words", func() {
				//This test needs to be better
				Expect(len(stringArrayOfText)).ToNot(Equal(0))
			})
			It("will return a nil error", func() {
				Expect(err).To(BeNil())
			})
		})
		BeforeEach(func() {
			link = "https://gist.github.com/cfreshman/40608e78e83eb4e1d60b285eb7e9732f/raw/wordle-nyt-allowed-guesses.txt"
		})
	})
	Describe("backUpFiles", func() {
		var filePath string

		var err error
		JustBeforeEach(func() {
			_, err = r.backUpFiles(filePath)
		})
		Context("when the file path is invalid", func() {
			BeforeEach(func() {
				filePath = "badLinkhere"
			})
			It("will return an error", func() {
				Expect(err).ToNot(BeNil())
			})
		})
		BeforeEach(func() {
			filePath = "rga/wordle-allowed-guesses.txt"
		})
	})
})
