func maxAlternatingSum(nums []int, swaps [][]int) int64 {
    n := len(nums)

	graph := make([][]int, n)
	for _, e := range swaps {
		i, j := e[0], e[1]
		graph[i] = append(graph[i], j)
		graph[j] = append(graph[j], i)
	}

	visited := make([]bool, n)
	total := int64(0)

	var dfs func(int, *[]int)
	dfs = func(u int, currentArray *[]int) {
		visited[u] = true
		*currentArray = append(*currentArray, u)
		for _, v := range graph[u] {
			if !visited[v] {
				dfs(v, currentArray)
			}
		}
	}

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		comp := []int{}
		dfs(i, &comp)

		vals := []int{}
		countEven := 0

		for _, idx := range comp {
			vals = append(vals, nums[idx])
			if idx % 2 == 0 {
				countEven++
			}
		}

		sort.Sort(sort.Reverse(sort.IntSlice(vals)))

		for index, value := range vals {
			if index < countEven {
				total += int64(value)
			} else {
				total -= int64(value)
			}
		}
	}

	return total
}