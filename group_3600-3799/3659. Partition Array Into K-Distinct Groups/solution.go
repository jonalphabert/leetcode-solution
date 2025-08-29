func partitionArray(nums []int, k int) bool {
    if len(nums) % k != 0 {
        return false
    }

    group := len(nums) / k
    freq := make(map[int]int)

    for _, val := range nums {
        freq[val] ++;
        if freq[val] > group {
            return false
        }
    }
    
    return true
}