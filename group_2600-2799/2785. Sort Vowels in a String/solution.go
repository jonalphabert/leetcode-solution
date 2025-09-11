func sortVowels(s string) string {
    isVowel := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
    byteArray := []byte(s)
    vowelArray := []byte{}
    posVowel := []int{}


    for i := 0; i < len(s); i++ {
		if isVowel[s[i]] {
			vowelArray = append(vowelArray, s[i])
            posVowel = append(posVowel, i)
		}
	}

    sort.Slice(vowelArray, func(i, j int) bool {
        return int(vowelArray[i]) < int(vowelArray[j]) 
    })

    for index, indexVowel := range posVowel {
        byteArray[indexVowel] = vowelArray[index]
    }

    return string(byteArray)
}