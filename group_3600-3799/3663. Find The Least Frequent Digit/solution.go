func getLeastFrequentDigit(n int) int {
    freq := make(map[int]int)

    for i:=0; i<=9; i++{
        freq[i] = 0;
    }

    for n > 0 {
        freq[n%10]++
        n /= 10
    }

    minFreq := -1
    minDigit := -1

    for i := 0; i <= 9; i++ {
        if freq[i] > 0 && (minFreq == -1 || freq[i] < minFreq) {
            minFreq = freq[i]
            minDigit = i
        }
    }

    return minDigit
}