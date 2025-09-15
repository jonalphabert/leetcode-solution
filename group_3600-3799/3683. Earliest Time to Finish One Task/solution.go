func earliestTime(tasks [][]int) int {
    minTime := math.MaxInt

    for _, val := range tasks {
        minTime = min(minTime, val[0] + val[1])
    }

    return minTime
}