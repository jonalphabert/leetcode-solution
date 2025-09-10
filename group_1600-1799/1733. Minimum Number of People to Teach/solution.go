func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
    m := len(languages)

    setFailedCommunication := make(map[int]bool)
    langs := make([]map[int]bool, m+1)
    countFailedLanguage := make([]int, n+1)
    maxFailedLanguage := 0

    for i := 1; i <= m; i++ {
        langs[i] = map[int]bool{}
        for _, lang := range languages[i-1] {
            langs[i][lang] = true
        }
    }

    canCommunicate := func(u, v int) bool {
        if len(langs[u]) > len(langs[v]) {
            u, v = v, u
        }
        for lang := range langs[u] {
            if langs[v][lang] {
                return true
            }
        }
        return false
    }

    for _, friendship := range friendships {
        if canCommunicate(friendship[0], friendship[1]) {
            continue
        }

        setFailedCommunication[friendship[0]] = true
        setFailedCommunication[friendship[1]] = true
    }

    for key, _ := range setFailedCommunication {
        for _, language := range languages[key-1] {
            countFailedLanguage[language]++
            maxFailedLanguage = max(maxFailedLanguage, countFailedLanguage[language])
        }
    }

    return len(setFailedCommunication) - maxFailedLanguage
}