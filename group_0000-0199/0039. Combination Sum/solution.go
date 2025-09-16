func combinationSum(candidates []int, target int) [][]int {
    res := [][]int{}

    sort.Ints(candidates)

    var combination func(start int, combinationArr []int, remaining int)
    combination = func(start int, combinationArr []int, remaining int){
        if remaining == 0 {
            temp := make([]int, len(combinationArr))
            copy(temp, combinationArr)
            res = append(res, temp)
        } else {
            for i:=start; i < len(candidates); i++ {
                if remaining - candidates[i] < 0 {
                    break
                }

                combinationArr = append(combinationArr, candidates[i])

                combination(i, combinationArr, remaining - candidates[i])

                combinationArr = combinationArr[:len(combinationArr)-1]
            }
        }
    }

    combination(0,[]int{}, target)

    return res
}