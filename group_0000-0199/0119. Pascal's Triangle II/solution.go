func getRow(rowIndex int) []int {
    res := make([]int, rowIndex+1)

    res[0] = 1

    for i := 1; i <= rowIndex; i++ {
        for j:= i/2; j >= 0; j-- {
            if j == 0 {
                res[j] = 1
                res[i-j] = 1
                continue
            }
            res[j] = res[j] + res[j-1]
            res[i-j] = res[j]
        }
    }
    return res
}