type Cell struct {
	h, i, j int
}

type PriorityQueue []Cell

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].h < pq[j].h } 
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Cell))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func trapRainWater(heightMap [][]int) int {
	m, n := len(heightMap), len(heightMap[0])
	if m <= 2 || n <= 2 {
		return 0
	}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	pq := &PriorityQueue{}
	heap.Init(pq)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 || i == m-1 || j == n-1 {
				heap.Push(pq, Cell{heightMap[i][j], i, j})
				visited[i][j] = true
			}
		}
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	res := 0

	for pq.Len() > 0 {
		cell := heap.Pop(pq).(Cell)

		for _, d := range dirs {
			ni, nj := cell.i+d[0], cell.j+d[1]
			if ni >= 0 && nj >= 0 && ni < m && nj < n && !visited[ni][nj] {
				visited[ni][nj] = true
				if heightMap[ni][nj] < cell.h {
					res += cell.h - heightMap[ni][nj]
					heap.Push(pq, Cell{cell.h, ni, nj})
				} else {
					heap.Push(pq, Cell{heightMap[ni][nj], ni, nj})
				}
			}
		}
	}

	return res
}