func countDigitOne(n int) int {
    digits := []int{}

    for _, ch := range strconv.Itoa(n) {
        digits = append(digits, int(ch-'0'))
    }

    dp := make(map[[3]int]int)

    var dfs func (pos int, countOne int, tight int) int 
    dfs = func (pos int, countOne int, tight int) int {
        if pos == len(digits) {
            return countOne
        }

        key := [3]int{pos, countOne, tight}
        if val, exist := dp[key]; exist {
            return val
        }

        limit := 9
        res := 0

        if tight == 1 {
            limit = digits[pos]
        }

        for i := 0; i <= limit; i++ {
            newTight := 0
            newCountOne := countOne
            if tight == 1 && i == limit {
                newTight = 1
            }

            if i == 1 {
                newCountOne ++
            }

            res += dfs(pos+1, newCountOne, newTight)
        }

        dp[key] = res
        return res
    }

    return dfs(0, 0, 1)
}