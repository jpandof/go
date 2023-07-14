package palindromo

func PalindromoString(value string) bool {

	if len(value) <= 1 {
		return true
	}

	for i := 0; i < len(value)/2; i++ {
		firstLetter := value[i]
		lastLetter := value[len(value)-(i+1)]

		if firstLetter != lastLetter {
			return false
		}
	}

	return true

}
