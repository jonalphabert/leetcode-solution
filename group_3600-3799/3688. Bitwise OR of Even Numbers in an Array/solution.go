func evenNumberBitwiseORs(nums []int) int {
    res := 0

    for _, num := range nums {
        if num % 2 == 0 {
            res |= num
        }
    }
    return res
}