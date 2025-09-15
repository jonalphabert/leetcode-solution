func canBeTypedWords(text string, brokenLetters string) int {
    parts := strings.Split(text, " ")
	brokenLettersMap := make(map[rune]bool)

	for _, letter := range brokenLetters {
		brokenLettersMap[letter] = true
	}

	brokenWord := 0

	for _, part := range parts {
		for _, letter := range part {
			if brokenLettersMap[letter] {
				brokenWord++
				break
			}
		}
	}

	return len(parts) - brokenWord
}