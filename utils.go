package randomize

// StringIsInArray will check if a string is in an array of strings.
func StringIsInArray(test string, target []string) bool {
	return ElementIsInArray(test, target)
}

// ElementIsInArray will check if an element is in an array of elements.
func ElementIsInArray[T comparable](testElement T, array []T) bool {
	for _, v := range array {
		if testElement == v {
			return true
		}
	}
	return false
}
