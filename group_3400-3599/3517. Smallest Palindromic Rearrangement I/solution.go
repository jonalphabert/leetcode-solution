func smallestPalindrome(s string) string {
    // buat frekuensi per charnya
    mapCharCount := make([]int, 26)

    for i := 0; i < len(s); i++ {
        mapCharCount[s[i]-'a']++
    }

    var stringBuilder strings.Builder

    for i := 0; i<26; i++ {
        stringBuilder.WriteString(strings.Repeat(string(i+'a'), mapCharCount[i]/2))
    }

    if len(s)%2 == 1 {
        stringBuilder.WriteByte(s[len(s)/2])
    }

    for i := 25; i>=0; i-- {
        stringBuilder.WriteString(strings.Repeat(string(i+'a'), mapCharCount[i]/2))
    }

    return stringBuilder.String()
}