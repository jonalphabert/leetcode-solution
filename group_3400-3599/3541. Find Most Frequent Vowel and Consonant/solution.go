func maxFreqSum(s string) int {
    frequency := make([]int, 26)
    maxVowel, maxConsonant := 0, 0

    for i:=0; i<len(s); i++ {
        char := byte(s[i])
        frequency[int(char-'a')]++

        if char == 'a' || char == 'i' || char == 'u' || char == 'e' || char == 'o' {
            maxVowel = max(maxVowel, frequency[int(char-'a')])
        } else {
            maxConsonant = max(maxConsonant, frequency[int(char-'a')])
        }
    }

    return maxConsonant + maxVowel
}