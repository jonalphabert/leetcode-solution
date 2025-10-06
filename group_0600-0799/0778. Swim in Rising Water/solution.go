type Node struct {
	time, i, j int
}

type PriorityQueue []Node

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].time < pq[j].time
}
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func swimInWater(grid [][]int) int {
	n := len(grid)
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	pq := &PriorityQueue{{grid[0][0], 0, 0}}
	heap.Init(pq)
	res := 0

	for pq.Len() > 0 {
		node := heap.Pop(pq).(Node)
		i, j, t := node.i, node.j, node.time
		if visited[i][j] {
			continue
		}
		visited[i][j] = true
		if i == n-1 && j == n-1 {
			if t > res {
				res = t
			}
			return res
		}
		if t > res {
			res = t
		}

		for _, d := range dirs {
			ni, nj := i+d[0], j+d[1]
			if ni >= 0 && ni < n && nj >= 0 && nj < n && !visited[ni][nj] {
				heap.Push(pq, Node{grid[ni][nj], ni, nj})
			}
		}
	}
	return res
}