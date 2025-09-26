func groupAnagrams(strs []string) [][]string {
    hashMap := make(map[string][]string)

    for _, s := range strs {
        b := []byte(s)

        sort.Slice(b, func(i, j int) bool {
            return b[i] < b[j]
        })

        sorted := string(b)

        hashMap[sorted] = append(hashMap[sorted], s)
    }

    var res [][]string

    for _, value := range hashMap {
        res = append(res, value)
    }

    return res
}