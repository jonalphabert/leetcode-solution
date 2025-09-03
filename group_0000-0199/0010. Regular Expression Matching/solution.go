func isMatch(s string, p string) bool {
    var dfs func(i, j int) bool 
    dfs = func(i, j int) bool {
        if j == len(p) {
            return i == len(s)
        }

        match := i < len(s) && (s[i] == p[j] || p[j] == '.')

        if j+1 < len(p) && p[j+1] == '*' {
            return dfs(i, j+2) || (match && dfs(i+1, j))
        } else {
            return match && dfs(i+1, j+1)
        }
    }

    return dfs(0,0)
}