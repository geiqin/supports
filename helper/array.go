package helper

func InArray(s [] string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}