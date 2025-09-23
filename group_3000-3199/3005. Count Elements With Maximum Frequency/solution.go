func maxFrequencyElements(nums []int) int {
    freq := make(map[int]int)

    for _, val := range nums {
        freq[val]++
    }

    maxFreq := 0
    result := 0

    for _, val := range freq {
        if val > maxFreq {
            maxFreq = val
            result = val
        } else if val == maxFreq {
            result += val
        }

    }

    return result
}