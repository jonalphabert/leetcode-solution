func countDigitOne(n int) int {
    factor := 1000000000
    res := 0
    var high, currentDigit int

    for factor > 0 {
        high = n / (factor * 10)
        currentDigit = (n / factor) % 10
        // fmt.Println(high, currentDigit, factor)
        res += factor * high
        if currentDigit == 1 {
            res += n % factor + 1
        } else if currentDigit > 1 {
            res += factor
        }
        factor /= 10
    }

    return res
}