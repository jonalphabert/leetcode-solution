func checkAnyZeroDigit(num int) bool {
    for num > 0 {
        if num % 10 == 0 {
            return false
        }
        num /= 10
    }
    return true
}

func getNoZeroIntegers(n int) []int {
    for i := 1; i<=n/2; i++ {
        if checkAnyZeroDigit(i) && checkAnyZeroDigit(n-i) {
            return []int{i, n-i}
        }
    }
    return []int{}
}