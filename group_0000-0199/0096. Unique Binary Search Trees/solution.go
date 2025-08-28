var memo = map[int]int{0: 1}

func numTrees(n int) int {
    if val, ok := memo[n]; ok {
		return val
	}
	var res int = 0
	for i := 0; i < n; i++ {
		res += numTrees(i) * numTrees(n-1-i)
	}
	memo[n] = res
	return res
}