func sumZero(n int) []int {
    res := make([]int, 0)

    if n%2 == 1 {
        res = append(res, 0)
    }

    for i:=1; i<=n/2; i++ {
        res = append(res, i)
        res = append(res, -1*i)
    }

    return res
}