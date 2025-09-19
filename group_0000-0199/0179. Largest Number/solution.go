func largestNumber(nums []int) string {
    strNums := make([]string, len(nums))
    for i, n := range nums {
        strNums[i] = strconv.Itoa(n)
    }

    sort.Slice(strNums, func(i, j int) bool {
        return strNums[i]+strNums[j] > strNums[j]+strNums[i]
    })

    result := strings.Join(strNums, "")

    if result[0] == '0' {
        return "0"
    }

    return result
}