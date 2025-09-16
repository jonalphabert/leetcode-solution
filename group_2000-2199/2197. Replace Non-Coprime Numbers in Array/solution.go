func gcd(a, b int) int {
    for b > 0 {
        rem := a % b
        a = b
        b = rem
    }
    return a
}

func replaceNonCoprimes(nums []int) []int {
    res := []int{}
    res = append(res, nums[0])

    indexRes := 0

    for i:=1; i<len(nums); i++ {
        gcdAdjNum := gcd(nums[i], res[indexRes])

        if gcdAdjNum == 1 {
            res = append(res, nums[i])
            indexRes++
        } else {
            res[indexRes] = res[indexRes] * nums[i] / gcdAdjNum

            for indexRes > 0 {
                val := gcd(res[indexRes], res[indexRes-1])

                if val == 1 {
                    break
                }

                res[indexRes-1] = res[indexRes] * res[indexRes-1] / val
                res = res[:indexRes]
                indexRes--
            }
        }
    }

    return res
}