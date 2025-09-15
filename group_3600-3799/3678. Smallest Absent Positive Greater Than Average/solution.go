func smallestAbsent(nums []int) int {
    isThere := make(map[int]bool)
    sum := 0
    n := len(nums)

    for _, val := range nums {
        sum += val
        isThere[val] = true
    }

    avg :=sum/n
    avg++
    if avg <= 0 {
        avg = 1
    }

    for isThere[avg] {
        avg++
    }

    return avg
}