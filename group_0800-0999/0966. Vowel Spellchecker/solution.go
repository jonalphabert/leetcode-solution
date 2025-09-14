func spellchecker(wordlist []string, queries []string) []string {
    wordSet := make(map[string]bool)
	caseMap := make(map[string]string)
	vowelMap := make(map[string]string)

	devowel := func(word string) string {
		vowels := "aeiou"
		lower := strings.ToLower(word)
		var sb strings.Builder
		for _, ch := range lower {
			if strings.ContainsRune(vowels, ch) {
				sb.WriteRune('*')
			} else {
				sb.WriteRune(ch)
			}
		}
		return sb.String()
	}

	for _, word := range wordlist {
		wordSet[word] = true

		lower := strings.ToLower(word)
		if _, ok := caseMap[lower]; !ok {
			caseMap[lower] = word
		}

		vword := devowel(word)
		if _, ok := vowelMap[vword]; !ok {
			vowelMap[vword] = word
		}
	}

	ans := make([]string, len(queries))
	for i, query := range queries {
		if wordSet[query] {
			ans[i] = query
		} else if val, ok := caseMap[strings.ToLower(query)]; ok {
			ans[i] = val
		} else if val, ok := vowelMap[devowel(query)]; ok {
			ans[i] = val
		} else {
			ans[i] = ""
		}
	}

	return ans
}