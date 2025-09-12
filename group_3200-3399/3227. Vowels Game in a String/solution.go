func doesAliceWin(s string) bool {
    for _, char := range s {
        if char == 'a' || char == 'i' || char == 'u' || char == 'e' || char == 'o' {
            return true
        }
    }

    return false
}