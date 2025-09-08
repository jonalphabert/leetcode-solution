func longestValidParentheses(s string) int {
    maxValue := 0
    if len(s) < 2 {
        return maxValue
    }

    dp := make([]int, len(s))

    dp[0] = 0

    for i := 1; i < len(s); i++ {
        if s[i] == ')' {
            if s[i-1] == '(' {
                // case "()"
                if i-2 >= 0 {
                    dp[i] = dp[i-2] + 2
                } else {
                    dp[i] = 2
                }
            } else if s[i-1] == ')' {
                // case "...))"
                prevLen := dp[i-1]
                if i-prevLen-1 >= 0 && s[i-prevLen-1] == '(' {
                    if i-prevLen-2 >= 0 {
                        dp[i] = prevLen + 2 + dp[i-prevLen-2]
                    } else {
                        dp[i] = prevLen + 2
                    }
                }
            }
            maxValue = max(maxValue, dp[i])
        }
    }

    return maxValue
}