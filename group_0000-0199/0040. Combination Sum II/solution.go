func combinationSum2(candidates []int, target int) [][]int {
    res := [][]int{}
    sort.Ints(candidates)

    var dfs func(start int, comb []int, remain int)
    dfs = func(start int, comb []int, remain int) {
        if remain == 0 {
            temp := make([]int, len(comb))
            copy(temp, comb)
            res = append(res, temp)
            return
        }

        for i := start; i < len(candidates); i++ {
            // skip duplicate di level yang sama
            if i > start && candidates[i] == candidates[i-1] {
                continue
            }

            if candidates[i] > remain {
                break
            }

            comb = append(comb, candidates[i])
            dfs(i+1, comb, remain - candidates[i])
            comb = comb[:len(comb)-1]
        }
    }

    dfs(0, []int{}, target)
    return res
}
