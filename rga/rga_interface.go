package rga

type privateMethodInterface interface {
	mapIt(listOfWords []string) map[string]string
	getText(link string) ([]string, error)
	backUpFiles(backUpFile string) ([]string, error)
	fileRetriever(fileLink string, backUpLocalPath string) ([]string, error)
}

type validatorInterface interface {
	validateBackUpPath(path string) error
}
