func majorityFrequencyGroup(s string) string {
    freq := make(map[rune]int)

    for _, char := range s {
        freq[char]++
    }

    mapGroup := make(map[int][]rune)

    for key, val := range freq {
        mapGroup[val] = append(mapGroup[val], key)
    }

    maxSize := 0
	bestFreq := 0 
	var bestGroup []rune

	for f, chars := range mapGroup {
		if len(chars) > maxSize || (len(chars) == maxSize && f > bestFreq) {
			maxSize = len(chars)
			bestFreq = f
			bestGroup = chars
		}
	}

	return string(bestGroup)
}