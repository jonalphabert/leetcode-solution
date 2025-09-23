type query struct{
    value int
    left int
    right int
}

type MaxHeap []query

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].value > h[j].value }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(query))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func maxTotalValue(nums []int, k int) int64 {
    n := len(nums)

    seen := make(map[query]bool)

    // for query purpose
    logArr := make([]int, n+1)

    for i := 2; i <= n; i++ {
        logArr[i] = logArr[i/2] + 1
    }

    // sparse table max and min
    maxSt := make([][]int, logArr[n]+1)
    minSt := make([][]int, logArr[n]+1)

    for i:= 0; i<=logArr[n]; i++ {
        maxSt[i] = make([]int, n)
        minSt[i] = make([]int, n)
    }

    for i:= 0; i<n; i++ {
        maxSt[0][i]=nums[i]
        minSt[0][i]=nums[i]
    }

    for pw:=1; pw <= logArr[n]; pw++ {
        for i:=0; i+(1 << pw) <= n; i++ {
            minSt[pw][i] = min(minSt[pw-1][i], minSt[pw-1][i + (1 << (pw-1))])
            maxSt[pw][i] = max(maxSt[pw-1][i], maxSt[pw-1][i + (1 << (pw-1))])
        }
    }

	querySparseTable := func(l, r int) int{
        log := logArr[r-l+1]
        return max(maxSt[log][l], maxSt[log][r - (1 << log) +1]) - min(minSt[log][l], minSt[log][r - (1 << log) +1])
    }

    pq := &MaxHeap{}
    heap.Init(pq)

    for i:=0; i<n; i++ {
		queryObject := query{querySparseTable(i, n-1), i, n-1}
        seen[queryObject] = true
		heap.Push(pq, queryObject)
    }

    ans := int64(0)

    for k > 0 {
        maxValue := heap.Pop(pq).(query)

        ans+=int64(maxValue.value)
		l := maxValue.left
		r := maxValue.right

        // add subset on that pruning right of selected substring
        if l+1 <= r{
            object := query{querySparseTable(l+1, r), l+1, r}
            if !seen[object] {
			    heap.Push(pq, object)
                seen[object] = true
            }
        }

        // add subset on that pruning left of selected substring
		if l <= r-1 {
			object := query{querySparseTable(l, r-1), l, r-1}
			if !seen[object] {
			    heap.Push(pq, object)
                seen[object] = true
            }
		}

		k--
    }

    return ans
}