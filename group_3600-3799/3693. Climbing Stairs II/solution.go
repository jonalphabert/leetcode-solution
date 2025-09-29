func climbStairs(n int, costs []int) int {
	bestRoute := make([]int, n+1)
	bestRoute[0] = 0

	for i := 1; i <= n; i++ {
		bestRoute[i] = 1<<31 - 1
		for j := 1; j <= 3; j++ {
			if i-j >= 0 {
				bestRoute[i] = min(bestRoute[i], bestRoute[i-j]+costs[i-1]+j*j)
			}
		}
	}

	return bestRoute[n]
}