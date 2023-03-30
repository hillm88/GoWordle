package rga

type validator struct{}

// // validateLink Takes a link in the form of a string and checks that it is not empty.
// func (validator) validateLink(link string) error {
// 	if link == "" {
// 		return linkEmpty
// 	}
// 	return nil
// }

// validateBackUpPath takes in the path to the set of backup files and checks that it is not empty.
func (validator) validateBackUpPath(path string) error {
	if path == "" {
		return backUpFilePathEmpty
	}
	return nil
}
