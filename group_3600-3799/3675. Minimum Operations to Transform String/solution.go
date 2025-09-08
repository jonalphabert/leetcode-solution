func minOperations(s string) int {
    res := -1
    for _, val := range s {
        res = max(res, (26 - int(val - 'a'))%26)
    }
    return res
}