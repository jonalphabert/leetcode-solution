func minSplitMerge(nums1 []int, nums2 []int) int {
	arrCopy := make([]int, len(nums1))
	copy(arrCopy, nums1)

	encode := func(arr []int) string {
		s := ""
		for i, v := range arr {
			if i > 0 {
				s += ","
			}
			s += fmt.Sprintf("%d", v)
		}
		return s
	}

	targetKey := encode(nums2)
	startKey := encode(arrCopy)

	if startKey == targetKey {
		return 0
	}

	type State struct {
		arr   []int
		steps int
	}

	queue := []State{{arrCopy, 0}}
	visited := map[string]bool{startKey: true}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		n := len(cur.arr)
		for L := 0; L < n; L++ {
			for R := L; R < n; R++ {
				sub := append([]int{}, cur.arr[L:R+1]...)
				remain := append([]int{}, cur.arr[:L]...)
				remain = append(remain, cur.arr[R+1:]...)

				for pos := 0; pos <= len(remain); pos++ {
					newArr := append([]int{}, remain[:pos]...)
					newArr = append(newArr, sub...)
					newArr = append(newArr, remain[pos:]...)

					key := encode(newArr)
					if visited[key] {
						continue
					}
					if key == targetKey {
						return cur.steps + 1
					}
					visited[key] = true
					queue = append(queue, State{newArr, cur.steps + 1})
				}
			}
		}
	}

	return -1
}