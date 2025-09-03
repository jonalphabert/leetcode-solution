func romanToInt(s string) int {
    values := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    res := 0
    lastValue := math.MaxInt

    for i := 0; i < len(s); i++ {
        char := s[i]
        if lastValue < values[char] {
            res += values[char] - lastValue*2
        } else {
            lastValue = values[char]
            res += values[char]
        }
    }

    return res
}