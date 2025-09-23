// ChatGPT Generation for testing memory allocation only
type SegmentTree struct {
	n   int
	max []int
	min []int
}

func NewSegmentTree(nums []int) *SegmentTree {
	n := len(nums)
	size := 1
	for size < n {
		size <<= 1
	}
	maxTree := make([]int, 2*size)
	minTree := make([]int, 2*size)

	// init leaves
	for i := 0; i < n; i++ {
		maxTree[size+i] = nums[i]
		minTree[size+i] = nums[i]
	}
	for i := n; i < size; i++ {
		maxTree[size+i] = -1 << 31
		minTree[size+i] = 1<<31 - 1
	}

	// build
	for i := size - 1; i > 0; i-- {
		maxTree[i] = max(maxTree[2*i], maxTree[2*i+1])
		minTree[i] = min(minTree[2*i], minTree[2*i+1])
	}

	return &SegmentTree{size, maxTree, minTree}
}

func (st *SegmentTree) Query(l, r int) int {
	l += st.n
	r += st.n
	mx, mn := -1 << 31, 1<<31-1
	for l <= r {
		if l%2 == 1 {
			mx = max(mx, st.max[l])
			mn = min(mn, st.min[l])
			l++
		}
		if r%2 == 0 {
			mx = max(mx, st.max[r])
			mn = min(mn, st.min[r])
			r--
		}
		l /= 2
		r /= 2
	}
	return mx - mn
}

// ================== Max Heap ==================

type query struct {
	value int
	left  int
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

// ================== Main Logic ==================

func maxTotalValue(nums []int, k int) int64 {
	n := len(nums)
	st := NewSegmentTree(nums)

	// seen set pakai bit-packing biar hemat memori
	seen := make(map[uint64]struct{})

	pack := func(l, r int) uint64 {
		return (uint64(l) << 32) | uint64(r)
	}

	pq := &MaxHeap{}
	heap.Init(pq)

	for i := 0; i < n; i++ {
		q := query{st.Query(i, n-1), i, n - 1}
		heap.Push(pq, q)
		seen[pack(i, n-1)] = struct{}{}
	}

	ans := int64(0)

	for k > 0 && pq.Len() > 0 {
		maxValue := heap.Pop(pq).(query)
		ans += int64(maxValue.value)
		l, r := maxValue.left, maxValue.right

		if l+1 <= r {
			key := pack(l+1, r)
			if _, ok := seen[key]; !ok {
				heap.Push(pq, query{st.Query(l+1, r), l + 1, r})
				seen[key] = struct{}{}
			}
		}
		if l <= r-1 {
			key := pack(l, r-1)
			if _, ok := seen[key]; !ok {
				heap.Push(pq, query{st.Query(l, r-1), l, r - 1})
				seen[key] = struct{}{}
			}
		}
		k--
	}

	return ans
}