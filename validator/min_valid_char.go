package validator

// minValidChar returns true when char count >= min
func minValidChar(min int, value string, validateChar func(rune) bool) bool {
	var count = 0
	for _, r := range value {
		if validateChar(r) {
			count += 1
		}
	}

	return count >= min
}
