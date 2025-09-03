func isMatch(s string, p string) bool {
    memo := make(map[[2]int]bool)

    var dfs func(i, j int) bool
    dfs = func(i, j int) bool {
        if val, ok := memo[[2]int{i, j}]; ok {
            return val
        }

        if j == len(p) {
            return i == len(s)
        }

        match := i < len(s) && (s[i] == p[j] || p[j] == '.')

        var ans bool
        if j+1 < len(p) && p[j+1] == '*' {
            ans = dfs(i, j+2) || (match && dfs(i+1, j))
        } else {
            ans = match && dfs(i+1, j+1)
        }

        memo[[2]int{i, j}] = ans
        return ans
    }

    return dfs(0, 0)
}